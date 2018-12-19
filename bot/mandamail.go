package main

import (
	"log"

	gomail "gopkg.in/gomail.v2"
)

func mandamail(from, to, device string) {

	subject := "Allarme ppp su " + device
	body := "Ciao <b>Gringo</b> <hr> rilevato abbassamento anomalo sessioni ppp su " + device

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", "alberto.bregliano@telecomitalia.it", "Alberto Bregliano")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewPlainDialer(configuration.SmtpServer, configuration.SmtpPort, configuration.SmtpUser, configuration.SmtpPassword)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Printf("impossibile inviare mail %s\n", err.Error())
	}
}