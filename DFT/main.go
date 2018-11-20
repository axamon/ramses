package main

import "math"
import "fmt"

func DFT_naive(input []float64) ([]float64, []float64) {
	real := make([]float64, len(input))
	imag := make([]float64, len(input))
	arg := -2.0 * math.Pi / float64(len(input))
	for k := 0; k < len(input); k++ {
		r, i := 0.0, 0.0
		for n := 0; n < len(input); n++ {
			r += input[n] * math.Cos(arg*float64(n)*float64(k))
			i += input[n] * math.Sin(arg*float64(n)*float64(k))
		}
		real[k], imag[k] = r, i
	}
	return real, imag
}

func Amplitude(real, imag []float64) []float64 {
	amp := make([]float64, len(real))
	for i := 0; i < len(real); i++ {
		amp[i] = math.Sqrt(real[i]*real[i] + imag[i]*imag[i])
	}
	return amp
}

func main() {
	// create input data
	len := 256
	x := make([]float64, len)
	for i := range x {
		x[i] = math.Sin(8.0 * math.Pi * float64(i) / float64(len))
	}

	// DFT
	real, imag := DFT_naive(x)

	// Amplitude
	amp := Amplitude(real, imag)

	// Print result
	for key, value := range amp {
		fmt.Println(key, value)
	}
}
