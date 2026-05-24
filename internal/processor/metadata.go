package processor

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

var (
	// LeakCache intentionally grows forever in leaky mode.
	// It is used to demonstrate heap profiling in Lab 3.
	LeakCache = make(map[string][]byte)

	fixedCacheMu sync.Mutex
	fixedCache   = make(map[string][]byte)

	imageDataPattern = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)
)

const maxFixedCacheItems = 100

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
		LeakCache[key] = make([]byte, 1024*10) // 10 KB leaked per processed image.
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
	}

	return matched
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

func ProcessImageSlowForBenchmark(workerID int) bool {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())
	matched, _ := regexp.MatchString(`^image_data_\d+_timestamp_\d+$`, data)

	return matched
}

func ProcessImageOptimizedForBenchmark(workerID int) bool {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())

	return imageDataPattern.MatchString(data)
}
