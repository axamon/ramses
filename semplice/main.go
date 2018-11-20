package main

import (
	"math"
	"math/rand"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

var numSamples int
var sigma float64

func main() {
	sigma = 2.0

	numSamples, _ = strconv.Atoi(os.Args[1])

	x := func(n int) float64 {
		wave0 := 50.0 + (rand.Float64()+float64(n))*59
		//wave0 := 3*math.Sin(2*math.Pi*float64(n)/20.0) + 2.0 + float64(n)/10
		wave1 := 30 * math.Sin(2*math.Pi*float64(n)/1000) //settimanale
		wave3 := 12.0 * math.Sin(2*math.Pi*float64(n))    //gionaliero
		return wave0 + wave1 + wave3
	}

	xs := make([]float64, numSamples)
	// Discretize our function by sampling at numSamples points.
	a := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		a[i] = x(i)
		xs[i] = float64(i)
	}

	indiceanomalo, _ := strconv.Atoi(os.Args[2])
	anomalo, _ := strconv.Atoi(os.Args[3])
	a[indiceanomalo] = float64(anomalo)

	//fmt.Println(xs, a)
	xdet, ydet := detrend(xs, a)

	elaboraserie(xdet, ydet, "prova", "eth0", "fuffa")
}

//detrend
func detrend(x, y []float64) (xdet, ydet []float64) {
	//Calcolo dei pesi per elimare la eteroschedasticità
	//Ogni peso è dato dall'inverso della varianza del punto
	var weights = make([]float64, numSamples)
	for i := 0; i < len(y); i++ {
		if i <= 2 {
			weights[i] = 0
			continue
		}
		variance := stat.Variance(y[:i], nil)
		weights[i] = 1 / variance
	}
	alpha, beta := stat.LinearRegression(x, y, weights, false)

	ydet = make([]float64, numSamples)
	xdet = make([]float64, numSamples)

	for i := 0; i < len(y); i++ {
		ydet[i] = y[i] - (alpha + x[i]*beta)
		xdet[i] = x[i]
	}

	return
}
