package processor

import "testing"

var benchmarkResult bool

func BenchmarkProcessImageSlow(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = ProcessImageSlowForBenchmark(i)
	}
	benchmarkResult = result
}

func BenchmarkProcessImageOptimized(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = ProcessImageOptimizedForBenchmark(i)
	}
	benchmarkResult = result
}
