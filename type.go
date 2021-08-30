package AppExporter


import (
	"github.com/prometheus/client_golang/prometheus"
	"context"
	"log"
)



type Scraper interface {
	Name() string

	Help() string

	Version() float64

	Scraper(ctx context.Context,Client interface{},ch chan<-prometheus.Metric, logger log.Logger) error


}