package algoritmi

import (
	"sort"

	"gonum.org/v1/gonum/stat"
)

// ScremaValori elimina i valori oltre il qmax percentile
// e sotto il qmin percentile
func ScremaValori(values []float64, qmax, qmin float64) (speeds []float64) {

	// Verifiche quartili
	// Creo una struttura di appoggio
	var q = make([]float64, len(values))
	speeds = make([]float64, len(values))

	// Devo copiare la slide in entrata su una struttura di appoggio
	// per sortarla
	copy(q, values)

	// Sorto la struttura di appoggio e NON i valori reali
	sort.Float64s(q)

	// Calcola il qmin quartile
	percentile1 := stat.Quantile(qmin, stat.Empirical, q, nil)

	// Calcola il qmax quartile
	percentile99 := stat.Quantile(qmax, stat.Empirical, q, nil)

	// Elimina valori troppo strani
	for i := 0; i < len(values); i++ {
		speeds[i] = values[i]

		//se il valoro è troppo estremo lo sostituisce con 0
		if values[i] > percentile99 || values[i] < percentile1 {
			speeds[i] = 0
		}

	}
	return speeds
}
