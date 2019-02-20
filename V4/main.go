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

var configuration Configuration

var antistorm = NewTTLMap(24 * time.Hour)
var violazioni = NewTTLMap(24 * time.Hour)

//recupera la soglia percentuale di allarmi per cui allarmarsi
//var soglia = configuration.Soglia //è commentato se non viene preso il valore dopo nel codice

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
