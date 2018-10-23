package main

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	//rand.Seed(int64(0))

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
