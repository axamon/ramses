package funzioni

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/axamon/ramses/algoritmi"
	"gonum.org/v1/gonum/stat"
)

const sessioniPPP = "https://ipw.telecomitalia.it/ipwmetrics/api/v1/rawmetrics/kpi.ppoe.slot?device="

var wgppp sync.WaitGroup

//Nasppp recupera le sessioni ppp dei NAS
func Nasppp() (output chan string) {
	//Creo il contesto inziale che verrà propagato alle go-routine
	ctx := context.Background()

	//Creo la variabile dove accodare i nomi dei nas
	var devices []string

	//identifico il file json con le informazioni da parsare
	filelistapparati := "nasInventory.json"

	//leggo il file in memoria
	body, err := ioutil.ReadFile(filelistapparati)
	if err != nil {
		log.Printf("%s Error Impossibile recuperare lista\n", filelistapparati)
	}

	//Creo la variabile dove conservare i dati parsati
	var listalistanas [][]TNAS
	errjson := json.Unmarshal(body, &listalistanas)
	if errjson != nil {
		log.Printf("%s Error Impossibile parsare dati\n", filelistapparati)
	}

	//Istanzio un contatore per contare i nas trovati
	var i int

	//listalistanas è una lista di liste quindi bisogna fare un doppio ciclo for
	for _, listanas := range listalistanas {
		for _, nas := range listanas {
			//fmt.Println(n, nas.Name) //debug

			//considero solo gli apparati che abbiano "NAS" all'inzio del campo Service
			//e EDGE_BRAS come dominio e MX960 come chassis
			if strings.HasPrefix(nas.Service, "NAS") && strings.Contains(nas.Domain, "EDGE_BRAS") && strings.Contains(nas.ChassisName, "MX960") {
				//incremento il contatore
				i++

				//Appendo in devices il nome nas trovato
				devices = append(devices, nas.Name)
			}
		}
	}
	//loggo il numero di NAS identificati
	log.Printf("%d INFO numero di NAS trovati\n", i)
	time.Sleep(3 * time.Second)

	//recuperaSessioniPPP è una funzione che recupera i dati ppp dei nas
	recuperaSessioniPPP := func() {
		//espando il contesto inziale inserendo un timeout
		ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)

		//prima di chiudere vengono rilasciate tutte le risorse
		defer cancel()

		//Avvio un ciclo infinito
		for {
			select {
			//se viene raggiunto il timeout la funzione viene killata
			case <-ctx.Done():
				log.Printf("All Error Tempo per completare recupero dati terminato\n")
				return

			//finchè non si raggiunge il timeout viene eseguito il codice di default
			default:
				for _, device := range devices {
					wgppp.Add(1)
					log.Printf("%s Info Inzio verifica device\n", device)
					//go nasppp2(device)
					nasppp2(ctx, device, output)
				}
				return
			}
		}

	}

	//Prima esecuzione del recupero dati dall'avvio dell'applicazione
	recuperaSessioniPPP()

	//Attende che tutte le richieste siano terminate prima di proseguire
	wgppp.Wait()
	fmt.Println("Dopo primo run") //debug

	//imposta un refesh ogni tot minuti
	t := time.Tick(30 * time.Second)
	c := time.Tick(5 * time.Minute)
	for {
		select {
		case <-c:
			recuperaSessioniPPP()
			wgppp.Wait()
		case <-t:
			fmt.Println(".")
		}
	}

}

