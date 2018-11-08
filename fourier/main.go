package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"github.com/mjibson/go-dsp/spectral"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

//PwelchOptions computa la potenza spettrale
type PwelchOptions struct {
	// NFFT is the number of data points used in each block for the FFT. Must be
	// even; a power 2 is most efficient. This should *NOT* be used to get zero
	// padding, or the scaling of the result will be incorrect. Use Pad for
	// this instead.
	//
	// The default value is 256.
	NFFT int

	// Window is a function that returns an array of window values the length
	// of its input parameter. Each segment is scaled by these values.
	//
	// The default (nil) is window.Hann, from the go-dsp/window package.
	Window func(int) []float64

	// Pad is the number of points to which the data segment is padded when
	// performing the FFT. This can be different from NFFT, which specifies the
	// number of data points used. While not increasing the actual resolution of
	// the psd (the minimum distance between resolvable peaks), this can give
	// more points in the plot, allowing for more detail.
	//
	// The value default is 0, which sets Pad equal to NFFT.
	Pad int

	// Noverlap is the number of points of overlap between blocks.
	//
	// The default value is 0 (no overlap).
	Noverlap int

	// Specifies whether the resulting density values should be scaled by the
	// scaling frequency, which gives density in units of Hz^-1. This allows for
	// integration over the returned frequency values. The default is set for
	// MATLAB compatibility. Note that this is the opposite of matplotlib style,
	// but with equivalent defaults.
	//
	// The default value is false (enable scaling).
	Scale_off bool
}

var pippo spectral.PwelchOptions

