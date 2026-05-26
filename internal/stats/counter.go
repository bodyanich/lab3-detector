// Package stats contains safe and unsafe counters used for race detector experiments.
package stats

import "sync"

// UnsafeCounter stores image processing statistics without synchronization.
type UnsafeCounter struct {
	values map[string]int
}

// NewUnsafeCounter creates a new unsafe counter for race detector demonstration.
func NewUnsafeCounter() *UnsafeCounter {
	return &UnsafeCounter{values: make(map[string]int)}
}

// IncrementProcessed increments the processed counter for the given image type without synchronization.
func (c *UnsafeCounter) IncrementProcessed(imageType string) {
	c.values[imageType]++
}

// Value returns the current counter value for the given image type.
func (c *UnsafeCounter) Value(imageType string) int {
	return c.values[imageType]
}

// SafeCounter stores image processing statistics using a mutex.
type SafeCounter struct {
	mu     sync.RWMutex
	values map[string]int
}

// NewSafeCounter creates a new synchronized counter.
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{values: make(map[string]int)}
}

// IncrementProcessed increments the processed counter for the given image type using synchronization.
func (c *SafeCounter) IncrementProcessed(imageType string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[imageType]++
}

// Value returns the current synchronized counter value for the given image type.
func (c *SafeCounter) Value(imageType string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.values[imageType]
}
