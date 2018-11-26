package main

import (
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/axamon/ramses/algoritmi"
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
	xdet, ydet := algoritmi.Detrend(xs, a)

	elaboraserie(xdet, ydet, "prova", "eth0", "fuffa")
}