func main() {
	//pippo.Scale_off = true
	numSamples, _ := strconv.Atoi(os.Args[1])

	// Equation 3-10.
	x := func(n int) float64 {
		wave0 := 2.0 + (rand.Float64() + float64(n)/10) + rand.Float64()*5
		//wave0 := 3*math.Sin(2*math.Pi*float64(n)/20.0) + 2.0 + float64(n)/10
		wave1 := 3 * math.Sin(2*math.Pi*float64(n)/6.0)
		wave3 := 4.0 * math.Sin(2*math.Pi*float64(n)/3.0)
		return wave0 + wave1 + wave3
	}

	xs := make([]float64, numSamples)
	// Discretize our function by sampling at numSamples points.
	a := make([]float64, numSamples)
	for i := 0; i < numSamples; i++ {
		a[i] = x(i)
		xs[i] = float64(i)
	}

	//Effettua la regressione lineare
	var weights []float64
	alpha, beta := stat.LinearRegression(xs, a, weights, false)
	//r2 := stat.RSquared(xs, a, weights, alpha, beta)
	//fmt.Println(alpha, beta, r2)

	//creo un risultato anomalo
	indiceanomalo, _ := strconv.Atoi(os.Args[2])
	anomalo, _ := strconv.Atoi(os.Args[3])
	a[indiceanomalo] = float64(anomalo)

	//X := fft.FFTReal(a)
	//l := fft.IFFT(X)

	//fmt.Println(a)
	//ll := make([]complex128, numSamples)
	diff := make([]float64, numSamples)

	//Elimina il trend
	for i := 0; i < numSamples; i++ {
		//p, f := cmplx.Polar(X[i])
		///= p * math.Exp(f*math.Sqrt(-1))
		diff[i] = a[i] - (alpha + beta*xs[i])
		//diff[i] = a[i] - beta*xs[i] //lascio alpha

	}

	//Analisi spettrale
	//	Pxx, ff := spectral.Pwelch(diff, float64(300), &pippo)
	Pxx, ff := spectral.Pwelch(diff, float64(300), &pippo)

	fmt.Println(len(Pxx))
	//pxxmedia := stat.Mean(Pxx, nil)
	//ffmedia := stat.Mean(Pxx, nil)

	//Crea mappa delle frequenze
	frequenze := make(map[float64]float64)

	//Associa frequenza a Potenza spettrale

	for l := 0; l < len(Pxx); l++ {
		/*//Se l'ampiezza è troppo bassa o la frequenza troppo alta escludi armonica
		if Pxx[l] < pxxmedia || ff[l] > 2*ffmedia {
			continue
		}
		*/
		frequenze[Pxx[l]] = ff[l]
	}

	fmt.Println(len(frequenze))

	//Mette in ordine crescente i poteri spettrali
	sort.Sort(sort.Reverse(sort.Float64Slice(Pxx)))

	mediaPxx := stat.Mean(Pxx, nil)
	lenp := len(Pxx)
	sumPxx := mediaPxx * float64(lenp)

	//Debug
	//for n, p := range Pxx {
	//	fmt.Printf("%v,%2.2f\n", n, p/sumPxx*100)
	//}
	//Mette in ordine decrescente i poteri spettrali
	//sort.Sort(sort.Float64Slice(Pxx))

	//Crea una sinusoide con frequenza e ampiezza date
	sinwave := func(n int, ampiezza, frequenza float64) float64 {
		wave := ampiezza / sumPxx * math.Sin(2*math.Pi*frequenza*float64(n))
		return wave
	}

	//discretizza la sinusoide creata
	sinwavediscrete := make([]float64, numSamples)
	sinwavediscrete1 := make([]float64, numSamples)
	sinwavediscrete2 := make([]float64, numSamples)
	sinwavediscrete3 := make([]float64, numSamples)
	sinwavediscrete4 := make([]float64, numSamples)
	sinwavediscrete5 := make([]float64, numSamples)
	sinwavediscrete6 := make([]float64, numSamples)
	sinwavediscrete7 := make([]float64, numSamples)
	sinwavediscrete8 := make([]float64, numSamples)
	sinwavediscrete9 := make([]float64, numSamples)
	sinwavediscrete10 := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {

		sinwavediscrete[i] = sinwave(i, Pxx[0], frequenze[Pxx[0]])
		sinwavediscrete1[i] = sinwave(i, Pxx[1], frequenze[Pxx[1]])
		sinwavediscrete2[i] = sinwave(i, Pxx[2], frequenze[Pxx[2]])
		sinwavediscrete3[i] = sinwave(i, Pxx[3], frequenze[Pxx[3]])
		sinwavediscrete4[i] = sinwave(i, Pxx[4], frequenze[Pxx[4]])
		sinwavediscrete5[i] = sinwave(i, Pxx[5], frequenze[Pxx[5]])
		sinwavediscrete6[i] = sinwave(i, Pxx[6], frequenze[Pxx[6]])
		sinwavediscrete7[i] = sinwave(i, Pxx[7], frequenze[Pxx[7]])
		sinwavediscrete8[i] = sinwave(i, Pxx[8], frequenze[Pxx[8]])
		sinwavediscrete9[i] = sinwave(i, Pxx[9], frequenze[Pxx[9]])
		sinwavediscrete10[i] = sinwave(i, Pxx[10], frequenze[Pxx[10]])

	}

	//mediaA, _ := stat.Mode(a, nil)

	//Elimina la stagionalità
	var stagionalita []float64
	for i := 0; i < numSamples; i++ {
		//p, f := cmplx.Polar(X[i])
		///= p * math.Exp(f*math.Sqrt(-1))
		//diff[i] = diff[i] - math.Abs(diff[i]*sinwavediscrete[i]) - math.Abs(diff[i]*sinwavediscrete1[i]*sinwavediscrete2[i])
		sumsinewaves := (sinwavediscrete[i] +
			sinwavediscrete1[i] +
			sinwavediscrete2[i] +
			sinwavediscrete3[i] +
			sinwavediscrete4[i] +
			sinwavediscrete5[i] +
			sinwavediscrete6[i] +
			sinwavediscrete7[i] +
			sinwavediscrete8[i] +
			sinwavediscrete9[i] +
			sinwavediscrete10[i])

		stagionalita = append(stagionalita, sumsinewaves*alpha+alpha)

		diff[i] = diff[i] - sumsinewaves

	}

	//diffmean := sPPANICtat.Mean(diff, nil)
	//diffmode, _ := stat.Mode(diff, nil)
	diffmean, diffstd := stat.MeanStdDev(diff, nil)
	for i := 0; i < len(diff); i++ {
		if math.Abs(diff[i]) < diffmean+3*diffstd {
			diff[i] = 0
		}
	}
	/*
		mediadiff, _ := stat.Mode(diff, nil)
		fmt.Println(mediadiff)

		for i := 0; i < len(diff); i++ {
			diff[i] = diff[i] * math.Abs(mediadiff)
		}
	*/
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Outlier artificiale individuato"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		//
		"Andamento reale", points(a, numSamples),
		"Stagionalità", points(stagionalita, numSamples),
		//"SS", points(ss, numSamples),
		//"Fourier", points(l, numSamples),
		//
		"Anomalie", points(diff, numSamples),
	)

	if err != nil {
		log.Println(err.Error())
	}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func points(line []float64, n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = 0
		} else {
			pts[i].X = pts[i-1].X + 1
		}
		pts[i].Y = line[i]
	}
	return pts
}
