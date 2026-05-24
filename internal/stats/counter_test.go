package stats

import (
	"os"
	"sync"
	"testing"
)

func TestUnsafeCounterRace(t *testing.T) {
	if os.Getenv("RUN_UNSAFE_RACE") != "1" {
		t.Skip("set RUN_UNSAFE_RACE=1 to demonstrate the race detector failure")
	}

	counter := NewUnsafeCounter()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.IncrementProcessed("jpg")
			}
		}()
	}

	wg.Wait()
}

func TestSafeCounterNoRace(t *testing.T) {
	counter := NewSafeCounter()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.IncrementProcessed("jpg")
			}
		}()
	}

	wg.Wait()

	if got := counter.Value("jpg"); got != 100000 {
		t.Fatalf("counter.Value(jpg) = %d; want 100000", got)
	}
}
