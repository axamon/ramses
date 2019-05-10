package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/axamon/stringset"
)

// Creo set di Nas da ignorare
var ignoraNasSet = stringset.NewStringSet()

func ignoraNas() {
	// Identifico il file json con la lista NAS da ignorare
	filelistaNasDaIgnorare := configuration.NasDaIgnorare
	log.Println(filelistaNasDaIgnorare) //debug

	// Leggo il file in memoria
	ignoranasbody, errignoranas := ioutil.ReadFile(filelistaNasDaIgnorare)
	if errignoranas != nil {
		log.Printf(
			"Error Impossibile recuperare lista dei nas da ignorare dal file: %s Errore: %s\n",
			filelistaNasDaIgnorare,
			errignoranas.Error())
	}

	// Creo variabile che contiene lista nas da ignorare
	var listaNasDaIgnorare map[string][]string
	errjsonNasdaignorare := json.Unmarshal(ignoranasbody, &listaNasDaIgnorare)
	if errjsonNasdaignorare != nil {
		log.Printf(
			"Error Impossibile parsare dati dal file: %s Errore: %s\n",
			listaNasDaIgnorare,
			errjsonNasdaignorare.Error())
	}

	var ignora = make(map[string]bool)
	for _, nasignorato := range listaNasDaIgnorare["nasdaignorare"] {
		ignora[nasignorato] = true
		ignoraNasSet.Add(nasignorato)
	}
}