func nasppp2(ctx context.Context, device string, output chan string) {
	//Riceve il contesto padre e aggiunge un timeout
	//massimo per terminare la richiesta dati
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()
	defer log.Printf("%s Info Recupero dati terminato\n", device)
	defer wgppp.Done()

	for {
		select {
		case <-ctx.Done():
			log.Printf("%s Error Superato tempo massimo per raccolta dati\n", device)
			return

		default:
			//Attendo un tempo random per evitare di fare troppe query insieme
			randomdelay := rand.Intn(100)
			time.Sleep(time.Duration(randomdelay) * time.Millisecond)

			//Ripulisce eventiali impostazioni di proxy a livello di sistema
			os.Setenv("HTTP_PROXY", "")
			os.Setenv("HTTPS_PROXY", "")
			fmt.Println(os.Getenv("HTTP_PROXY"))
			fmt.Println(os.Getenv("HTTPS_PROXY"))

			//fmt.Println(device)

			var sigma float64
			sigma = 2.0

			//Recupera la variabile d'ambiente
			username, err := Recuperavariabile("username")
			if err != nil {
				log.Fatal(err)
				return
			}

			//Recupera la variabile d'ambiente
			password, err := Recuperavariabile("password")
			if err != nil {
				log.Fatal(err)
				return
			}
			url := sessioniPPP + device + "&start=7d-ago&end=5m-ago&aggregator=sum"

			req, _ := http.NewRequest("GET", url, nil)

			//qui su costringe il client ad accettare anche certificati https non validi o scaduti, non anrebbe fatto ma bisogna fare di necessità virtù
			transCfg := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
			}

			//req.Header.Add("content-type", "application/json;charset=UTF-8")
			req.SetBasicAuth(username, password)
			//req.Header.Add("authorization", "Basic MDAyNDY1MDY6Y2Z4VyRsTVM2ZA==")
			req.Header.Add("cache-control", "no-cache")

			client := &http.Client{Transport: transCfg}

			res, _ := client.Do(req)
			//res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err.Error(), device)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(err.Error(), device)
			}
			defer res.Body.Close()
			//fmt.Println(res)
			//	fmt.Println(string(body))
			var result []interface{}
			err = json.Unmarshal(body, &result)
			if err != nil {
				log.Println(err.Error(), device)
				return
			}
			if len(result) < 1 {
				log.Printf("%s Error Non ci sono abbastanza info\n", device)
				return
			}
			d := result[0].(map[string]interface{})
			dp := d["dps"].(map[string]interface{})

			//Metti i tempi in ordine
			tempi := make([]string, 0)
			for t := range dp {
				tempi = append(tempi, t)
			}
			//Ordina i tempi in maniera crescente
			sort.Strings(tempi)

			//Crea variabili da'appoggio
			var seriepppvalue []float64
			var serieppptime []float64

			//Cicla i tempi
			for _, t := range tempi {
				tint, _ := strconv.Atoi(t)
				serieppptime = append(serieppptime, float64(tint))
				seriepppvalue = append(seriepppvalue, dp[t].(float64))
				//fmt.Println("orario: ", t, "valore: ", dp[t])
			}

			//elimino il trend
			xdet, ydet := algoritmi.Detrend(serieppptime, seriepppvalue)

			//applico derivata terza alle ordinate
			yderived, _ := algoritmi.Derive3(ydet)

			//Elimina i valori troppo estremi
			yvalues := algoritmi.ScremaValori(yderived, 0.98, 0.02)

			mean, stdev := stat.MeanStdDev(yvalues, nil)
			log.Printf("%s Info media: %2.f stdev: %2.f", device, mean, stdev)

			//Passo le info alla fuzione di elaborazione e grafico
			wg.Add()
			elaboraseriePPP(ctx, xdet, yvalues, device, "test", "ppp")

			//Verifica se ci sono errori da segnalare
			for _, v := range seriepppvalue[len(seriepppvalue)-1:] {
				//fmt.Println(v)
				//Se le sessioni salgono non è importante
				// if v > mean+sigma*stdev {
				// 	log.Printf("Alert su %s, forte innalzamento sessioni ppp\n", device)
				// 	grafanaurl := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
				// 	msg <- fmt.Sprintf("Alert su %s, forte innalzamento sessioni ppp, %s\n", device, grafanaurl)
				// }
				if v < mean-sigma*stdev {
					log.Printf("%s Alert, forte abbassamento sessioni ppp\n", device)
					grafanaurl := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
					output <- fmt.Sprintf("Alert su %s, forte abbassamento sessioni ppp, %s\n", device, grafanaurl)
				}
			}

			return
		}
	}
}
