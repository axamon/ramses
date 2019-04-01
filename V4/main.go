package main

import (
	"log"
	"os"
	"time"

	"github.com/remeh/sizedwaitgroup"
	"github.com/tkanos/gonfig"
)

// wg è un Waitgroup che gestisce il throtteling
var wg = sizedwaitgroup.New(80)

// Canale per invio messaggi
var msg = make(chan string, 1)

// Obsoleto canale per salvare grafici
var image = make(chan string, 1)

// RiceviResult riceve una stringa e la invia a telegram //OBSOLETO
func RiceviResult(result string) {
	msg <- result
	return
}

// Crea variabile con le configurazioni del file passato come argomento
var configuration Configuration

// Crea delle mappe a tempo per storicizzare avventimenti
var antistorm = NewTTLMap(24 * time.Hour)
var violazioni = NewTTLMap(24 * time.Hour)
var nientedatippp = NewTTLMap(12 * time.Hour)

func main() {

	// Scrive su standard output la versione di Ramses
	log.Println("Avvio Ramses version: 4")

	// Recupera valori dal file di configurazione passato come argomento
	err := gonfig.GetConf(os.Args[1], &configuration)
	if err != nil {
		log.Printf("errore: %s", err.Error())
		os.Exit(1)
	}

	// GatherInfo recupera informazioni di sevizio sul funzionamento dell'APP
	GatherInfo()

	nasppp()

}
