package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/remeh/sizedwaitgroup"
)

const (
	ipdomainurl string = "https://ipw.telecomitalia.it/ipwmetrics/api/v1/metrics/"
)

var metriche = []string{
	"net.volume.in",
	"net.volume.out",
	"net.errors.in",
	"net.errors.out",
	"net.discards.in",
	"net.discards.out",
	"net.throughput.in",
	"net.throughput.out"}

//Dove salvere il nome delle interfacce
var listainterfacce []string

//Waitgroupche gestisce il throtteling
var wg = sizedwaitgroup.New(80)

//var wg waitgroup //waitgroup vecchio stile

//Gestione sigma
var sigma float64

func init() {
	flag.Float64Var(&sigma, "s", 2, "imposta il numero di deviazioni standard da considerare")
}

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Sintassi: -s=<sigma da usare> <device da controllare>")
		return
	}

	var device string

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		device = arg
	}

	fmt.Println(sigma, device)

	log.Printf("Elaborazione per %s Iniziata\n", device)

	//	ifnames := ifNames(device)

	recuperajson(device)

	log.Printf("Elaborazione per %s terminata\n", device)
}
