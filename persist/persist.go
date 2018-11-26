package persist

import (
	"encoding/gob"
	"os"
	"sync"
)

//Crea un mutex per gestire l'accesso concorrente delle go-routines
var gobfileLock = sync.RWMutex{}

//SaveLocally encodes in Gob the passed data in a local file
func SaveLocally(path string, object interface{}) error {

	//Crea il file dove persistere i dati
	file, err := os.Create(path)

	//Prima di terminare la funzione chiude il file
	defer file.Close()

	//Se non ci sono errori
	if err == nil {

		//Crea un encoder
		encoder := gob.NewEncoder(file)

		//Applica il mutex a lettura e scrittura
		gobfileLock.Lock()

		//Encoda i dati
		encoder.Encode(object)

		//Sblocca il mutex
		gobfileLock.Unlock()
	}

	return err
}

//Load retrieves and decodes Gob ciphered data form local file
func Load(path string, object interface{}) error {
	gobfileLock.RLock()
	defer gobfileLock.RUnlock()
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// func check(e error) {
// 	if e != nil {
// 		_, file, line, _ := runtime.Caller(1)
// 		log.Printf("Errore su linea %d, per file %s: %s\n", line, file, e.Error())
// 	}
// }
