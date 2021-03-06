package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sort"
	"time"
)

// recuperaNAS si connette a IPDOM per scaricare TUTTE le info sui nas.
func recuperaNAS(ctx context.Context) (nasList [][]TNAS, err error) {

	//var bjson []byte

	username := configuration.IPDOMUser
	password := configuration.IPDOMPassword

	urlricerca := configuration.IPDOMUrlRicerca

	sigle := []string{"AG", "AL", "AN", "AO", "AR", "AP", "AT", "AV", "BA",
		"BT", "BL", "BN", "BG", "BI", "BO", "BZ", "BS", "BR", "CA", "CL", "CB",
		"CI", "CE", "CT", "CZ", "CH", "CO", "CS", "CR", "KR", "CN", "EN", "FM",
		"FE", "FI", "FG", "FC", "FR", "GE", "GO", "GR", "IM", "IS", "SP", "AQ",
		"LT", "LE", "LC", "LI", "LO", "LU", "MC", "MN", "MS", "MT", "ME", "MI",
		"MO", "MB", "NA", "NO", "NU", "OT", "OR", "PD", "PA", "PR", "PV", "PG",
		"PU", "PE", "PC", "PI", "PT", "PN", "PZ", "PO", "RG", "RA", "RC", "RE",
		"RI", "RN", "RM", "RO", "SA", "VS", "SS", "SV", "SI", "SR", "SO", "TA",
		"TE", "TR", "TO", "OG", "TP", "TN", "TV", "TS", "UD", "VA", "VE", "VB",
		"VC", "VR", "VV", "VI", "VT"}

	// Riordina le sigle in ordine alfabetico.
	sort.Strings(sigle)

	// Cicla le sigle.
	for _, sigla := range sigle {

		// Aggiunge uno al waitgroup wg.
		wg.Add()

		// Avvia la go routine.
		go func(sigla string) {

			// A fine funzione diminuisce di uno il wg.
			defer wg.Done()

			// Avvio conteggio tempo.
			start := time.Now()

			// Creo un contenitore per il nuovo NAS.
			var nasholder []TNAS

			// Attende un secondo per non sovraccaricare IPDOM.
			time.Sleep(1 * time.Second)

			// Crea lan regex da passare nella url.
			nas := "^r-" + sigla

			// log.Printf(
			//	"INFO Inizio recupero informazioni NAS provincia %s\n", sigla)

			// Crea la url definitiva.
			url := urlricerca + nas

			// Crea la web request da inviare.
			req, err := http.NewRequest("GET", url, nil)

			if err != nil {
				log.Printf(
					"Error Impossibile creare web request per provincia %s\n",
					sigla)
			}

			// Configura credenziale accesso web a IPDOM.
			req.SetBasicAuth(username, password)

			// Aggiunge header per gestione json.
			req.Header.Add("content-type", "application/json")

			// Aggiunge header per evitare di recuperare dati obsoleti.
			req.Header.Add("cache-control", "no-cache")

			// Aggiunge contesto alla web request.
			req.WithContext(ctx)

			// Ignora certificati https scaduti o errati.
			transCfg := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}

			client := &http.Client{Transport: transCfg}

			res, err := client.Do(req)
			if err != nil {
				log.Printf(
					"Error Impossibile eseguire il client http: %s",
					err.Error())
				return //nil, err
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Printf(
					"Error Impossibile leggere risposta client http: %s",
					err.Error())
				return //nil, err
			}

			// Scrive i dati su file json.
			// err = ioutil.WriteFile(sigla+"nasInventory.json", body, 0644)

			// Chiude il corpo della response quando la funzione finisce.
			defer res.Body.Close()

			//check(err)

			// ------ read from file -------
			//bjson, err = ioutil.ReadFile("nasInventory.json")
			// -------------------
			// ****************************************************************
			//recupera il risultato della query a ipdom
			//	var d []TNAS

			err = json.Unmarshal(body, &nasholder)
			if err != nil {
				log.Printf(
					"Error Impossibile eseguire unmarshal dei dati per %s: %s",
					sigla,
					err.Error())
				// return nil, err
			}

			nasList = append(nasList, nasholder)

			log.Printf(
				"INFO Finito recupero informazioni NAS provincia %s, tempo impiegato: %v",
				sigla,
				time.Since(start))

			runtime.Gosched()
		}(sigla)
	}
	wg.Wait()

	/* inventory := &nasList
	b, err := json.Marshal(inventory)
	if err != nil {
		log.Println(err.Error())
	}
	err = ioutil.WriteFile("nasInventoryNew.json", b, 0644) //scrive i dati su file json
	if err != nil {
		log.Println(err.Error())
	} */
	return nasList, nil
}

// NasInventory2Csv write a nas list to a csv file
/* func NasInventory2Csv(nasList []TNAS, filename string) {
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
*/
