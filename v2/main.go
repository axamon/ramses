package main

import (
	"log"
	"os"
	"time"

	"github.com/axamon/ramses/funzioni"
)

var msg = make(chan string, 1)
var image = make(chan string, 1)

var serie [][]string

var fileNasInventory = "nasInventory.json"

func main() {

	if _, err := os.Stat(fileNasInventory); os.IsNotExist(err) {
		log.Printf("Non trovo il file %s, riscarico la lista NAS\n", fileNasInventory)
		funzioni.RecuperaNAS()
		log.Printf("%s creato\n", fileNasInventory)

		time.Sleep(5 * time.Second)
	}

	var alert string
	for {
		select {
		case alert = <-funzioni.Nasppp():
			funzioni.Mandamail("alberto.bregliano@gmail.com", "alberto.bregliano@protonmail.com", alert)
		}
	}

}
