package main

import (
	"log"
	"os"
	"time"

	"github.com/remeh/sizedwaitgroup"
	"github.com/tkanos/gonfig"
)

//Waitgroupche gestisce il throtteling
var wg = sizedwaitgroup.New(80)

var msg = make(chan string, 1)
var image = make(chan string, 1)

//RiceviResult riceve una stringa e la invia a telegram
func RiceviResult(result string) {
	msg <- result
	return
}

//crea variabile con le configurazioni del file passato come argomento
var configuration Configuration

//Crea delle mappe a tempo della durata di 24 ore per storicizzare avventimenti
var antistorm = NewTTLMap(24 * time.Hour)
var violazioni = NewTTLMap(24 * time.Hour)
var nientedatippp = NewTTLMap(12 * time.Hour)

//recupera la soglia percentuale di allarmi per cui allarmarsi
//var soglia = configuration.Soglia //è commentato altrimenti non viene preso il valore dopo nel codice

func main() {

	//recupera valori dal file di configurazione passato come argomento
	err := gonfig.GetConf(os.Args[1], &configuration)
	if err != nil {
		log.Printf("errore: %s", err.Error())
		os.Exit(1)
	}

	GatherInfo()

	nasppp()

}
