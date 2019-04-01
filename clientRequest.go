package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func clientRequest(url, username, password, device string) (result []interface{}) {

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
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err.Error(), device)
		return result
	}
	if len(result) < 1 {
		log.Printf("Error %s Non ci sono abbastanza info\n", device)
		err := CreaTrap(device, "No data ppp", "Assenza dati sulle sessioni ppp", listanasip[device], 2, 4)
		if err != nil {
			log.Printf("Error %s Impossibile inviare Trap\n", device)
		}
		nastrappati[device] = true
		return result
	}
	return result

}
