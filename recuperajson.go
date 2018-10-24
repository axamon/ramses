package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func recuperavariabile(variabile string) (result string, err error) {
	if result, ok := os.LookupEnv(variabile); ok && len(result) != 0 {
		return result, nil
	}
	err = fmt.Errorf("la variabile %s non esiste o Ã¨ vuota", variabile)
	fmt.Fprintln(os.Stderr, err.Error())
	return
}

//Perchè la concorrency non si incasini serve un bel waitgroup
//var wg sync.WaitGroup

func recuperajson(device string) (values TDATA) {

	//Recupera la variabile d'ambiente
	username, err := recuperavariabile("username")
	if err != nil {
		log.Fatal(err)
		return
	}

	password, err := recuperavariabile("password")
	if err != nil {
		log.Fatal(err)
		return
	}
	//device = xrs-mi001
	//for {
	//Scaricare in parallelo i 3 file xml che float64eressano
	//wg.Add(1)
	//getxml(, device+".json")
	//wg.Wait()
	//time.Sleep(5 * time.Minute)
	//	}
	//}
	url := "https://ipw.telecomitalia.it/ipwmetrics/api/v1/metrics/net.throughput.out/" + device
	file := device + ".json"

	//getxml scarica il file in "url" e lo salva in locale con nome "file"
	//func getxml(url string, file string)  {

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
	var test XrsMi001Stru
	if err := json.Unmarshal([]byte(body), &test); err != nil {
		log.Fatal(err.Error())
	}

	values = test.NetThroughputOut.XrsMi001.Five11100GigEthernetICRC0023878MetropolitanoHNE500CMICLD50SEABONELAG100.Data

	// tlen := len(t) - 1
	// for i := 0; i <= tlen; i++ {
	// 	values[i] = t[i].Value
	// 	//fmt.Println(t[i].Value)
	// }
	// //fmt.Println(tlen)

	//Crea il file dove salvare i dati, se non ci risce impanica tutto ed esce.
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(body)
	//wg.Done()

	return values
}
