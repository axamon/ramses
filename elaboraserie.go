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

func elaboraserie(lista []float64, interfaccia string) {

	speeds := lista

	re := regexp.MustCompile("(ICR-.[0-9]+/[0-9]+)")
	nameICR := re.FindStringSubmatch(interfaccia)[0]
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	nomeimmagine := reg.ReplaceAllString(nameICR, "")

	// fmt.Println(choisedinterface)

	//var numchunks int
	//numchunks = len(speeds)

	//mean := stat.Mean(speeds, nil)
	//fmt.Printf("Media: %.3f\n", stat.Mean(speeds, nil))
	//harmonicmean := stat.HarmonicMean(speeds, nil)

	//fmt.Printf("MediaArmonica: %.3f\n", stat.HarmonicMean(speeds, nil))
	//mode, _ := stat.Mode(speeds, nil)
	//mean, _ := stat.Mode(speeds, nil)

	//fmt.Printf("Moda: %.3f\n", mode)
	//nums := speeds
	//fmt.Println(len(nums))
	//entropy := stat.Entropy(nums)
	//sort.Float64s(nums) //Mette in ordine nums
	//fmt.Printf("Mediana: %.3f\n", stat.Quantile(0.5, stat.Empirical, nums, nil))
	//median := stat.Quantile(0.5, stat.Empirical, nums, nil)
	//percentile95 := stat.Quantile(0.95, stat.Empirical, nums, nil)

	//stdev := stat.StdDev(speeds, nil)
	//stderr := stat.StdErr(stdev, float64(numchunks))
	//fmt.Printf("StDev: %.3f\n", stat.StdDev(speeds, nil))
	//skew := stat.Skew(speeds, nil)
	//fmt.Printf("Skew: %.3f\n", stat.Skew(speeds, nil))
	//curtosi := stat.ExKurtosis(speeds, nil)
	//chisquare := stat.ChiSquare(nums, speeds)

	/* l, err := json.Marshal(fe)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(string(l)) */
	// 	kalmansample(speeds, stdev)

	// 	return

	// }

	// func kalmansample(speeds []float64, stdev float64) {

	sma3 := ma.ThreadSafe(ma.NewSMA(3))   //creo una moving average a 3
	sma7 := ma.ThreadSafe(ma.NewSMA(7))   //creo una moving average a 7
	sma20 := ma.ThreadSafe(ma.NewSMA(20)) //creo una moving average a 20
	sma100 := ma.ThreadSafe(ma.NewSMA(100))
	//
	// kalman filter
	//

	//sstd := 0.000001
	//sstd := stdev
	//ostd := 0.5

	// trend model
	//filter, err := kalman.New(&kalman.Config{
	// 	F: mat64.NewDense(2, 2, []float64{2, -1, 1, 0}),
	// 	G: mat64.NewDense(2, 1, []float64{1, 0}),
	// 	Q: mat64.NewDense(1, 1, []float64{sstd}),
	// 	H: mat64.NewDense(1, 2, []float64{1, 0}),
	// 	R: mat64.NewDense(1, 1, []float64{ostd}),
	// })
	// if err != nil {
	// 	panic(err)
	// }

	n := len(speeds)
	fmt.Println(n)
	//s := mat64.NewDense(1, n, nil)
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

		//Verifica anomalia
		if i > len(speeds)-120 {
			if yaryOrig[i] > ma20Upperband[i] {
				fmt.Fprint(os.Stderr, "violazione soglia:", yaryOrig[i], xary[i], "\n")
				// if i >= len(yaryOrig)-120 {

				// 	ptp := make([]plotter.XYer, 120)
				// 	xys := make(plotter.XYs, 120)
				// 	ptp[i] = xys
				// 	xys[i].X = xary[i]
				// 	xys[i].Y = yaryOrig[i]
				// 	plotutil.AddScatters(p, ptp[i]
				//)

				// }
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
	if err := p.Save(8*vg.Inch, 4*vg.Inch, nomeimmagine+".png"); err != nil {
		panic(err)
	}
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
