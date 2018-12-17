package funzioni

import (
	"fmt"
	"os"
)

//Recuperavariabile recupera valori dalle variabili d'ambiente
func Recuperavariabile(variabile string) (result string, err error) {
	if result, ok := os.LookupEnv(variabile); ok && len(result) != 0 {
		return result, nil
	}
	err = fmt.Errorf("la variabile %s non esiste o è vuota", variabile)
	fmt.Fprintln(os.Stderr, err.Error())
	return
}
