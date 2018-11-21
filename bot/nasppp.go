package main

import (
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
	"sync"
	"time"

	"gonum.org/v1/gonum/stat"
)

const sessioniPPP = "https://ipw.telecomitalia.it/ipwmetrics/api/v1/rawmetrics/kpi.ppoe.slot?device="

var wgppp sync.WaitGroup

func nasppp() {

	//elenco NAS da monitorare
	//devices := []string{"r-rm899", "r-rm900"}
	//devices := recuperaNAS()

	devices := []string{"r-al899", "r-al900", "r-an899", "r-an900", "r-ba900", "r-bg900", "r-bo900", "r-bs122", "r-bs900", "r-bs899", "r-bz900", "r-ca900", "r-co900", "r-ct899", "r-ct900", "r-cz900", "r-FI899-re0", "r-ge900", "r-mi506", "r-mi890", "r-mi895", "r-mi898", "r-mi899", "r-mi900", "r-mo898", "r-mo899", "r-na899", "r-na900", "r-nl897", "r-nl899", "r-pa900", "r-pe899", "r-pd166", "r-pd900", "r-pe900", "r-pg900", "r-pi095", "r-pi899", "r-pi900", "r-rm613", "r-rm890", "r-rm897", "r-rm898", "r-rm895", "r-rm899", "r-rm900", "r-rn899", "r-rn900", "r-SV900", "r-ta899", "r-ta900", "r-to189", "r-to900", "r-ts900", "r-vr900", "r-fi900", "r-nl900", "r-ve900", "r-BG899", "r-mi897", "r-nl898", "r-pa899-re0", "r-rm896", "r-to899", "r-ve899", "r-bo890", "r-nl890", "r-fi898", "r-cz899", "r-ct898", "r-mo900", "r-na898", "r-ba899", "r-bo899", "r-mi896", "r-ts899", "r-pd899"}

	for n, device := range devices {

		fmt.Println(n, device)
	}

	for _, device := range devices {
		wgppp.Add(1)
		fmt.Printf("%s, verifico device %s\n", time.Now().Format("20060102T15:04:05"), device)
		//time.Sleep(200 * time.Millisecond)
		go nasppp2(device)
	}
	wgppp.Wait()

	//imposta un refesh ogni tot minuti
	c := time.Tick(15 * time.Minute)
	for now := range c {
		for _, device := range devices {
			wgppp.Add(1)
			fmt.Printf("%s, verifico device %s\n", now.Format("20060102T15:04:05"), device)
			//time.Sleep(200 * time.Millisecond)
			nasppp2(device)
		}
	}

}

func nasppp2(device string) {
	defer wgppp.Done()
	//Attendo un tempo random per evitare di fare troppe query insieme
	randomdelay := rand.Intn(500)
	time.Sleep(time.Duration(randomdelay) * time.Millisecond)

	os.Setenv("HTTP_PROXY", "")
	os.Setenv("HTTPS_PROXY", "")
	fmt.Println(os.Getenv("HTTP_PROXY"))
	fmt.Println(os.Getenv("HTTPS_PROXY"))

	fmt.Println(device)

	var sigma float64
	sigma = 2.5

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
		fmt.Println(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	//fmt.Println(res)
	//	fmt.Println(string(body))
	var result []interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(result) < 1 {
		log.Println("Non ci sono abbastanza info")
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
	var seriepppvalue []float64
	var serieppptime []float64
	for _, t := range tempi {
		tint, _ := strconv.Atoi(t)
		serieppptime = append(serieppptime, float64(tint))
		seriepppvalue = append(seriepppvalue, dp[t].(float64))
		//fmt.Println("orario: ", t, "valore: ", dp[t])
	}
	mean, stdev := stat.MeanStdDev(seriepppvalue, nil)
	fmt.Println(mean, stdev)
	wg.Add()

	//Passo le info alla fuzione di elaborazione e grafico
	elaboraseriePPP(serieppptime, seriepppvalue, device, "test", "ppp")

	//Verifica se ci sono errori da segnalare
	for _, v := range seriepppvalue[len(seriepppvalue)-3:] {
		//fmt.Println(v)
		//Se le sessioni salgono non è importante
		// if v > mean+sigma*stdev {
		// 	log.Printf("Alert su %s, forte innalzamento sessioni ppp\n", device)
		// 	grafanaurl := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
		// 	msg <- fmt.Sprintf("Alert su %s, forte innalzamento sessioni ppp, %s\n", device, grafanaurl)
		// }
		if v < mean-sigma*stdev {
			log.Printf("Alert su %s, forte abbassamento sessioni ppp\n", device)
			grafanaurl := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
			msg <- fmt.Sprintf("Alert su %s, forte abbassamento sessioni ppp, %s\n", device, grafanaurl)
		}
	}

}
