package main

import (
	"log"
	"os"
	"time"

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

// CreaTrap invia trap snmp v1 per notificare gli eventi
func CreaTrap(device, argomento, summary, ipdevice string, specific, severity int) (err error) {

	// Se si tratta di inviare trap per mancanza di dati
	// (specific =1 e severity > 0)
	if specific == 1 && severity > 0 {
		//aggiungo il device alla lista per 8 ore
		nientedatippp.AddWithTTL(device, true, 8*time.Hour)
	}

	adesso := time.Now()

	// Creo la variabile trapMancanoDatiInviata come falsa
	var trapMancanoDatiInviata = false

	// Verifico se il device è nella lista di quelli che non hanno dati
	elements := nientedatippp.GetAll()
	for el := range elements {
		if el == device {
			// Sse è presente cambio la variabile trapMancanoDatiInviata
			// su true vuol dire che è stata inviata una trap di problema
			// nelle 8 ore precedenti.
			trapMancanoDatiInviata = true
		}
	}

	// Se vuoi inviare una trap per mancanza di dati e non sono le 10 di mattina
	// oppure il problema è stato già notificato allora esce.
	if specific == 1 && severity > 0 && adesso.Hour() != 10 && trapMancanoDatiInviata == true {
		return
	}

	// Se invece si deve comunicare la risoluzione di un problema di tipo
	// mancanza dati (specific=1e severity=0) non importa a che ora si risolve
	// e se la variabile trapMancanoDatiInviata è falsa vuol dire che
	// non è stata notificata la trap del problema
	if specific == 1 && severity == 0 && trapMancanoDatiInviata == false {
		// quindi si esce perchè non c'è nessuna trap di risoluzione da inviare
		return
	}

	// In tutti gli altri casi si può inviare la trap

	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = configuration.IPDOMSnmpReceiver
	g.Default.Port = configuration.IPDOMSnmpPort
	g.Default.Version = g.Version1
	g.Default.Community = "public"
	g.Default.Logger = log.New(os.Stdout, "", 0)

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

	_, err = g.Default.SendTrap(trap)
	if err != nil {
		log.Printf("Error SendTrap() err: %s\n", err.Error())
	}

	return err
}
