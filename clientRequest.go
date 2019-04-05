package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func clientRequest(ctx context.Context, url, username, password, device string) (result []interface{}) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error Get new request %s\n", err.Error())
	}

	// Costringe il client ad accettare anche certificati https non validi
	// o scaduti.
	transCfg := &http.Transport{
		// Ignora certificati SSL scaduti.
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.SetBasicAuth(username, password)
	req.Header.Add("cache-control", "no-cache")
	req.WithContext(ctx)

	client := &http.Client{Transport: transCfg}

	res, err := client.Do(req)
	//res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error HTTP Client Do impossibile raggiungere %s: %s\n", device, err.Error())
		return nil
	}

	if res.StatusCode > 300 {
		log.Printf("Error Ricevuto un errore http: %d\n", res.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error impossibile ricevere HTTP body per %s, %s\n", device, err.Error())
		// os.Exit(1)
	}
	defer res.Body.Close()
	// fmt.Println(res)
	// fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Error unmarshal impossibile per %s, %s\n", device, err.Error())
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
