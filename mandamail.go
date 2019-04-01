package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	gomail "gopkg.in/gomail.v2"
)

func mandamailAlert(from, to, device string) {

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
	body := fmt.Sprintf("Alert %s Forte abbassamento sessioni ppp, %s\n", device, grafanaurl)

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

func mandamail(from, to, scopo string) (err error) {

	var subject, body string

	switch scopo {
	case "Avvio":
		subject = "Ramses - Avvio applicazione"
		body = "Ramses avviato"
	case "Update":
		subject = "Ramses - applicazione attiva"
		body = "Ramses è ancora attivo"

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

func mandamailChiusura(from, to string) (err error) {

	subject := "Ramses - Arresto applicazione"
	body := "Ramses arrestato"

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
