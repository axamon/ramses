package main

import (
	"regexp"

	//"github.com/spf13/viper"

	ma "github.com/mxmCherry/movavg"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	//"compress/gzip"

	"fmt"
	"log"
	"os"
)

func elaboraserie(lista []float64, device, interfaccia string) {
	defer wg.Done()

	speeds := lista

	if len(speeds) < 120 {
		log.Println("Non ci sono abbastanza dati:", device, interfaccia)
		return
	}

	// re := regexp.MustCompile("(ICR-.[0-9]+/[0-9]+)")
	// subnames := re.FindStringSubmatch(interfaccia)
	// if len(subnames) >= 0 {
	// 	nameICR := subnames[0]
	// }

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	nomeimmagine := reg.ReplaceAllString(interfaccia, "")

	sma3 := ma.ThreadSafe(ma.NewSMA(3))   //creo una moving average a 3
	sma7 := ma.ThreadSafe(ma.NewSMA(7))   //creo una moving average a 7
	sma20 := ma.ThreadSafe(ma.NewSMA(20)) //creo una moving average a 20
	sma100 := ma.ThreadSafe(ma.NewSMA(100))

	n := len(speeds)
	//fmt.Println(n)

	x, dx := 0.0, 0.01
	xary := make([]float64, 0, n)
	yaryOrig := make([]float64, 0, n)
	ma3 := make([]float64, 0, n)
	ma7 := make([]float64, 0, n)
	ma20 := make([]float64, 0, n)
	ma100 := make([]float64, 0, n)
	ma20Upperband := make([]float64, 0, n)
	ma20Lowerband := make([]float64, 0, n)

	//
	// plot
	//

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	var i int
	//aryOrig := speeds
	for i = 0; i < n; i++ {
		//y := math.Sin(x) + 0.1*(rand.NormFloat64()-0.5)
		y := speeds[i]
		//s.Set(0, i, y)
		x += dx

		xary = append(xary, x)

		ma3 = append(ma3, sma3.Add(y))       //aggiung alla media mobile il nuovo valore e storo la media
		ma7 = append(ma7, sma7.Add(y))       //aggiung alla media mobile il nuovo valore e storo la media
		ma20 = append(ma20, sma20.Add(y))    //aggiung alla media mobile il nuovo valore e storo la media
		ma100 = append(ma100, sma100.Add(y)) //aggiung alla media mobile il nuovo valore e storo la media
		//yaryOrig = append(yaryOrig, y-ma20[i])
		yaryOrig = append(yaryOrig, y)

		var devstdBands float64
		if i >= 100 {
			devstdBands = stat.StdDev(speeds[i-99:i], nil)
		}
		// ma20Upperband = append(ma20Upperband, sma20.Avg()+3*devstdBands)
		// ma20Lowerband = append(ma20Lowerband, sma20.Avg()-3*devstdBands)
		var sigma float64
		sigma = 2
		ma20Upperband = append(ma20Upperband, sma20.Avg()+sigma*devstdBands)
		ma20Lowerband = append(ma20Lowerband, sma20.Avg()-sigma*devstdBands)

		//Verifica anomalie
		if i > len(speeds)-3 { //Confronto solo gli ultimi3 valori per un ROPLdi 15 minuti
			if yaryOrig[i] > ma20Upperband[i] {
				fmt.Fprint(os.Stderr, "violazione soglia alta:", device, interfaccia, yaryOrig[i], xary[i], "\n")
			}

			if yaryOrig[i] < ma20Lowerband[i] {
				fmt.Fprint(os.Stderr, "violazione soglia bassa:", device, interfaccia, yaryOrig[i], xary[i], "\n")
			}
		}
	}

	//filtered := filter.Filter(s)
	//yaryFilt := mat64.Row(nil, 0, filtered)

	err = plotutil.AddLinePoints(p,

		//"Filtered", generatePoints(xary, yaryFilt[len(yaryFilt)-120:]),
		//"MA3", generatePoints(xary, ma3),
		//"MA7", generatePoints(xary, ma7),
		"Up 2 sigma", generatePoints(xary, ma20Upperband[len(ma20Upperband)-120:len(ma20Upperband)-1]),
		"Original", generatePoints(xary, yaryOrig[len(yaryOrig)-120:len(yaryOrig)-1]),
		"Media mobile 20", generatePoints(xary, ma20[len(ma20)-120:]),
		"Media mobile 100", generatePoints(xary, ma100[len(ma20)-120:]),
		"Low 2 sigma", generatePoints(xary, ma20Lowerband[len(ma20Lowerband)-120:len(ma20Lowerband)-1]),
	)
	if err != nil {
		log.Println(err)
	}

	// Save the plot to a PNG file.
	p.Title.Text = device + "\n " + interfaccia
	if err := p.Save(8*vg.Inch, 4*vg.Inch, nomeimmagine+".png"); err != nil {
		panic(err)
	}
	return
}

func generatePoints(x []float64, y []float64) plotter.XYs {
	//pts := make(plotter.XYs, len(x))
	pts := make(plotter.XYs, 119)

	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}

	return pts
}
