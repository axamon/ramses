package main

import (
	"strings"

	"github.com/axamon/stringset"
)

func selezionaNas() (nomiNas []string, err error) {

	// Creo la variabile dove accodare i nomi dei nas
	//var devices []string

	// TODO: creare il file con i nomi NAS dinamicamente

	// Identifico il file json con le informazioni da parsare
	//filelistapparati := configuration.NasInventory

	// Leggo il file in memoria
	//body, err := ioutil.ReadFile(filelistapparati)
	//if err != nil {
	//	log.Printf("Error Impossibile recuperare lista %s\n", filelistapparati)
	//}

	// Creo la variabile dove conservare i dati parsati
	//var listalistanas [][]TNAS
	//errjson := json.Unmarshal(body, &listalistanas)
	//if errjson != nil {
	//	log.Printf("Error Impossibile parsare dati %s\n", filelistapparati)
	//}

	// Identifico il file json con la lista NAS da ignorare
	//filelistaNasDaIgnorare := configuration.NasDaIgnorare
	//log.Println(filelistaNasDaIgnorare) //debug

	// Leggo il file in memoria
	//ignoranasbody, errignoranas := ioutil.ReadFile(filelistaNasDaIgnorare)
	//if errignoranas != nil {
	//	log.Printf("Error Impossibile recuperare lista dei nas da ignorare %s %s\n", filelistaNasDaIgnorare, errignoranas.Error())
	//}

	// Creo variabile che contiene lista nas da ignorare
	//var listaNasDaIgnorare map[string][]string
	//errjsonNasdaignorare := json.Unmarshal(ignoranasbody, &listaNasDaIgnorare)
	//if errjsonNasdaignorare != nil {
	//	log.Printf("Error Impossibile parsare dati %s , %s\n", listaNasDaIgnorare, errjsonNasdaignorare.Error())
	//}
	//fmt.Println(listaNasDaIgnorare)

	// Creo set di Nas da ignorare
	//ignoraNasSet := stringset.NewStringSet()

	//var ignora = make(map[string]bool)
	//for _, nasignorato := range listaNasDaIgnorare["nasdaignorare"] {
	//	ignora[nasignorato] = true
	//	ignoraNasSet.Add(nasignorato)
	//}

	// listalistanas è una lista di liste quindi bisogna fare un doppio ciclo for
	for _, listanas := range listalistanas {
		for _, nas := range listanas {
			// fmt.Println(n, nas.Name) //debug

			// Escludo i NAS in da ignorare
			//if _, ok := ignora[nas.Name]; ok {
			//	log.Printf("INFO %s ignorato\n", nas.Name)
			//	continue
			//}

			// Considero solo gli apparati che abbiano "NAS" all'inzio del campo Service
			// e EDGE_BRAS come dominio
			// e MX960 come chassis
			if strings.HasPrefix(nas.Service, "NAS") && strings.Contains(nas.Domain, "EDGE_BRAS") &&
				strings.Contains(nas.ChassisName, "MX960") {

				// Appendo in nomiNAs il nome nas trovato
				nomiNas = append(nomiNas, nas.Name)

				// Per inviare trap serve conoscere l'ip di management del NAS uffa che barba che noia
				listanasip[nas.Name] = nas.ManIPAddress
				// log.Printf("Info %v ignorato\n", devices) //debug

			}
		}
	}

	// Tolgo dal set devices i nas da ignorare e salvo in nomiNasSet
	//nomiNasSet = devices.Difference(ignoraNasSet)
	nomiNasSet := stringset.NewStringSet()
	for _, nome := range nomiNas {
		nomiNasSet.Add(nome)
	}

	listaNomiNas := nomiNasSet.Strings()

	return listaNomiNas, nil

}
