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
var image = make(chan string, 1)

//RiceviResult riceve una stringa e la invia a telegram
func RiceviResult(result string) {
	msg <- result
	return
}

//rendiamo il bot b usabile anche in altre funzioni

//var m *tb.Message
var b *tb.Bot

func main() {

	GatherInfo()

	go nasppp()

	//Recupera la variabile d'ambiente
	TELEGRAMTOKEN, err := recuperavariabile("TELEGRAMTOKEN")
	if err != nil {
		log.Println(err)
		return
	}

	b, err = tb.NewBot(tb.Settings{
		Token:  TELEGRAMTOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Println(err.Error())
	}

	b.Handle("/version", func(m *tb.Message) {
		b.Send(m.Chat, "Ramses_bot v3.5.3 beta")
		b.Send(m.Chat, "In continuo cambiamento")
	})

	ramses, _ := regexp.Compile(`^[rR]amses\s.*`)

	b.Handle(tb.OnText, func(m *tb.Message) {
		go func() {
			for {
				b.Send(m.Chat, <-msg)
			}
		}()
		go func() {
			for {
				p := &tb.Photo{File: tb.FromDisk(<-image)}
				b.Send(m.Chat, p)
			}
		}()
		//b.Send(m.Sender, m.Text)
		//cerca se nella stringa di testo è presente fd + cache
		fmt.Println(m.Text) //debug

		if ramses.MatchString(m.Text) == true {

			//prende il secondo parametro passato via chat

			device := strings.Split(m.Text, " ")[1]
			fmt.Println(device) //debug
			//answer := "Vediamo cosa trovo su " + device
			//b.Send(m.Chat, answer)
			// b.Reply(m, "Vediamo cosa trovo per", device)

			//prende identificativo interfaccia
			//icr := strings.Split(m.Text, " ")[2]

			go recuperajson(device)

			//p := &tb.Photo{File: tb.FromDisk(image)}

			//msg := fmt.Sprintf("Ecco cosa ho trovato per: %s", device)

		}

	})

	b.Start()
}
