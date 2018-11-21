package main

import "gonum.org/v1/gonum/stat"

func detrend(x, y []float64) (xdet, ydet []float64) {
	//Calcolo dei pesi per elimare la eteroschedasticità
	//Ogni peso è dato dall'inverso della varianza del punto
	numSamples := len(y)
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
