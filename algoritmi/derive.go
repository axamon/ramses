package algoritmi

import (
	"gonum.org/v1/gonum/stat"
)

//Derive execute the first derivate of the serie
func Derive(serie []float64) (dserie []float64, err error) {
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

//Derive2 execute the second derivate of the serie
func Derive2(serie []float64) (ddserie []float64, err error) {

	dserie, err := Derive(serie)
	ddserie, err = Derive(dserie)

	return ddserie, err
}

//Derive3 execute the third derivate of the serie
func Derive3(serie []float64) (dddserie []float64, err error) {

	dserie, err := Derive(serie)
	ddserie, err := Derive(dserie)
	dddserie, err = Derive(ddserie)

	return dddserie, err
}
