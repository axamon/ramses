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

	"gonum.org/v1/gonum/stat"
)

const sessioniPPP = "https://ipw.telecomitalia.it/ipwmetrics/api/v1/rawmetrics/kpi.ppoe.slot?device="

var wgppp sync.WaitGroup

func nasppp() {

	var devices []string
	listalistanas := recuperaNAS()

	var i int
	for _, listanas := range listalistanas {
		for _, nas := range listanas {
			//fmt.Println(n, nas.Name)
			if strings.HasPrefix(nas.Service, "NAS") {
				i++
				devices = append(devices, nas.Name)
			}
		}
	}
	log.Printf("I Nas trovati sono %d\n", i)

	recuperaSessioniPPP := func() {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 4*time.Minute)
		defer cancel()
		for _, device := range devices {
			wgppp.Add(1)
			log.Printf("Verifico device %s\n", device)
			go nasppp2(device)
		}

		return
	}

	recuperaSessioniPPP()
	wgppp.Wait()
	fmt.Println("Dopo primo run")

	//imposta un refesh ogni tot minuti
	t := time.Tick(1 * time.Minute)
	c := time.Tick(5 * time.Minute)
	for {
		select {
		case <-c:
			recuperaSessioniPPP()
			wgppp.Wait()
		case <-t:
			fmt.Println(".")
		}
	}

}

func nasppp2(device string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	defer wgppp.Done()
	//Attendo un tempo random per evitare di fare troppe query insieme
	randomdelay := rand.Intn(100)
	time.Sleep(time.Duration(randomdelay) * time.Millisecond)

	os.Setenv("HTTP_PROXY", "")
	os.Setenv("HTTPS_PROXY", "")
	fmt.Println(os.Getenv("HTTP_PROXY"))
	fmt.Println(os.Getenv("HTTPS_PROXY"))

	//fmt.Println(device)

	var sigma float64
	sigma = 2.0

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
		log.Printf("Non ci sono abbastanza info per %s", device)
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
	log.Printf("%s: media: %2.f stdev: %2.f", device, mean, stdev)

	//Passo le info alla fuzione di elaborazione e grafico
	//wg.Add()
	//elaboraseriePPP(serieppptime, seriepppvalue, device, "test", "ppp")

	//Verifica se ci sono errori da segnalare
	for _, v := range seriepppvalue[len(seriepppvalue)-1:] {
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

	return
}
