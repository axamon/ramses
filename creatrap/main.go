package main

import (
	"fmt"
	"log"

	g "github.com/soniah/gosnmp"
)

/*
	pdu := g.SnmpPDU{
		Name:  "1.3.6.1.2.1.1.6",
		Type:  g.OctetString,
		Value: "Oval Office",
	}
*/

/*
device := "r-al900"
argomento := "sessioni ppp"
severity := 5
summary := "Forte abbassamento sessioni ppp"
*/

func main() {
	result, err := CreaTrap("finto", "sessioni_ppp", "Forte abbassamento sessioni ppp", "10.10.10.10", 1, 5)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, value := range result.Variables {
		fmt.Println(value.Name, value.Type, value.Value)
	}
}

// CreaTrap invia trap snmp v1 per notificare gli eventi
func CreaTrap(device, argomento, summary, ipdevice string, specific, severity int) (result *g.SnmpPacket, err error) {

	// In tutti gli altri casi si può inviare la trap

	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = "127.0.0.1"
	g.Default.Port = 162
	g.Default.Version = g.Version1
	g.Default.Community = "public"
	//g.Default.Logger = log.New(os.Stdout, "Ramses ", 0)

	err = g.Default.Connect()
	if err != nil {
		log.Printf("Error Connect() err: %s\n", err.Error())
	}
	defer g.Default.Conn.Close()

	pdu1 := g.SnmpPDU{
		Name:  "1.3.6.1.4.1.8888.200.34.34.1.1",
		Type:  g.OctetString,
		Value: "RAMSES",
	}

	pdu2 := g.SnmpPDU{
		Name:  "1.3.6.1.4.1.8888.200.34.34.1.2",
		Type:  g.OctetString,
		Value: device,
	}

	pdu3 := g.SnmpPDU{
		Name:  "1.3.6.1.4.1.8888.200.34.34.1.3",
		Type:  g.OctetString,
		Value: argomento,
	}

	pdu4 := g.SnmpPDU{
		Name:  "1.3.6.1.4.1.8888.200.34.34.1.4",
		Type:  g.Integer,
		Value: severity,
	}

	pdu5 := g.SnmpPDU{
		Name:  "1.3.6.1.4.1.8888.200.34.34.1.5",
		Type:  g.OctetString,
		Value: summary,
	}

	trap := g.SnmpTrap{
		Variables:    []g.SnmpPDU{pdu1, pdu2, pdu3, pdu4, pdu5},
		Enterprise:   ".1.3.6.1.4.1.8888.200.34.34.34",
		AgentAddress: ipdevice,
		GenericTrap:  6,
		SpecificTrap: specific,
		Timestamp:    300,
	}

	result, err = g.Default.SendTrap(trap)
	if err != nil {
		log.Printf("Error SendTrap() err: %s\n", err.Error())
	}
	for _, value := range trap.Variables {
		fmt.Println(value.Name, value.Value)
	}
	if result.Error != 0 {
		return nil, fmt.Errorf("Errore nell'invio della trap")
	}
	fmt.Println(result.Error)
	return result, err
}
