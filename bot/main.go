package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var msg = make(chan string, 1)

//RiceviResult riceve una stringa e la invia a telegram
func RiceviResult(result string) {
	msg <- result
	return
}

func main() {

	//Recupera la variabile d'ambiente
	TELEGRAMTOKEN, err := recuperavariabile("TELEGRAMTOKEN")
	if err != nil {
		log.Println(err)
		return
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  TELEGRAMTOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	b.Handle("/version", func(m *tb.Message) {
		b.Send(m.Chat, "Ramses_bot v2.4.1 beta")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		//b.Send(m.Sender, m.Text)
		//cerca se nella stringa di testo è presente fd + cache
		ramses, _ := regexp.Compile(`^[rR]amses\s.*`)
		if ramses.MatchString(m.Text) == true {

			//prende il secondo parametro passato via chat
			device := strings.Split(m.Text, " ")[1]

			image, err := recuperajson(device)
			if err != nil {
				log.Println(err.Error())
			}
			p := &tb.Photo{File: tb.FromDisk(image)}

			msg := fmt.Sprintf("Ecco cosa ho trovato per: %s", device)
			b.Reply(m, msg)
			b.Send(m.Chat, p)
		}

	})

	b.Start()
}
