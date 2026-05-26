package processor

import "testing"

func BenchmarkProcessImageSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessImageSlowForBenchmark(i)
	}
}

func BenchmarkProcessImageOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessImageOptimizedForBenchmark(i)
	}
}
