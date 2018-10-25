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
	err = fmt.Errorf("la variabile %s non esiste o Ã¨ vuota", variabile)
	fmt.Fprintln(os.Stderr, err.Error())
	return
}

//Perchè la concorrency non si incasini serve un bel waitgroup
//var wg sync.WaitGroup

var test XrsMi001Stru

func recuperajson(device, choisedinterface string) (values []float64) {

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

	// if err := json.Unmarshal([]byte(body), &test); err != nil {
	// 	log.Fatal(err.Error())
	// }

	//choisedinterface = "Lag99LAGGroup00000000LOGICORMI595Ae5OFFRAMPToRMI595200G"
	// jsonstring := "test.NetThroughputOut." + device + "." + choisedinterface + ".Data"
	// fmt.Println(jsonstring)
	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		log.Println("errore: ", err.Error())
	}

	device = "xrs-mi001"
	NET := result["net.throughput.out"].(map[string]interface{})
	DEVICE := NET[device].(map[string]interface{})
	INT := DEVICE[choisedinterface].(map[string]interface{})
	DATA := INT["data"].([]interface{})

	for k, v := range DATA {
		fmt.Println(k, v.(map[string]interface{})["time"], v.(map[string]interface{})["value"])
		value := fmt.Sprint(v.(map[string]interface{})["value"])
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println("value non converitibile in float64", err.Error())
		}

		values = append(values, val)
	}
	//fmt.Println(detail)
	// st := refle.sct.TypeOf(test)
	// field := st.Field(0)
	// v := field.Tag.Get("2/1/2-100-Gig-Ethernet-ICR-C00228/05-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG100-")
	// fmt.Println(v)
	//fmt.Println(result["2/1/2-100-Gig-Ethernet-ICR-C00228/05-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG100-"])
	// e := reflect.TypeOf(&test).Elem()
	// getValues(e, device, choisedinterface)
	// s := structs.New(test)
	// p := s.Field("NetThroughputOut").Field("XrsMi001").Field("A110100EthernetTX").Field("Data").Fields()
	// fmt.Println(p)
	//values = "test.NetThroughputOut." + device + "." + choisedinterface + ".Data"

	//values = structpath
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
