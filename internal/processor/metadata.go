// Package processor contains image metadata processing logic used for profiling experiments.
package processor

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
)

var (
	// LeakCache intentionally grows forever in leaky mode.
	// It is used to demonstrate heap profiling in Lab 3.
	LeakCache = make(map[string][]byte)
	leakMu    sync.Mutex

	fixedCacheMu sync.Mutex
	fixedCache   = make(map[string][]byte)

	imageDataPattern = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)

	processedImages atomic.Uint64
)

const maxFixedCacheItems = 100

func logProcessedImage(workerID int, mode string) {
	processedCount := processedImages.Add(1)

	if processedCount%1000 == 0 {
		zap.L().Info(
			"images processed",
			zap.Uint64("processed_count", processedCount),
			zap.Int("worker_id", workerID),
			zap.String("mode", mode),
		)
	}
}

// RunLeakyWorkerPool starts worker goroutines that intentionally leak memory.
func RunLeakyWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImageLeaky(id)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	select {}
}

// RunFixedWorkerPool starts worker goroutines with optimized processing and bounded memory usage.
func RunFixedWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImageFixed(id)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	select {}
}

func processImageLeaky(workerID int) bool {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())

	// Inefficient hot path: regexp is compiled on every call.
	matched, _ := regexp.MatchString(`^image_data_\d+_timestamp_\d+$`, data)
	if matched {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())

		leakMu.Lock()
		LeakCache[key] = make([]byte, 1024*10) // 10 KB leaked per processed image.
		leakMu.Unlock()

		logProcessedImage(workerID, "leaky")
	}

	return matched
}

func processImageFixed(workerID int) bool {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())

	// Optimized hot path: regexp is compiled once at package initialization.
	matched := imageDataPattern.MatchString(data)
	if matched {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())
		storeBounded(key, make([]byte, 1024*10))

		logProcessedImage(workerID, "fixed")
	}

	return matched
}

// ProcessImageSlowForBenchmark runs the slow image processing path for benchmark tests.
func ProcessImageSlowForBenchmark(workerID int) bool {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())

	matched, _ := regexp.MatchString(`^image_data_\d+_timestamp_\d+$`, data)
	if matched {
		_ = strings.ToUpper(data)
	}

	return matched
}

// ProcessImageOptimizedForBenchmark runs the optimized image processing path for benchmark tests.
func ProcessImageOptimizedForBenchmark(workerID int) bool {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())

	if imageDataPattern.MatchString(data) {
		_ = strings.ToUpper(data)
		return true
	}

	return false
}

func storeBounded(key string, value []byte) {
	fixedCacheMu.Lock()
	defer fixedCacheMu.Unlock()

	if len(fixedCache) >= maxFixedCacheItems {
		for oldKey := range fixedCache {
			delete(fixedCache, oldKey)
			break
		}
	}

	fixedCache[key] = value
}
