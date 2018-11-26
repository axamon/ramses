package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lytics/anomalyzer"
)

func main() {
	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  9,
		LowerBound:  1, // ignore the lower bound
		ActiveSize:  4,
		//Delay:       true,
		NSeasons: 4,
		Methods:  []string{"diff", "fence", "highrank", "lowrank", "magnitude", "cdf", "ks"},
	}

	// initialize with empty data or an actual slice of floats
	data := []float64{0.1, 2.05, 1.5, 2.5, 2.6, 2.55, 3, 3, 3, 3, 3, 3, 3, 3, 3}

	anom, _ := anomalyzer.NewAnomalyzer(conf, data)

	// the push method automatically triggers a recalcuation of the
	// anomaly probability.  The recalculation can also be triggered
	// by a call to the Eval method.

	num, _ := strconv.Atoi(os.Args[1])
	prob := anom.Push(float64(num))
	fmt.Printf("Anomalous Probability: %2.2f\n", prob)
}
