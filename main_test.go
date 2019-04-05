package main_test

import (
	"flag"
	"log"
	"os"
	"testing"

	ramses "github.com/axamon/ramses"

	"github.com/tkanos/gonfig"
)

var configuration ramses.Configuration

func TestMain(m *testing.M) {

	// Recupera valori dal file di configurazione passato come argomento
	file := "configurationDev.json"
	err := gonfig.GetConf(file, &configuration)
	if err != nil {
		log.Printf("Error Impossibile recupere valori da %s: %s\n", file, err.Error())
		os.Exit(1)
	}
	flag.Parse()
	exitCode := m.Run()

	os.Exit(exitCode)
}
