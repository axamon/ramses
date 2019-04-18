package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// wgpp è un waitgroup per sincronizzare le goroutines.
var wgppp sync.WaitGroup

// Creo la mappa dove mettere nas name e ip di management insieme.
var listanasip = make(map[string]string)

// Creo la mappa dei NAS per cui è stata inviata una trap.
var nastrappati = make(map[string]bool)

func nasppp(ctx context.Context, nomiNas []string) {

	// recuperaSessioniPPP è una funzione che recupera i dati ppp dei nas.
	recuperaSessioniPPP := func(ctx context.Context) {
		// Espando il contesto inziale inserendo un timeout.
		ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)

		// Prima di chiudere vengono rilasciate tutte le risorse.
		defer cancel()

		// Avvio un ciclo infinito.
		for {
			select {
			// Se viene raggiunto il timeout la funzione viene killata.
			case <-ctx.Done():
				log.Printf(
					"Error All Tempo per completare recupero dati terminato\n")
				return

			// Finchè non si raggiunge il timeout
			// viene eseguito il codice di default.
			default:

				for _, device := range nomiNas {
					wgppp.Add(1)
					// log.Printf("%s Info Inzio verifica device\n", device)
					// go nasppp2(device)
					nasppp2(ctx, device)
				}
				return
			}
		}

	}

	// Prima esecuzione del recupero dati dall'avvio dell'applicazione.
	recuperaSessioniPPP(ctx)

	// Attende che tutte le richieste siano terminate prima di proseguire.
	wgppp.Wait()

	// Verifica l'avvio delle mail.
	// Se non riesce a mandare mail l'applicativo esce con errore.
	err := mandamail(
		configuration.SmtpFrom,
		configuration.SmtpTo,
		"Avvio",
		eventi)

	if err != nil {
		log.Printf(
			"Error Impossibile inviare mail: %s\n",
			err.Error())
		os.Exit(1)
	}

	log.Printf("INFO Mail di avvio inviata.\n")

	fmt.Println("Dopo primo run") //debug

	// c imposta l'intervallo di verifica delle sessioni nas.
	c := time.Tick(5 * time.Minute)

	// update imposta l'intervallo di inoltro della mail di notifica
	// per far sapere che Ramses è ancora attivo.
	update := time.Tick(24 * time.Hour)

	for {
		select {
		case <-update:
			// Ogni tot specificato dal canale update invia mail
			// per far sapere che il sistema è attivo.
			mandamail(
				configuration.SmtpFrom,
				configuration.SmtpTo,
				"Update",
				eventi)
		case <-c:
			// Ogni tot specificato dal canale c fa partire
			// il recupero dei dati di sessione PPP.
			recuperaSessioniPPP(ctx)
			wgppp.Wait()
		}
	}

}

func nasppp2(ctx context.Context, device string) {

	// Riceve il contesto padre e aggiunge un timeout.
	// massimo per terminare la richiesta dati.
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)

	// Esegue cancel a fine procedura.
	defer cancel()

	//defer log.Printf("%s Info Recupero dati terminato\n", device)
	defer wgppp.Done()

	for {
		select {
		case <-ctx.Done():
			log.Printf(
				"Error %s Superato tempo massimo per raccolta dati\n",
				device)
			return

		default:
			// Attendo un tempo random per evitare troppe query insieme.
			// se sono attive le goroutines.
			randomdelay := rand.Intn(100)
			time.Sleep(
				time.Duration(randomdelay) * time.Millisecond)

			// Ripulisce eventiali impostazioni di proxy
			// a livello di sistema.
			os.Setenv("HTTP_PROXY", "")
			os.Setenv("HTTPS_PROXY", "")
			fmt.Println(os.Getenv("HTTP_PROXY"))
			fmt.Println(os.Getenv("HTTPS_PROXY"))

			//fmt.Println(device)

			// Recupera le credenziali per IPDOM.
			username := configuration.IPDOMUser
			password := configuration.IPDOMPassword

			// Ricompongo la URL di IPDOM con il nome del NAS all'interno.
			url := configuration.URLSessioniPPP +
				device + configuration.URLTail7d

			// Avvia la richiesta web.
			result := clientRequest(ctx, url, username, password, device)

			// Elabora il risulatato della richiesta web.
			elaboroRequest(ctx, result, device)

			return
		}
	}
}
