package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	hdFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Name:        "hd_errors_total",
		Help:        "Number of hard-disk errors.",
		ConstLabels: prometheus.Labels{"version": "1234"},
	})
)

var server http.Server

func Init() {
	prometheus.MustRegister(hdFailures)
	server = http.Server{
		Addr: ":3000",
		Handler: promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{
				EnableOpenMetrics: false,
			},
		),
	}

	server.ListenAndServe()
}

func Close() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	server.Shutdown(ctx)
}
