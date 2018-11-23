package main

import (
	"gonum.org/v1/gonum/stat"
)

func derive(serie []float64) (dserie []float64, err error) {
	length := len(serie)
	limit := length - 1
	dserie = make([]float64, length)
	for i := 0; i < limit; i++ {
		dserie[i] = serie[i+1] - serie[i]
	}
	meandserie := stat.Mean(dserie, nil)
	dserie[limit] = meandserie
	return dserie, err
}

func derive2(serie []float64) (ddserie []float64, err error) {

	dserie, err := derive(serie)
	ddserie, err = derive(dserie)

	return ddserie, err
}
