package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thebinary/ntpq_exporter/exporter"
)

var Version = "rc"
var (
	showVersion     bool
	listenAddr      string
	metricsEndpoint string
)

func main() {
	flag.StringVar(
		&listenAddr, "l",
		":9123",
		"Prometheus Listening Address",
	)
	flag.StringVar(
		&metricsEndpoint, "m",
		"/metrics",
		"Prometheus metrics endpoint",
	)
	flag.BoolVar(&showVersion, "version", false, "show version")

	flag.Parse()

	if showVersion {
		fmt.Printf("version: %s\n", Version)
		os.Exit(0)
	}

	ntpqSysStatsExporter := exporter.NewSysStatsExporter()
	prometheus.MustRegister(ntpqSysStatsExporter)

	http.Handle(metricsEndpoint, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
			 <head><title>Radmin Exporter</title></head>
			 <body>
			 <h1>Radmin Exporter</h1>
			 <p><a href='` + metricsEndpoint + `'>Metrics</a></p>
			 </body>
			 </html>
			 `))
	})
	log.Printf("[INFO] Starting listener at: '%s'", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
