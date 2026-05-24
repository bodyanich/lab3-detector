package stats

import "sync"

type UnsafeCounter struct {
	values map[string]int
}

func NewUnsafeCounter() *UnsafeCounter {
	return &UnsafeCounter{values: make(map[string]int)}
}

func (c *UnsafeCounter) IncrementProcessed(imageType string) {
	c.values[imageType]++
}

func (c *UnsafeCounter) Value(imageType string) int {
	return c.values[imageType]
}

type SafeCounter struct {
	mu     sync.RWMutex
	values map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{values: make(map[string]int)}
}

func (c *SafeCounter) IncrementProcessed(imageType string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[imageType]++
}

func (c *SafeCounter) Value(imageType string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.values[imageType]
}
