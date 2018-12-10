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
)

const sessioniPPP = "https://ipw.telecomitalia.it/ipwmetrics/api/v1/rawmetrics/kpi.ppoe.slot?device="

var wgppp sync.WaitGroup

func nasppp() {
	//Creo il contesto inziale che verrà propagato alle go-routine
	ctx := context.Background()

	//Creo la variabile dove accodare i nomi dei nas
	var devices []string

	//identifico il file json con le informazioni da parsare
	filelistapparati := fileNasInventory

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

	return

	// //imposta un refesh ogni tot minuti
	// t := time.Tick(30 * time.Second)
	// c := time.Tick(5 * time.Minute)
	// for {
	// 	select {
	// 	case <-c:
	// 		recuperaSessioniPPP()
	// 		wgppp.Wait()
	// 	case <-t:
	// 		fmt.Println(".")
	// 	}
	// }

}

func nasppp2(ctx context.Context, device string) {
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

			//Recupera la variabile d'ambiente
			username, err := recuperavariabile("username")
			if err != nil {
				log.Fatal(err)
				return
			}

			//Recupera la variabile d'ambiente
			password, err := recuperavariabile("password")
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

			//Cicla i tempi
			for _, t := range tempi {
				value := fmt.Sprint(dp[t].(float64))
				sec, _ := strconv.ParseInt(t[:10], 10, 64)
				epoch := time.Unix(sec, 0)
				infotemporali := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v",
					epoch.YearDay(), epoch.Year(), epoch.Month(),
					epoch.Day(), epoch.Hour(), epoch.Minute(),
					epoch.Weekday())
				serie = append(serie, []string{device, t, infotemporali, value})
				//fmt.Println("orario: ", t, "valore: ", dp[t])
			}

			return
		}
	}
}
