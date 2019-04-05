package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/remeh/sizedwaitgroup"
	"github.com/tkanos/gonfig"
)

// wg è un Waitgroup che gestisce il throtteling
var wg = sizedwaitgroup.New(80)

// Crea variabile con le configurazioni del file passato come argomento
var configuration Configuration

// Crea delle mappe a tempo per storicizzare avventimenti
var antistorm = NewTTLMap(24 * time.Hour)
var violazioni = NewTTLMap(24 * time.Hour)
var nientedatippp = NewTTLMap(12 * time.Hour)

var version = "version: 4"

func main() {
	// Creo il contesto inziale che verrà propagato alle go-routine
	ctx := context.Background()

	// Prima di terminare la funzione invia una mail
	defer mandamail(configuration.SmtpFrom, configuration.SmtpTo, "Chiusura", eventi)

	// Scrive su standard output la versione di Ramses
	log.Printf("Avvio Ramses %s\n", version)

	// Recupera valori dal file di configurazione passato come argomento
	file := os.Args[1]
	err := gonfig.GetConf(file, &configuration)
	if err != nil {
		log.Printf("Error Impossibile recupere valori da %s: %s\n", file, err.Error())
		os.Exit(1)
	}

	// GatherInfo recupera informazioni di sevizio sul funzionamento dell'APP
	GatherInfo()

	nasppp(ctx)

}
