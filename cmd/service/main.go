// Package main starts the Image Metadata Processor service.
package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"

	"lab3-detector/internal/processor"
)

func main() {
	mode := flag.String("mode", "leaky", "worker mode: leaky or fixed")
	workers := flag.Int("workers", 5, "number of workers")
	flag.Parse()

	go func() {
		log.Println("pprof server started on http://localhost:6060/debug/pprof/")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Printf("pprof server stopped: %v", err)
		}
	}()

	log.Printf("Image Metadata Processor started from main branch in %q mode with %d workers", *mode, *workers)

	switch *mode {
	case "leaky":
		processor.RunLeakyWorkerPool(*workers)
	case "fixed":
		processor.RunFixedWorkerPool(*workers)
	default:
		log.Fatalf("unknown mode %q: use leaky or fixed", *mode)
	}
}
