package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	gomail "gopkg.in/gomail.v2"
)

func mandamailAlert(from, to, device string, evento *Jerk) {

	// Verifica che device non sia nella mappa antistorm, se c'è esce
	elements := antistorm.GetAll()
	for el := range elements {
		if el == device {
			log.Printf("Error %s Segnalazione già inviata recentemente.\n", device)
			return
		}
	}

	// Se device non è nella mappa antistorm allora lo inserisce
	antistorm.AddWithTTL(device, true, 30*time.Minute)

	// Setta l'oggetto della mail
	subject := "Allarme ppp su " + device

	// Crea il contenuto della mail
	grafanaurl := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
	body := fmt.Sprintf("Alert %s Forte abbassamento sessioni ppp, valore riscontrato %d alle %v %s\n", device, int(evento.pppValue), evento.Timestamp.UTC(), grafanaurl)

	// Aggiunge i destinatari in to
	tomultiplo := strings.Split(to, ",")

	t := make(map[string][]string)

	t["To"] = tomultiplo

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	//m.SetHeader("To", to)
	m.SetHeaders(t)
	m.SetAddressHeader("Cc", "alberto.bregliano@telecomitalia.it", "Alberto Bregliano")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewPlainDialer(configuration.SmtpServer, configuration.SmtpPort, configuration.SmtpUser, configuration.SmtpPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("impossibile inviare mail %s\n", err.Error())
	}
}

func mandamail(from, to, scopo string, eventi Jerks) (err error) {

	var listaeventi []string
	for _, evento := range eventi {
		singoloevento := fmt.Sprintf(`
		%s %s %d\n
		`, evento.Timestamp.UTC().Format("20060102T15:04"), evento.NasName, int(evento.pppValue))
		listaeventi = append(listaeventi, singoloevento)
	}

	var subject, body string

	switch scopo {
	case "Avvio":
		subject = "Ramses - Avvio applicazione"
		body = fmt.Sprintf("Ramses avviato\n%v\n", listaeventi)
	case "Update":
		subject = "Ramses - applicazione attiva"
		body = fmt.Sprintf("Ramses è ancora attivo\n%v\n", listaeventi)
	case "Chiusura":
		subject = "Ramses - Arresto applicazione"
		body = fmt.Sprintf("Ramses arrestato\n%v\n", listaeventi)
	}

	tomultiplo := strings.Split(to, ",")

	t := make(map[string][]string)

	t["To"] = tomultiplo

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	//m.SetHeader("To", to)
	m.SetHeaders(t)
	m.SetAddressHeader("Cc", "alberto.bregliano@telecomitalia.it", "Alberto Bregliano")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(configuration.SmtpServer, configuration.SmtpPort, configuration.SmtpUser, configuration.SmtpPassword)

	errdialandsend := d.DialAndSend(m)
	if errdialandsend != nil {
		err = fmt.Errorf("Error Impossibile inviare mail %s", errdialandsend.Error())
	}
	return err
}
