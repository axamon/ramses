package main

import (
	"fmt"
	"math/rand"
	"reflect"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var listainterfacce []string

func printTags(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			column := field.Tag.Get("json")
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
	//rand.Seed(int64(0))

	dev := XrsMi001Stru{}

	e := reflect.TypeOf(&dev).Elem()

	printTags(e)
	fmt.Println(len(listainterfacce))

	for i := 0; i < len(listainterfacce); i++ {
		fmt.Println(listainterfacce[i])

	}

	// for i := 0; i < e.NumField(); i++ {
	// 	varName := e.Type().Field(i).Name
	// 	fmt.Printf("%v\n", varName)
	// }

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	device := "xrs-mi001"
	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	values := recuperajson(device)

	sample := elaborapunti(values)

	var t []float64
	for i := 0; i < len(values); i++ {
		t = append(t, values[i].Value)
	}
	elaboraserie(t)

	err = plotutil.AddLinePoints(p,
		"pippo", sample)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func elaborapunti(values TDATA) plotter.XYs {
	pts := make(plotter.XYs, len(values))
	for i := range pts {
		if i == 0 {
			pts[i].X = values[0].Time
		} else {
			pts[i].X = values[i].Time
		}
		pts[i].Y = values[i].Value
	}
	return pts
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
