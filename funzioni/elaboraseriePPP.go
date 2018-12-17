package funzioni

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	//"github.com/spf13/viper"

	ma "github.com/mxmCherry/movavg"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	//"compress/gzip"

	"log"
)

func elaboraseriePPP(ctx context.Context, x, y []float64, device, interfaccia, metrica string) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	defer log.Printf("%s Info Recupero dati terminato\n", device)
	//Finita la funzione notifica il waitgroup
	defer wg.Done()

	//Load data from local persisting storage

	//Execute reconciliation of data removing outliers

	//Add new data and look for anomalies

	//Persist data

	for {
		select {
		case <-ctx.Done():
			log.Printf("%s Error Superato tempo massimo per elaborazione dati\n", device)
			return
		default:
			//Se non trovo abbastanza dati esci
			if len(y) < 300 {
				log.Println("Non ci sono abbastanza dati: ", device)
				return
			}

			//forzo il valore di sigma
			sigma = 2.5

			//sendimage serve a decidere in seguito se inviare immagine grafico a telegram
			//Viene settata su true nella fase di verifica anomalie
			var sendimage bool

			//Creazione medie mobili di interesse
			//sma3 := ma.ThreadSafe(ma.NewSMA(3))     //creo una moving average a 3
			//sma7 := ma.ThreadSafe(ma.NewSMA(7))     //creo una moving average a 7
			sma20 := ma.ThreadSafe(ma.NewSMA(20)) //creo una moving average a 20
			//sma100 := ma.ThreadSafe(ma.NewSMA(100)) //creo una moving average a 100

			n := len(y)
			fmt.Println(n) //debug

			//Crea contenitori parametrizzati al numero n di elementi in entrata
			xary := make([]float64, 0, n)
			yaryOrig := make([]float64, 0, n)
			//ma3 := make([]float64, 0, n)
			//ma7 := make([]float64, 0, n)
			ma20 := make([]float64, 0, n)
			//ma100 := make([]float64, 0, n)
			ma20Upperband := make([]float64, 0, n)
			ma20Lowerband := make([]float64, 0, n)
			//
			// plot
			//

			//Inzializza il grafico
			p, err := plot.New()
			if err != nil {
				panic(err)
			}

			//inzializza il puntatore
			var i int

			//aryOrig := speeds
			for i = 0; i < n-1; i++ {
				//applico uno smoothing delle ordinate
				//speeds[i] = speeds[i+1] - speeds[i]
				ypoint := y[i]

				xary = append(xary, x[i])

				//	ma3 = append(ma3, sma3.Add(y))       //aggiung alla media mobile il nuovo valore e storo la media
				//	ma7 = append(ma7, sma7.Add(y))       //aggiung alla media mobile il nuovo valore e storo la media
				ma20 = append(ma20, sma20.Add(ypoint)) //aggiung alla media mobile il nuovo valore e storo la media
				//	ma100 = append(ma100, sma100.Add(y)) //aggiung alla media mobile il nuovo valore e storo la media
				//yaryOrig = append(yaryOrig, y-ma20[i])
				yaryOrig = append(yaryOrig, ypoint)

				var devstdBands float64
				if i >= 300 {
					//Calcola le bande di bollinger su 288 punti che sono una settimana
					devstdBands = stat.StdDev(y[i-288:i], nil)
				}
				// ma20Upperband = append(ma20Upperband, sma20.Avg()+3*devstdBands)
				// ma20Lowerband = append(ma20Lowerband, sma20.Avg()-3*devstdBands)

				ma20Upperband = append(ma20Upperband, sma20.Avg()+sigma*devstdBands)
				ma20Lowerband = append(ma20Lowerband, sma20.Avg()-sigma*devstdBands)

				//Verifica anomalie
				if i > len(y)-3 { //Confronto solo gli ultimi3 valori per un ROPL di 15 minuti
					if yaryOrig[i] > ma20Upperband[i] {
						log.Printf("Violata soglia alta %s %s. Intf: %s, valore: %.2f", device, metrica, interfaccia, yaryOrig[i])
						//alert := fmt.Sprintf("Violata soglia alta %s %s. Intf: %s, valore: %.2f", device, metrica, interfaccia, yaryOrig[i])
						//msg <- alert
						//sendimage = true
						//TODO inviare alert

					}

					if yaryOrig[i] < ma20Lowerband[i] {
						log.Printf("Violata soglia bassa %s %s. Intf: %s, valore: %.2f", device, metrica, interfaccia, yaryOrig[i])
						//urlgrafana := "https://ipw.telecomitalia.it/grafana/dashboard/db/bnas?orgId=1&var-device=" + device
						//alert := fmt.Sprintf("Violata soglia bassa %s su %s. %s", metrica, device, urlgrafana)
						//TODO inviare alert
						//msg <- alert
						sendimage = true

					}
				}
			}

			//filtered := filter.Filter(s)
			//yaryFilt := mat64.Row(nil, 0, filtered)

			//salva sigma come string
			sigmastring := strconv.FormatFloat(sigma, 'f', 1, 64)

			err = plotutil.AddLinePoints(p,

				//"Filtered", generatePoints(xary, yaryFilt[len(yaryFilt)-120:]),
				//"MA3", generatePoints(xary, ma3),
				//"MA7", generatePoints(xary, ma7),

				"Up "+sigmastring+" sigma", generatePoints(xary, ma20Upperband),
				"Original", generatePoints(xary, yaryOrig),
				"Media mobile 20", generatePoints(xary, ma20),
				//"Media mobile 100", generatePoints(xary, ma100),
				"Low "+sigmastring+" sigma", generatePoints(xary, ma20Lowerband),
				//		"lower", generatePoints(xary, lower),
			)
			if err != nil {
				log.Println(err)
			}

			// Save the plot to a PNG file.

			//imposta su due righe del grafico nome apparato e interfaccia
			p.Title.Text = device + "\n " + interfaccia + "\n" + metrica

			path1 := "./grafici"
			path2 := path1 + "/" + device
			path3 := path2 + "/" + metrica

			paths := []string{path1, path2, path3}

			for _, path := range paths {
				if _, err := os.Stat(path); os.IsNotExist(err) {
					os.Mkdir(path, 664)
				}
			}
			fmt.Println(path3) //debug

			//SALVA IL GRAFICO
			if err := p.Save(8*vg.Inch, 4*vg.Inch, path3+"/"+device+".png"); err != nil {
				panic(err)
			}

			if sendimage == true {
				//image <- path3 + "/" + device + ".png"
			}

			return
		}
	}
}
