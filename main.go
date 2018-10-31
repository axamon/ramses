package main

import (
	"log"
	"os"

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
var wg = sizedwaitgroup.New(20)

//var wg waitgroup //waitgroup vecchio stile

func main() {

	device := os.Args[1]

	//	ifnames := ifNames(device)

	recuperajson(device)

	log.Printf("Elaborazione per %s terminata\n", device)
}
