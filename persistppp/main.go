package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

var serie [][]string

var fileNasInventory = "nasInventory.json"

func main() {

	if _, err := os.Stat(fileNasInventory); os.IsNotExist(err) {
		log.Printf("Non trovo il file %s, riscarico la lista NAS\n", fileNasInventory)
		recuperaNAS()
		log.Printf("%s creato\n", fileNasInventory)

		time.Sleep(5 * time.Second)
	}

	nasppp()

	file, _ := os.Create("result.csv")
	defer file.Close()

	w := csv.NewWriter(file)
	w.WriteAll(serie) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

}
