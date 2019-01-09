package main

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

var wgppp sync.WaitGroup

func nasppp() {
	//Creo il contesto inziale che verrà propagato alle go-routine
	ctx := context.Background()

	defer mandamailChiusura(configuration.SmtpFrom, configuration.SmtpTo)

	//verifica l'avvio di mail. Se non manda mail esce.
	err := mandamailAvvio(configuration.SmtpFrom, configuration.SmtpTo)

	if err != nil {
		log.Println(err.Error())
		log.Fatal(err.Error())
		os.Exit(1)
	}

	//Creo la variabile dove accodare i nomi dei nas
	var devices []string

	//TODO creare il file con i nomi NAS dinamicamente

	//identifico il file json con le informazioni da parsare
	filelistapparati := configuration.NasInventory

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
					//log.Printf("%s Info Inzio verifica device\n", device)
					//go nasppp2(device)
					nasppp2(ctx, device)
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
	update := time.Tick(24 * time.Hour)
	for {
		select {
		case <-update:
			mandamailUpdate(configuration.SmtpFrom, configuration.SmtpTo)
		case <-c:
			recuperaSessioniPPP()
			wgppp.Wait()
		case <-t:
			fmt.Println(".")
		}
	}

}

func nasppp2(ctx context.Context, device string) {
	//Riceve il contesto padre e aggiunge un timeout
	//massimo per terminare la richiesta dati
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()
	//defer log.Printf("%s Info Recupero dati terminato\n", device)
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
			sigma = configuration.Sigma

			//Recupera le credenziali per IPDOM
			username := configuration.IPDOMUser
			password := configuration.IPDOMPassword

			url := configuration.URLSessioniPPP + device + configuration.URLTail7d

			req, _ := http.NewRequest("GET", url, nil)

			//qui su costringe il client ad accettare anche certificati https non validi o scaduti, non anrebbe fatto ma bisogna fare di necessità virtù
			transCfg := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
			}

			//req.Header.Add("content-type", "application/json;charset=UTF-8")
			req.SetBasicAuth(username, password)
			req.Header.Add("cache-control", "no-cache")

			client := &http.Client{Transport: transCfg}

			res, err := client.Do(req)
			//res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err.Error(), device)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println(err.Error(), device)
				os.Exit(1)
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

			//Se non ci sono abbastanza valori per la serie esci
			if len(seriepppvalue) < 300 {
				log.Printf("%s Error Non ci sono abbastanza dati per elaborare statistiche", device)
				return
			}

			//Calcola statistiche sulla serie prima che sia elaborata
			//meanreale, stdevreale := stat.MeanStdDev(seriepppvalue, nil)
			//log.Printf("%s Info media: %2.f stdev: %2.f", device, meanreale, stdevreale)

			//elimino il trend
			//xdet, ydet := algoritmi.Detrend(serieppptime, seriepppvalue)
			xdet, ydet := algoritmi.Detrend(serieppptime, seriepppvalue)

			//applico derivata terza alle ordinate
			yderived, _ := algoritmi.Derive3(ydet)

			//Elimina i valori troppo alti o bassi
			y := algoritmi.ScremaValori(yderived, 0.99, 0.01)

			//Passo le info alla fuzione di elaborazione e grafico
			wg.Add()
			elaboraseriePPP(ctx, xdet, y, device, "test", "ppp")
			wg.Wait()

			//Calcola statistiche sulla serie elaborata
			mean, stdev := stat.MeanStdDev(y, nil)
			//log.Printf("%s Info media: %2.f stdev: %2.f", device, mean, stdev)

			//Verifica se ci sono errori da segnalare negli ultimi valori y della serie
			for _, v := range y[len(y)-2 : len(y)-1] {
				log.Printf("%s Info media: %2.f stdev: %2.f Lastvalue: %2.f", device, mean, stdev, v)
				//fmt.Println(v)
				//Se le sessioni salgono non è importante
				// if v > mean+sigma*stdev {
				// 	log.Printf("Alert su %s, forte innalzamento sessioni ppp\n", device)
				// 	grafanaurl := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
				// 	msg <- fmt.Sprintf("Alert su %s, forte innalzamento sessioni ppp, %s\n", device, grafanaurl)
				// }

				//Se il valore è minore di sigma volte la media allora allarma
				if v < mean-sigma*stdev {
					log.Printf("%s Alert, forte abbassamento sessioni ppp\n", device)
					//grafanaurl := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
					//<- fmt.Sprintf("Alert su %s, forte abbassamento sessioni ppp, %s\n", device, grafanaurl)
					mandamailAlert(configuration.SmtpFrom, configuration.SmtpTo, device)
				}
			}

			return
		}
	}
}
