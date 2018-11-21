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

const urlricerca = "https://ipw.telecomitalia.it/ipwinventory/api/v1/devices/?limit=1000skip=1000&name="

func recuperaNAS() (nasList [][]TNAS) {

	//var bjson []byte

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

	sigle := []string{"AG", "AL", "AN", "AO", "AR", "AP", "AT", "AV", "BA", "BT", "BL", "BN", "BG", "BI", "BO", "BZ", "BS", "BR", "CA", "CL", "CB", "CI", "CE", "CT", "CZ", "CH", "CO", "CS", "CR", "KR", "CN", "EN", "FM", "FE", "FI", "FG", "FC", "FR", "GE", "GO", "GR", "IM", "IS", "SP", "AQ", "LT", "LE", "LC", "LI", "LO", "LU", "MC", "MN", "MS", "MT", "ME", "MI", "MO", "MB", "NA", "NO", "NU", "OT", "OR", "PD", "PA", "PR", "PV", "PG", "PU", "PE", "PC", "PI", "PT", "PN", "PZ", "PO", "RG", "RA", "RC", "RE", "RI", "RN", "RM", "RO", "SA", "VS", "SS", "SV", "SI", "SR", "SO", "TA", "TE", "TR", "TO", "OG", "TP", "TN", "TV", "TS", "UD", "VA", "VE", "VB", "VC", "VR", "VV", "VI", "VT"}

	for _, sigla := range sigle {
		//creo un contenitore per il nuovo NAS
		var nasholder []TNAS

		nas := "^r-" + sigla
		url := urlricerca + nas
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
		//err = ioutil.WriteFile(sigla+"nasInventory.json", body, 0644) //scrive i dati su file json
		defer res.Body.Close()
		//check(err)

		// ------ read from file -------
		//bjson, err = ioutil.ReadFile("nasInventory.json")
		// -------------------
		// *****************************************************************************
		//recupera il risultato della query a ipdom
		//	var d []TNAS

		err = json.Unmarshal(body, &nasholder)
		if err != nil {
			fmt.Println(err)
		}

		nasList = append(nasList, nasholder)

	}
	inventory := &nasList
	b, err := json.Marshal(inventory)
	if err != nil {
		log.Println(err.Error())
	}
	err = ioutil.WriteFile("nasInventory.json", b, 0644) //scrive i dati su file json
	if err != nil {
		log.Println(err.Error())
	}
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
