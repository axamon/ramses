package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/axamon/ramses/algoritmi"
	"gonum.org/v1/gonum/stat"
)

var sigma = configuration.Sigma

func elaboroRequest(result []interface{}, device string) {

	// Estraggo serie dati dal risultato query http
	d := result[0].(map[string]interface{})
	dp := d["dps"].(map[string]interface{})

	// Mette i timestamps in ordine
	tempi := make([]string, 0)
	for t := range dp {
		tempi = append(tempi, t)
	}
	// Ordina i timestamps in maniera crescente
	sort.Strings(tempi)

	// Crea variabili di appoggio
	var seriepppvalue []float64
	var serieppptime []float64

	// Cicla i tempi
	for _, t := range tempi {
		tint, _ := strconv.Atoi(t)
		serieppptime = append(serieppptime, float64(tint))
		seriepppvalue = append(seriepppvalue, dp[t].(float64))
		//fmt.Println("orario: ", t, "valore: ", dp[t])
	}

	// Se non ci sono abbastanza valori per la serie esce
	if len(seriepppvalue) < 300 {
		log.Printf("Error %s Non ci sono abbastanza dati per elaborare statistiche\n", device)
		return
	}

	// Mofifica serie prima che sia elaborata

	// Elimino il trend
	xdet, ydet := algoritmi.Detrend(serieppptime, seriepppvalue)

	// Applico Derivata terza
	y, _ := algoritmi.Derive3(ydet)

	// Calcolo statistiche sulla serie elaborata
	mean, stdev := stat.MeanStdDev(y, nil)
	// log.Printf("%s Info media: %2.f stdev: %2.f", device, mean, stdev) // debug

	for i := 10; i < len(y); i++ {
		// Individuo se è avvenuto un Jerk
		if y[i] < mean-sigma*stdev {
			unixtimeUTC := time.Unix(int64(xdet[i]/1000), 0)
			// Serve per avere il timestamp di quando c'è stato il problema
			unixtimeinRFC3339 := unixtimeUTC.Format(time.RFC3339)

			// Devo verificare se valori futuri dopo il Jerk hanno avuto problemi
			numvalori := len(seriepppvalue)
			for l := 0; l <= 6; l++ {

				// Evita che si arrivi alla fine della serie di valori
				if i+l > numvalori-1 {
					break
				}
				// Verifica i valori dopo il jerk
				limite := (seriepppvalue[i] - seriepppvalue[i+l]) / seriepppvalue[i]

				// Se il limite è negativo non ci interessa
				if limite < 0 {
					continue
				}

				//log.Println(seriepppvalue[1-1], seriepppvalue[i], limite)
				//fmt.Printf("%s %s Jerk Ultimovalore: %2.f, Penultimovalore: %2.f, Limite: %.4f, %v\n", unixtimeinRFC3339, device, seriepppvalue[i], seriepppvalue[i+l], limite, l)
				//fmt.Printf("%s Info media: %2.f stdev: %2.f , Penultimovalore: %2.f, Differenza: %2.f\n", device, mean, stdev, seriepppvalue[i-1], seriepppvalue[i]-seriepppvalue[i-1])

				if limite > configuration.Soglia {
					summary := fmt.Sprintf("abbassamento sessioni ppp superiore al %2.0f%%\n", configuration.Soglia*100)
					// Attenzione NON usare log.Print perchè serve printare il timestamp non attuale ma di quando si è verificato il problema
					fmt.Printf("%s Alert %s, %s\n", unixtimeinRFC3339, device, summary)

					// Mandamail di notifica solo se siamo negli ultimi 6 valori
					if i > (numvalori - 6) {
						mandamailAlert(configuration.SmtpFrom, configuration.SmtpTo, device)
						err := CreaTrap(device, "sessioni ppp", summary, listanasip[device], 1, 5)
						if err != nil {
							log.Printf("Error %s Impossibile inviare trap\n", device)
						}
						nastrappati[device] = true
					}
				}
			}
		}
	}

}
