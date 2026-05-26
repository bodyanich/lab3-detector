// Package main starts the image metadata processor service.
package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"

	"lab3-detector/internal/logging"
	"lab3-detector/internal/processor"

	"go.uber.org/zap"
)

func main() {
	mode := flag.String("mode", "leaky", "worker mode: leaky or fixed")
	workers := flag.Int("workers", 5, "number of workers")
	flag.Parse()

	logger, err := logging.NewDevelopmentLogger()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("failed to sync logger: %v", err)
		}
	}()

	restoreGlobals := zap.ReplaceGlobals(logger)
	defer restoreGlobals()

	go func() {
		log.Println("pprof server started on http://localhost:6060/debug/pprof/")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Printf("pprof server stopped: %v", err)
		}
	}()

	logger.Info(
		"image metadata processor started",
		zap.String("mode", *mode),
		zap.Int("workers", *workers),
	)

	switch *mode {
	case "leaky":
		processor.RunLeakyWorkerPool(*workers)
	case "fixed":
		processor.RunFixedWorkerPool(*workers)
	default:
		logger.Fatal("unknown worker mode", zap.String("mode", *mode))
	}
}
