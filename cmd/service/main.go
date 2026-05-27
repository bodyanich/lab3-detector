// Package main starts the image metadata processor service.
package main

import (
	"flag"
	"log"
	"net/http"

	_ "net/http/pprof"

	"lab3-detector/internal/processor"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var version = "dev"

func main() {
	mode := flag.String("mode", "fixed", "processor mode: leaky or fixed")
	workers := flag.Int("workers", 5, "number of worker goroutines")
	flag.Parse()

	log.Printf("Image Metadata Processor version: %s", version)

	go startPprofServer()
	go startMetricsServer()

	log.Printf("Image Metadata Processor started in %q mode with %d workers", *mode, *workers)

	switch *mode {
	case "leaky":
		processor.RunLeakyWorkerPool(*workers)
	case "fixed":
		processor.RunFixedWorkerPool(*workers)
	default:
		log.Fatalf("unknown mode: %s", *mode)
	}
}

func startPprofServer() {
	log.Println("pprof server started on http://localhost:6060/debug/pprof/")
	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
		log.Printf("pprof server stopped: %v", err)
	}
}

func startMetricsServer() {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	log.Println("metrics server started on http://localhost:2112/metrics")
	if err := http.ListenAndServe(":2112", mux); err != nil {
		log.Printf("metrics server stopped: %v", err)
	}
}
