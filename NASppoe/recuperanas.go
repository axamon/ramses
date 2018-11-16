package main

import (
	"crypto/tls"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

// TNAS is a structure Type for contain NAS interface
type TNAS struct {
	ID                int
	Name              string
	Description       string
	Location          string
	Service           string
	ServiceAdded      string
	Pop               string
	ManIPAddress      string
	Network           string
	Centrale          string
	Alias             string
	Domain            string
	CreateTime        string
	ChangeTime        string
	SysContact        string
	SysDescr          string
	SysLocation       string
	SysName           string
	SysObject         string
	Aisle             string
	Altitude          string
	ChassisUUID       string
	FRUNumber         string
	FRUSerialNumber   string
	FWRevision        string
	HWRevision        string
	Manufacturer      string
	ManufacturingDate string
	Model             string
	ChassisName       string
	PartNumber        string
	RfidTag           string
	RackPosition      string
	RelativePosition  string
	CdmRow            string
	SwRevision        string
	SerialNumber      string
	SystemBoardUUID   string
	CdmType           string
	UserTracking      string
	XCoordinate       string
	YCoordinate       string
	PhysicalIndex     string
	VendorType        string
	ClassName         string
	Uptime            string
	Services          string
	InterfaceCount    string
	IsIPForwarding    string
	AccessIPAddress   string
	AccessProtocol    string
	DiscoveryTime     string
	Timestamp         string
	Routing           string
	Metrics           string
	Events            string
	Device            string
	Interfaces        string
	Fans              string
	Cards             string
	Slots             string
	Psus              string
	Configurations    string
}

func recuperaNAS() (nasList []TNAS) {

	var bjson []byte

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

	url := "https://ipw.telecomitalia.it/ipwinventory/api/v1/devices/?limit=1000skip=1000&name=^r-rm"
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)

	//Header che forse potrebbero essere tolti ma male non fanno
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	//***************************************************************************
	//------- read from ipdom --------
	//qui su costringe il client ad accettare anche certificati https non validi o scaduti, non anrebbe fatto ma bisogna fare di necessità virtù
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	client := &http.Client{Transport: transCfg}
	res, _ := client.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	err = ioutil.WriteFile("nasInventory.json", body, 0644) //scrive i dati su file json
	defer res.Body.Close()
	check(err)

	// // Scrivi il json su file
	// 	// Create a json file
	// fj, err := os.Create("nasInventory.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer fj.Close()
	// //Write Unmarshaled json data to json file
	// w := csv.NewWriter(fj)

	// ------ read from file -------
	bjson, err = ioutil.ReadFile("nasInventory.json")
	// -------------------
	// *****************************************************************************
	//recupera il risultato della query a ipdom
	//	var d []TNAS

	err = json.Unmarshal(bjson, &nasList)
	//	err = json.Unmarshal([]byte(body), &result)

	if err != nil {
		fmt.Println(err)
	}

	// 	err = ioutil.WriteFile("output.txt", body, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(nasList)
	// for _, myNas := range nasList {

	// 	fmt.Println(myNas)
	// }

	// for _, obj := range result {
	// 	var record []string

	// 	record = append(record, obj["id"])
	// 	record = append(record, get_value(result, "name"))
	// 	record = append(record, get_value(result, "pop"))

	// 	w.Write(record)
	// }
	// w.Flush()

	return
}

// NasInventory2Csv write a nas list to a csv file
func NasInventory2Csv(nasList []TNAS, filename string) {
	// var auxNas TNAS
	// sn := reflect.ValueOf(&auxNas).Elem()
	//	tnasType := sn.Type()
	sep := ";"
	fo, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	//	fo.WriteString("ID,Name,Description\n")
	// for i := 0; i < tnasType.NumField(); i++ {
	// 	fmt.Println(tnasType.Field(i).Name)
	// }
	inmezzo := "\"" + sep + "\""
	isFirst := true
	for _, myNas := range nasList {
		s := reflect.ValueOf(&myNas).Elem()
		typeOfT := s.Type()
		if isFirst == true {
			isFirst = false
			for i := 0; i < typeOfT.NumField(); i++ {
				if i != 0 {
					fo.WriteString(sep)
				}
				fo.WriteString(typeOfT.Field(i).Name)
			}
			fo.WriteString("\n")
		}
		fo.WriteString(strconv.Itoa(myNas.ID))
		fo.WriteString(sep + "\"")
		fo.WriteString(myNas.Name)
		// fo.WriteString(inmezzo)
		// fo.WriteString(myNas.Description)

		for i := 2; i < s.NumField(); i++ {
			f := s.Field(i)
			fo.WriteString(inmezzo)
			fo.WriteString(f.String())
			fmt.Printf("%d: %s %s = %v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
		// for _, d:= range []string{ myNas.Location, myNas.Pop, } {
		// 	//do something with the d
		//   }

		fo.WriteString("\"\n")

		//fmt.Println(myNas)
	}

	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
