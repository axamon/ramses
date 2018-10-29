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

func main() {
	ifs := ifNames("xrs-mi001")
	fmt.Println(ifs)
}

func ifNames(device string) (interfacce []string) {

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

	url := "https://ipw.telecomitalia.it/ipwmetrics/api/v1/metrics/net.throughput.out/" + device

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

	var result map[string]interface{}
	errResult := json.Unmarshal([]byte(body), &result)
	if errResult != nil {
		log.Println("errore: ", err.Error())
	}

	//device = "xrs-mi001"
	NET := result["net.throughput.out"].(map[string]interface{})
	DEVICE := NET[device].(map[string]interface{})

	for ifname, _ := range DEVICE {
		//fmt.Println(k, v.(map[string]interface{})["time"], v.(map[string]interface{})["value"])
		interfacce = append(interfacce, ifname)
	}

	return
}
