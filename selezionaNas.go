package main

import (
	"log"
	"strings"

	"github.com/axamon/stringset"
)

func selezionaNas() (nomiNas []string, err error) {

	// Creo la variabile dove accodare i nomi dei nas
	//var devices []string

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

	// listalistanas è una lista di liste quindi bisogna fare un doppio ciclo for
	for _, listanas := range listalistanas {
		for _, nas := range listanas {
			// fmt.Println(n, nas.Name) //debug

			// Escludo i NAS in da ignorare
			/* if _, ok := ignora[nas.Name]; ok {
			log.Printf("INFO %s ignorato\n", nas.Name)
			continue
			} */

			// Considero solo gli apparati che abbiano
			// "NAS" all'inzio del campo Service
			// e EDGE_BRAS come dominio
			// e MX960 come chassis
			if strings.HasPrefix(nas.Service, "NAS") &&
				strings.Contains(nas.Domain, "EDGE_BRAS") &&
				strings.Contains(nas.ChassisName, "MX960") {

				// Appendo in nomiNAs il nome nas trovato
				nomiNas = append(nomiNas, nas.Name)

				// Per inviare trap serve conoscere
				// l'ip di management del NAS
				listanasip[nas.Name] = nas.ManIPAddress
				// log.Printf("Info %v ignorato\n", devices) //debug

			}
		}
	}

	nomiNasSet := stringset.NewStringSet()

	for _, nome := range nomiNas {
		nomiNasSet.Add(nome)
	}

	// Tolgo dal set devices i nas da ignorare e salvo in nomiNasSet
	ignoraNas()
	log.Printf("INFO Lista NAS ignorati: %s\n", ignoraNasSet.Strings())
	nomiNasSet = nomiNasSet.Difference(ignoraNasSet)

	// listaNomiNas è una slice di stringhe con tutti i nomi nas.
	listaNomiNas := nomiNasSet.Strings()

	return listaNomiNas, nil

}
