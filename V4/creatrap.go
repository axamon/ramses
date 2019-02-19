package main

import (
	"log"
	"os"

	g "github.com/soniah/gosnmp"
)

/*
device := "r-al900"
argomento := "sessioni ppp"
severity := 5
summary := "Forte abbassamento sessioni ppp"
*/

//Creatrap invia trap snmp v1 per notificare gli eventi
func Creatrap(device, argomento, summary string, severity int) (err error) {

	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = configuration.IPDOMSnmpReceiver
	g.Default.Port = 162
	g.Default.Version = g.Version1
	g.Default.Community = "public"
	g.Default.Logger = log.New(os.Stdout, "", 0)

	err = g.Default.Connect()
	if err != nil {
		log.Printf("Connect() err: %v\n", err)
	}
	defer g.Default.Conn.Close()

	/*
		pdu := g.SnmpPDU{
			Name:  "1.3.6.1.2.1.1.6",
			Type:  g.OctetString,
			Value: "Oval Office",
		}
	*/

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
		AgentAddress: "127.0.0.1", //TODO: cambiare con ip apparato
		GenericTrap:  6,
		SpecificTrap: 0,
		Timestamp:    300,
	}

	_, err = g.Default.SendTrap(trap)
	if err != nil {
		log.Printf("SendTrap() err: %v\n", err)
	}

	return err
}
