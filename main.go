package main

import (
	"log"
	"os"

	"github.com/remeh/sizedwaitgroup"
)

const (
	ipdomainurl string = "https://ipw.telecomitalia.it/ipwmetrics/api/v1/metrics/net.throughput.out/"
)

//Dove salvere il nome delle interfacce
var listainterfacce []string

//Waitgroupche gestisce il throtteling
var wg = sizedwaitgroup.New(8)

//var wg waitgroup //waitgroup vecchio stile

func main() {

	device := os.Args[1]

	//	ifnames := ifNames(device)

	recuperajson(device)

	log.Printf("Elaborazione per %s terminata\n", device)
}
