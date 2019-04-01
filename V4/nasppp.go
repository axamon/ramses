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

	"github.com/axamon/stringset"

	"github.com/axamon/ramses/algoritmi"
	"gonum.org/v1/gonum/stat"
)

// wgpp è un waitgroup per sincronizzare le goroutines
var wgppp sync.WaitGroup

// Creo la mappa dove mettere nas name e ip insieme
var listanasip = make(map[string]string)

// Creo la mappa dei NAS per cui è stata inviata una trap
var nastrappati = make(map[string]bool)

func nasppp() {
	// Creo il contesto inziale che verrà propagato alle go-routine
	ctx := context.Background()

	// Prima di terminare la funzione invia una mail
	defer mandamailChiusura(configuration.SmtpFrom, configuration.SmtpTo)

	// Verifica l'avvio di mail. Se non riesce a mandare mail esce.
	err := mandamailAvvio(configuration.SmtpFrom, configuration.SmtpTo)
	if err != nil {
		log.Printf("Error Impossibile inviare mail: %s\n", err.Error())
		os.Exit(1)
	}

	// Creo la variabile dove accodare i nomi dei nas
	var devices *stringset.StringSet

	// TODO: creare il file con i nomi NAS dinamicamente

	// Identifico il file json con le informazioni da parsare
	filelistapparati := configuration.NasInventory

	// Leggo il file in memoria
	body, err := ioutil.ReadFile(filelistapparati)
	if err != nil {
		log.Printf("Error Impossibile recuperare lista %s\n", filelistapparati)
	}

	// Creo la variabile dove conservare i dati parsati
	var listalistanas [][]TNAS
	errjson := json.Unmarshal(body, &listalistanas)
	if errjson != nil {
		log.Printf("Error Impossibile parsare dati %s\n", filelistapparati)
	}

	// Identifico il file json con la lista NAS da ignorare
	filelistaNasDaIgnorare := configuration.NasDaIgnorare
	log.Println(filelistaNasDaIgnorare) //debug

	// Leggo il file in memoria
	ignoranasbody, errignoranas := ioutil.ReadFile(filelistaNasDaIgnorare)
	if errignoranas != nil {
		log.Printf("Error Impossibile recuperare lista %s\n", filelistaNasDaIgnorare)
	}

	// Creo variabile che contiene lista nas da ignorare
	var listaNasDaIgnorare map[string][]string
	errjsonNasdaignorare := json.Unmarshal(ignoranasbody, &listaNasDaIgnorare)
	if errjsonNasdaignorare != nil {
		log.Printf("Error Impossibile parsare dati %s , %s\n", listaNasDaIgnorare, errjsonNasdaignorare.Error())
	}
	fmt.Println(listaNasDaIgnorare)

	// Creo set di Nas da ignorare
	ignoraNasSet := stringset.NewStringSet()

	var ignora = make(map[string]bool)
	for _, nasignorato := range listaNasDaIgnorare["nasdaignorare"] {
		ignora[nasignorato] = true
		ignoraNasSet.Add(nasignorato)
	}

	// Creo set che contiene i nomi dei NAS
	nomiNasSet := stringset.NewStringSet()

	// listalistanas è una lista di liste quindi bisogna fare un doppio ciclo for
	for _, listanas := range listalistanas {
		for _, nas := range listanas {
			// fmt.Println(n, nas.Name) //debug

			// Escludo i NAS in da ignorare
			//if _, ok := ignora[nas.Name]; ok {
			//	log.Printf("INFO %s ignorato\n", nas.Name)
			//	continue
			//}

			//considero solo gli apparati che abbiano "NAS" all'inzio del campo Service
			//e EDGE_BRAS come dominio e MX960 come chassis
			if strings.HasPrefix(nas.Service, "NAS") && strings.Contains(nas.Domain, "EDGE_BRAS") && strings.Contains(nas.ChassisName, "MX960") {

				//Appendo in devices il nome nas trovato
				devices.Add(nas.Name)

				//Per inviare trap serve conoscere l'ip di management del NAS uffa che barba che noia
				listanasip[nas.Name] = nas.ManIPAddress
				//log.Printf("Info %v ignorato\n", devices) //debug

			}
		}
	}

	// Tolgo dal set devices i nas da ignorare e salvo in nomiNasSet
	nomiNasSet = devices.Difference(ignoraNasSet)

	// loggo il numero di NAS identificati
	log.Printf("%v INFO numero di NAS trovati\n", nomiNasSet.Len())
	time.Sleep(3 * time.Second)

	// recuperaSessioniPPP è una funzione che recupera i dati ppp dei nas
	recuperaSessioniPPP := func() {
		// Espando il contesto inziale inserendo un timeout
		ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)

		// Prima di chiudere vengono rilasciate tutte le risorse
		defer cancel()

		// Avvio un ciclo infinito
		for {
			select {
			//se viene raggiunto il timeout la funzione viene killata
			case <-ctx.Done():
				log.Printf("Error All Tempo per completare recupero dati terminato\n")
				return

			// finchè non si raggiunge il timeout viene eseguito il codice di default
			default:

				for _, device := range nomiNasSet.Strings() {
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
			// Attendo un tempo random per evitare di fare troppe query insieme
			// se sono attive le goroutines.
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

			// Recupera le credenziali per IPDOM
			username := configuration.IPDOMUser
			password := configuration.IPDOMPassword

			url := configuration.URLSessioniPPP + device + configuration.URLTail7d

			req, _ := http.NewRequest("GET", url, nil)

			// Costringe il client ad accettare anche certificati https non validi
			// o scaduti.
			transCfg := &http.Transport{
				// Ignora certificati SSL scaduti.
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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
			// fmt.Println(res)
			// fmt.Println(string(body))
			var result []interface{}
			err = json.Unmarshal(body, &result)
			if err != nil {
				log.Println(err.Error(), device)
				return
			}
			if len(result) < 1 {
				log.Printf("%s Error Non ci sono abbastanza info\n", device)
				err := Creatrap(device, "No data ppp", "Assenza dati sulle sessioni ppp", listanasip[device], 2, 4)
				if err != nil {
					log.Printf("%s Error Impossibile inviare Trap\n", device)
				}
				nastrappati[device] = true
				return
			}
			d := result[0].(map[string]interface{})
			dp := d["dps"].(map[string]interface{})

			// Mette i timestamps in ordine
			tempi := make([]string, 0)
			for t := range dp {
				tempi = append(tempi, t)
			}
			// Ordina i timestamps in maniera crescente
			sort.Strings(tempi)

			// Crea variabili di appoggio
			var seriepppvalue []float64
			var serieppptime []float64

			// Cicla i tempi
			for _, t := range tempi {
				tint, _ := strconv.Atoi(t)
				serieppptime = append(serieppptime, float64(tint))
				seriepppvalue = append(seriepppvalue, dp[t].(float64))
				//fmt.Println("orario: ", t, "valore: ", dp[t])
			}

			// Se non ci sono abbastanza valori per la serie esce
			if len(seriepppvalue) < 300 {
				log.Printf("%s Error Non ci sono abbastanza dati per elaborare statistiche", device)
				return
			}

			// Mofifica serie prima che sia elaborata

			// Elimino il trend
			xdet, ydet := algoritmi.Detrend(serieppptime, seriepppvalue)

			// Applico Derivata terza
			y, _ := algoritmi.Derive3(ydet)

			// Calcolo statistiche sulla serie elaborata
			mean, stdev := stat.MeanStdDev(y, nil)
			// log.Printf("%s Info media: %2.f stdev: %2.f", device, mean, stdev) // debug

			for i := 10; i < len(y); i++ {
				// Individuo se è avvenuto un Jerk
				if y[i] < mean-sigma*stdev {
					unixtimeUTC := time.Unix(int64(xdet[i]/1000), 0)
					// Serve per avere il timestamp di quando c'è stato il problema
					unixtimeinRFC3339 := unixtimeUTC.Format(time.RFC3339)

					// Devo verificare se valori futuri dopo il Jerk hanno avuto problemi
					numvalori := len(seriepppvalue)
					for l := 0; l <= 6; l++ {

						// Evita che si arrivi alla fine della serie di valori
						if i+l > numvalori-1 {
							break
						}
						// Verifica i valori dopo il jerk
						limite := (seriepppvalue[i] - seriepppvalue[i+l]) / seriepppvalue[i]

						// Se il limite è negativo non ci interessa
						if limite < 0 {
							continue
						}

						//log.Println(seriepppvalue[1-1], seriepppvalue[i], limite)
						//fmt.Printf("%s %s Jerk Ultimovalore: %2.f, Penultimovalore: %2.f, Limite: %.4f, %v\n", unixtimeinRFC3339, device, seriepppvalue[i], seriepppvalue[i+l], limite, l)
						//fmt.Printf("%s Info media: %2.f stdev: %2.f , Penultimovalore: %2.f, Differenza: %2.f\n", device, mean, stdev, seriepppvalue[i-1], seriepppvalue[i]-seriepppvalue[i-1])

						if limite > configuration.Soglia {
							summary := fmt.Sprintf("abbassamento sessioni ppp superiore al %2.0f%%\n", configuration.Soglia*100)
							// Attenzione NON usare log.Print perchè serve printare il timestamp non attuale ma di quando si è verificato il problema
							fmt.Printf("%s Alert %s, %s\n", unixtimeinRFC3339, device, summary)

							// Mandamail di notifica solo se siamo negli ultimi 6 valori
							if i > (numvalori - 6) {
								mandamailAlert(configuration.SmtpFrom, configuration.SmtpTo, device)
								err := Creatrap(device, "sessioni ppp", summary, listanasip[device], 1, 5)
								if err != nil {
									log.Printf("Error %s Impossibile inviare trap\n", device)
								}
								nastrappati[device] = true
							}
						}
					}
				}
			}

			return
		}
	}
}
