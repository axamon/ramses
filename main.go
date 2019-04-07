package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
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
var listalistanas [][]TNAS


var version = "version: 4.1"

func main() {
	// Creo il contesto inziale che verrà propagato alle go-routine
	// con la funzione cancel per uscire dal programma in modo pulito.
	ctx, cancel := context.WithCancel(context.Background())

	// Creo il canale c di buffer 1 per gestire i segnali di tipo CTRT+
	c := make(chan os.Signal, 1)
	// Notifica al canale c i tipi di segnali di interrupt.
	signal.Notify(c, os.Interrupt)
	
	defer func() {
		signal.Stop(c) // Ipedisce a c di ricevere ulteriori segnalazioni.
		cancel() // Avvia la funzione di chiusura.
	}()
	// Avvia una go-routine in background in ascolto sul canale c.
	go func() {
		select {
		case <-c: // Se arriva qualche segnale in c.
			fmt.Println("Spengo Ramses")
			mandamail(configuration.SmtpFrom, configuration.SmtpTo, "Chiusura", eventi)
			cancel()
			os.Exit(0)
		case <-ctx.Done(): // Se il contesto ctx viene terminato.
		}
	}()

	// Prima di terminare la funzione invia una mail
	
	// Scrive su standard output la versione di Ramses
	log.Printf("Avvio Ramses %s\n", version)

	// Recupera valori dal file di configurazione passato come argomento.
	file := os.Args[1]
	err := gonfig.GetConf(file, &configuration)
	if err != nil {
		log.Printf("Error Impossibile recupere valori da %s: %s\n", file, err.Error())
		os.Exit(1)
	}

	// GatherInfo recupera informazioni di sevizio sul funzionamento dell'APP
	GatherInfo()

	_, err = recuperaNAS(ctx)
	if err != nil {
		log.Printf("Impossibile recuperare NAS da IPDOM %s\n", err.Error())
		// cancel()
		// os.Exit(1)
	}
	log.Printf("INFO Recuperati dati NAS da IPDOM\n")

	nasppp(ctx)

}
