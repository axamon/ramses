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
)

func recuperavariabile(variabile string) (result string, err error) {
	if result, ok := os.LookupEnv(variabile); ok && len(result) != 0 {
		return result, nil
	}
	err = fmt.Errorf("la variabile %s non esiste o è vuota", variabile)
	fmt.Fprintln(os.Stderr, err.Error())
	return
}

func recuperajson(device string) {

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

		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		//recupera il risultato della query a ipdom
		var result map[string]interface{}
		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			log.Println("errore: ", err.Error())
		}

		//Viene preso in considerazione solo il throughput.out ma si può gestire anche altre metriche
		NET := result[metrica].(map[string]interface{})
		DEVICE := NET[device].(map[string]interface{})

		//Ammucchiamo tutte le interfacce nella variabile listainterfacce
		for ifname := range DEVICE {
			//fmt.Println(k, v.(map[string]interface{})["time"], v.(map[string]interface{})["value"])
			listainterfacce = append(listainterfacce, ifname)
		}

		//Prendo una interfaccia alla volta ed eseguo il for
		for _, ifname := range listainterfacce {
			//log.Printf("Inzio elaborazione %s\n", ifname)

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
	// Attende che siano finite tutte le elaborazioni prima di chiudere
	wg.Wait()

	return
}
