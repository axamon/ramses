package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/asaskevich/govalidator"

	"gonum.org/v1/plot"
)

//Dove salvere il nome delle interfacce
var listainterfacce []string
var listainterfacce2 = make(map[string]string)

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

	//dev è un type dell'apparato scelto
	dev := XrsMi001Stru{}

	//Questo serve a recuperare dallo struct il nome delle interfacce
	e := reflect.TypeOf(&dev).Elem()
	printTags(e)
	//fmt.Println(listainterfacce2)
	//Stampa tutte le interfacce con il numero
	for n, i := range listainterfacce {
		fmt.Println(n, i)
	}

	//Crea un lettore di input da tastiera
CHOISE:
	var choise int
	//choise = 6
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() == false {
		//fmt.Println(scanner.Text())
	}
	if govalidator.IsInt(scanner.Text()) == false {
		goto CHOISE
	}
	choise, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("hai scelto:", listainterfacce[choise])

	// choisedinterface := listainterfacce2[listainterfacce[choise]]
	choisedinterface := listainterfacce[choise]

	

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	device := "xrs-mi001"
	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	values := recuperajson(device, choisedinterface)
	//fmt.Println(values) //debug
	elaboraserie(values, choisedinterface)
}

// 	sample := elaborapunti(values)

// 	var t []float64
// 	for i := 0; i < len(values); i++ {
// 		t = append(t, values[i].Value)
// 	}

// 	err = plotutil.AddLinePoints(p,
// 		"pippo", sample)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Save the plot to a PNG file.
// 	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
// 		panic(err)
// 	}
// }

// func elaborapunti(values []float64) plotter.XYs {
// 	pts := make(plotter.XYs, len(values))
// 	for i := range pts {
// 		if i == 0 {
// 			pts[i].X = values[0].Time
// 		} else {
// 			pts[i].X = values[i].Time
// 		}
// 		pts[i].Y = values[i].Value
// 	}
// 	return pts
// }

// // randomPoints returns some random x, y points.
// func randomPoints(n int) plotter.XYs {
// 	pts := make(plotter.XYs, n)
// 	for i := range pts {
// 		if i == 0 {
// 			pts[i].X = rand.Float64()
// 		} else {
// 			pts[i].X = pts[i-1].X + rand.Float64()
// 		}
// 		pts[i].Y = pts[i].X + 10*rand.Float64()
// 	}
// 	return pts
// }
