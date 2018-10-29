package main

import (
	"log"
	"os"
	"reflect"

	"github.com/remeh/sizedwaitgroup"
)

//Dove salvere il nome delle interfacce
var listainterfacce []string
var listainterfacce2 = make(map[string]string)

var wg = sizedwaitgroup.New(8)

//var wg waitgroup

func printTags(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			column := field.Tag.Get("json")
			column2 := field.Name
			listainterfacce2[column] = column2
			listainterfacce = append(listainterfacce, column)

			//fmt.Printf("interface: %v\n", column)
			printTags(field.Type)
			continue
		}

		// column := field.Tag.Get("json")

		// fmt.Printf("interface: %v\n", column)

	}
	return
}

func main() {

	device := os.Args[1]

	ifnames := ifNames(device)

	// //dev è un type dell'apparato scelto
	// dev := XrsMi001Stru{}

	// //Questo serve a recuperare dallo struct il nome delle interfacce
	// e := reflect.TypeOf(&dev).Elem()
	// printTags(e)
	// //fmt.Println(listainterfacce2)
	// //Stampa tutte le interfacce con il numero
	// for n, i := range listainterfacce {
	// 	fmt.Println(n, i)
	// }

	//Crea un lettore di input da tastiera
	// CHOISE:
	// var choise int
	// //choise = 6
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() == false {
	// 	//fmt.Println(scanner.Text())
	// }
	// if govalidator.IsInt(scanner.Text()) == false {
	// 	goto CHOISE
	// }
	// choise, err := strconv.Atoi(scanner.Text())
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println("hai scelto:", listainterfacce[choise])

	// choisedinterface := listainterfacce2[listainterfacce[choise]]
	// choisedinterface := listainterfacce[choise]

	// p, err := plot.New()
	// if err != nil {
	// 	panic(err)
	// }

	// //device := "xrs-mi001"
	// p.Title.Text = "Plotutil example"
	// p.X.Label.Text = "X"
	// p.Y.Label.Text = "Y"

	recuperajson(device, ifnames)

	log.Printf("Elaborazione per %s terminata\n", device)
}
