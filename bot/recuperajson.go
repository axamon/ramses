package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/remeh/sizedwaitgroup"
)

const (
	ipdomainurl string = "https://ipw.telecomitalia.it/ipwmetrics/api/v1/metrics/"
)

var metriche = []string{
	"net.volume.in",
	"net.volume.out",
	"net.errors.in",
	"net.errors.out",
	"net.discards.in",
	"net.discards.out",
	"net.throughput.in",
	"net.throughput.out"}

//Waitgroupche gestisce il throtteling
var wg = sizedwaitgroup.New(80)

//var wg waitgroup //waitgroup vecchio stile

//Gestione sigma
var sigma = float64(2)

func recuperajson(device string) (err error) {
	log.Printf("Elaborazione per %s Iniziata\n", device)
	msg <- "Inizio controllo"
	defer log.Printf("Elaborazione per %s Terminata\n", device)

	//Dove salvere il nome delle interfacce
	var listainterfacce []string

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

	for _, metrica := range metriche {
		url := ipdomainurl + metrica + "/" + device
		//fmt.Println(url)

		//url := ipdomainurl + device
		file := device + "_" + metrica + ".json"

		req, _ := http.NewRequest("GET", url, nil)

		//Se il sito richiede di passare una username e password questi sono i campi giusti da cambiare
		req.SetBasicAuth(username, password)

		//Header che forse potrebbero essere tolti ma male non fanno
		req.Header.Add("content-type", "application/json")
		req.Header.Add("cache-control", "no-cache")

		//qui su costringe il client ad accettare anche certificati https non validi o scaduti, non anrebbe fatto ma bisogna fare di necessità virtù
		transCfg := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}

		client := &http.Client{Transport: transCfg}

		res, _ := client.Do(req)

		fmt.Println(res.StatusCode)

		if res.StatusCode > 399 {
			err = fmt.Errorf("Impossibile raggiungere %s", url)
			msg <- "http status errato, controllo interrotto"
			return
		}

		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		//recupera il risultato della query a ipdom
		var result map[string]interface{}
		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			log.Println("errore: ", err.Error())
		}

		NET := result[metrica].(map[string]interface{})
		DEVICE := NET[device].(map[string]interface{})

		//Ammucchiamo tutte le interfacce nella variabile listainterfacce
		for ifname := range DEVICE {
			//fmt.Println(k, v.(map[string]interface{})["time"], v.(map[string]interface{})["value"])
			listainterfacce = append(listainterfacce, ifname)
		}

		//Prendo una interfaccia alla volta ed eseguo il for
		//reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
		for _, ifname := range listainterfacce {

			//Ripulisco la variabile values per ingestare i nuovi valori della nuova interfaccia
			var values []float64
			INT := DEVICE[ifname].(map[string]interface{})
			DATA := INT["data"].([]interface{})

			//Estraggo i valori per ogni interfaccia
			for _, v := range DATA {
				//fmt.Println(k, v.(map[string]interface{})["time"], v.(map[string]interface{})["value"])
				value := fmt.Sprint(v.(map[string]interface{})["value"])
				val, err := strconv.ParseFloat(value, 64)
				if err != nil {
					//se si fossero valori non numerici così me ne accorgo e non si impanica nulla
					log.Println("value non converitibile in float64", err.Error())
				}

				//appendo a values il nuovo valore
				values = append(values, val)

			}

			wg.Add()
			go elaboraserie(values, device, ifname, metrica)
		}

		//Crea il file dove salvare i dati, se non ci risce impanica tutto ed esce.
		if _, err := os.Stat("jsondb"); os.IsNotExist(err) {
			os.Mkdir("jsondb", 664)
		}
		f, err := os.Create("jsondb" + "/" + file)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.Write(body)
		//wg.Done()
	}
	//Attende che siano finite tutte le elaborazioni prima di chiudere
	wg.Wait()
	msg <- "Finito controllo"

	return
}
