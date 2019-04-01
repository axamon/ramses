package main

import (
	"sort"
	"strconv"
)

func estraiSerie(result []interface{}) (serieppptime, seriepppvalue []float64) {

	// Estraggo serie dati dal risultato query http
	d := result[0].(map[string]interface{})
	dp := d["dps"].(map[string]interface{})

	// Mette i timestamps in ordine
	tempi := make([]string, 0)
	for t := range dp {
		tempi = append(tempi, t)
	}
	// Ordina i timestamps in maniera crescente
	sort.Strings(tempi)

	// Cicla i tempi
	for _, t := range tempi {
		tint, _ := strconv.Atoi(t)
		serieppptime = append(serieppptime, float64(tint))
		seriepppvalue = append(seriepppvalue, dp[t].(float64))
		//fmt.Println("orario: ", t, "valore: ", dp[t])
	}

	return serieppptime, seriepppvalue
}
