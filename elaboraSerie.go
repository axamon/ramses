package main

import (
	"github.com/axamon/ramses/algoritmi"
	"github.com/gonum/stat"
)

func elaboraSerie(
	serieppptime, seriepppvalue []float64) (
	mean, stdev float64, xdet, y []float64) {

	// Mofifica serie prima che sia elaborata
	// Elimino il trend
	xdet, ydet := algoritmi.Detrend(serieppptime, seriepppvalue)

	// Applico Derivata terza
	y, _ = algoritmi.Derive3(ydet)

	// Calcolo statistiche sulla serie elaborata
	mean, stdev = stat.MeanStdDev(y, nil)
	// log.Printf("%s Info media: %2.f stdev: %2.f", device, mean, stdev)

	return mean, stdev, xdet, y
}
