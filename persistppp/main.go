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
		log.Println("Non trovo il file %s, riscarico la lista NAS", fileNasInventory)
		recuperaNAS()
		log.Println("%s creato", fileNasInventory)

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
