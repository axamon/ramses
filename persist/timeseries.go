package persist

import "golang.org/x/net/internal/timeseries"

func createtimeserie(name string) error {
	timeseries.NewTimeSeries(name)
}