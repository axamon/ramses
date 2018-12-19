package main

import (
	"fmt"
	"log"

	gomail "gopkg.in/gomail.v2"
)

func mandamailAlert(from, to, device string) {

	subject := "Allarme ppp su " + device
	//body := "Ciao <b>Gringo</b> <hr> rilevato abbassamento anomalo sessioni ppp su " + device

	grafanaurl := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
	body := fmt.Sprintf("Alert su %s, forte abbassamento sessioni ppp, %s\n", device, grafanaurl)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", "alberto.bregliano@telecomitalia.it", "Alberto Bregliano")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewPlainDialer(configuration.SmtpServer, configuration.SmtpPort, configuration.SmtpUser, configuration.SmtpPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("impossibile inviare mail %s\n", err.Error())
	}
}

func mandamailAvvio(from, to string) (err error) {

	subject := "Ramses - Avvio applicazione"
	body := "Ramses avviato"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", "alberto.bregliano@telecomitalia.it", "Alberto Bregliano")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(configuration.SmtpServer, configuration.SmtpPort, configuration.SmtpUser, configuration.SmtpPassword)

	errdialandsend := d.DialAndSend(m)
	if errdialandsend != nil {
		err = fmt.Errorf("Ramses Error Impossibile inviare mail %s", errdialandsend.Error())
	}
	return err
}

func mandamailChiusura(from, to string) (err error) {

	subject := "Ramses - Arresto applicazione"
	body := "Ramses arrestato"

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", "alberto.bregliano@telecomitalia.it", "Alberto Bregliano")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(configuration.SmtpServer, configuration.SmtpPort, configuration.SmtpUser, configuration.SmtpPassword)

	errdialandsend := d.DialAndSend(m)
	if errdialandsend != nil {
		err = fmt.Errorf("Ramses Error Impossibile inviare mail %s", errdialandsend.Error())
	}
	return err
}
