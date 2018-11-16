package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat"
)

func main() {

	//elenco NAS da monitorare
	//devices := []string{"r-rm899", "r-rm900"}
	//devices := recuperaNAS()

	devices := []string{"r-al899", "r-al900", "r-an899", "r-an900", "r-ba900", "r-bg900", "r-bo900", "r-bs122", "r-bs900", "r-bs899", "r-bz900", "r-ca900", "r-co900", "r-ct899", "r-ct900", "r-cz900", "r-FI899-re0", "r-ge900", "r-mi506", "r-mi890", "r-mi895", "r-mi898", "r-mi899", "r-mi900", "r-mo898", "r-mo899", "r-na899", "r-na900", "r-nl897", "r-nl899", "r-pa900", "r-pe899", "r-pd166", "r-pd900", "r-pe900", "r-pg900", "r-pi095", "r-pi899", "r-pi900", "r-rm613", "r-rm890", "r-rm897", "r-rm898", "r-rm895", "r-rm899", "r-rm900", "r-rn899", "r-rn900", "r-SV900", "r-ta899", "r-ta900", "r-to189", "r-to900", "r-ts900", "r-vr900", "r-fi900", "r-nl900", "r-ve900", "r-BG899", "r-mi897", "r-nl898", "r-pa899-re0", "r-rm896", "r-to899", "r-ve899", "r-bo890", "r-nl890", "r-fi898", "r-cz899", "r-ct898", "r-mo900", "r-na898", "r-ba899", "r-bo899", "r-mi896", "r-ts899", "r-pd899"}

	for n, device := range devices {

		fmt.Println(n, device)
	}

	for _, device := range devices {
		fmt.Printf("%v, verifico device %s\n", time.Now(), device)
		//time.Sleep(200 * time.Millisecond)
		nasppp(device)
	}

	//imposta un refesh ogni tot minuti
	c := time.Tick(5 * time.Minute)
	for now := range c {
		for _, device := range devices {
			fmt.Printf("%v, verifico device %s\n", now, device)
			//time.Sleep(200 * time.Millisecond)
			nasppp(device)
		}
	}

}

func nasppp(device string) {

	fmt.Println(device)

	var sigma float64
	sigma = 2.0

	url := "https://ipw.telecomitalia.it/grafana/api/datasources/proxy/1/api/query"

	adesso := time.Now()
	//mezzorafa := adesso.Add(time.Duration(-30) * time.Minute)
	quattroOreFa := adesso.Add(time.Duration(-4) * time.Hour)

	adessoEpoch := fmt.Sprint(adesso.Unix())
	//mezzorafaEpoch := fmt.Sprint(mezzorafa.Unix())
	quattroOreFaEpoch := fmt.Sprint(quattroOreFa.Unix())

	payload := strings.NewReader("{\"start\":" + quattroOreFaEpoch + ",\"queries\":[{\"metric\":\"kpi.ppoe.slot\",\"aggregator\":\"sum\",\"downsample\":\"2m-avg\",\"tags\":{\"device\":\"" + device + "\"}}],\"msResolution\":false,\"globalAnnotations\":true,\"end\":" + adessoEpoch + "}\r\n")

	//fmt.Println(payload, adesso.Unix(), quattroOreFa.Unix()) //debug

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("authorization", "Basic MDAyNDY1MDY6Y2Z4VyRsTVM2ZA==")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
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
	var serieppptime []string
	for _, t := range tempi {
		serieppptime = append(serieppptime, t)
		seriepppvalue = append(seriepppvalue, dp[t].(float64))
		//fmt.Println("orario: ", t, "valore: ", dp[t])
	}
	mean, stdev := stat.MeanStdDev(seriepppvalue, nil)
	fmt.Println(mean, stdev)
	//fmt.Println(serieppp)
	for _, v := range seriepppvalue[len(seriepppvalue)-3:] {
		//fmt.Println(v)
		if v > mean+sigma*stdev {
			fmt.Printf("Alert su %s, forte innalzamento sessioni ppp\n", device)
		}
		if v < mean-sigma*stdev {
			fmt.Printf("Alert su %s, forte abbassamento sessioni ppp\n", device)
		}
	}

}
