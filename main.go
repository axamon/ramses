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

// Versione attuale di Ramses.
var version = "version: 4.5"

// wg è un Waitgroup che gestisce quante richieste contemporanee fare a IPDOM
var wg = sizedwaitgroup.New(5)

// Crea variabile con le configurazioni del file passato come argomento
var configuration Configuration

// Crea delle mappe a tempo per storicizzare avvenimenti
var antistorm = NewTTLMap(24 * time.Hour)
var violazioni = NewTTLMap(24 * time.Hour)
var nientedatippp = NewTTLMap(12 * time.Hour)
var listalistanas [][]TNAS

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
		cancel()       // Avvia la funzione di chiusura.
	}()

	// Avvia una go-routine in background in ascolto sul canale c.
	go func() {
		select {
		case <-c: // Se arriva qualche segnale in c.
			fmt.Println()
			fmt.Println("Spengo Ramses, docilmente...")
			log.Println("INFO Invio mail per comunicare spegnimento Ramses")
			// Prima di terminare la funzione invia una mail
			mandamail(configuration.SmtpFrom, configuration.SmtpTo, "Chiusura", eventi)
			cancel()
			os.Exit(0)
		case <-ctx.Done(): // Se il contesto ctx viene terminato.
		}
	}()

	// Scrive su standard output la versione di Ramses.
	log.Printf("Avvio Ramses %s\n", version)

	// Recupera valori dal file di configurazione passato come argomento.
	file := os.Args[1]
	err := gonfig.GetConf(file, &configuration)
	// Se ci sono errori nel recuperare le informazioni l'applicazione viene chiusa.
	if err != nil {
		log.Printf("Error Impossibile recupere valori da %s: %s\n", file, err.Error())
		os.Exit(1)
	}

	// GatherInfo recupera informazioni di sevizio sul funzionamento dell'APP
	GatherInfo(ctx)

	log.Printf("INFO Inizio recupero informazioni NAS su IPDOM\n")

	listalistanas, err = recuperaNAS(ctx)
	// Se ci sono errori nel recuperare le informazioni l'applicazione viene chiusa.
	if err != nil {
		log.Printf("Impossibile recuperare NAS da IPDOM %s\n", err.Error())
		cancel()
		os.Exit(1)
	}
	if err == nil {
		log.Printf("INFO Recuperati dati NAS da IPDOM\n")
	}

	// Dalla lista NAS seleziona quelli da considerare.
	nomiNas, err := selezionaNas()
	if err != nil {
		log.Println(err.Error())
	}

	// Loggo il numero di NAS identificati
	log.Printf("INFO numero di NAS trovati: %d\n", len(nomiNas))
	time.Sleep(3 * time.Second)

	nasppp(ctx, nomiNas)

}
