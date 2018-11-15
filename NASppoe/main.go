package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	device := os.Args[1]

	url := "https://ipw.telecomitalia.it/grafana/api/datasources/proxy/1/api/query"

	adesso := time.Now()
	quattroOreFa := adesso.Add(time.Duration(-4) * time.Hour)

	adessoEpoch := fmt.Sprint(adesso.Unix())
	quattroOreFaEpoch := fmt.Sprint(quattroOreFa.Unix())

	payload := strings.NewReader("{\"start\":" + quattroOreFaEpoch + ",\"queries\":[{\"metric\":\"kpi.ppoe.slot\",\"aggregator\":\"sum\",\"downsample\":\"2m-avg\",\"tags\":{\"device\":\"" + device + "\"}}],\"msResolution\":false,\"globalAnnotations\":true,\"end\":" + adessoEpoch + "}\r\n")

	fmt.Println(payload, adesso.Unix(), quattroOreFa.Unix())

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("authorization", "Basic MDAyNDY1MDY6Y2Z4VyRsTVM2ZA==")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	//	fmt.Println(res)
	//	fmt.Println(string(body))
	var result []interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err.Error())
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
	var serieppp []float64
	for _, t := range tempi {
		tindex, _ := strconv.Atoi(t)
		serieppp[tindex] = dp[t].(float64)
		fmt.Println("orario: ", t, "valore: ", dp[t])
	}
	fmt.Println(serieppp)

}
