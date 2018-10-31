package main

import (
        "fmt"
        
        "github.com/mjibson/go-dsp/fft"
)

func main() {
        fmt.Println(fft.FFTReal([]float64 {1, 2, 3, 1,1}))
}