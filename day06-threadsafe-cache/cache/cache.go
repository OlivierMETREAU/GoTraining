package cache

import (
	"errors"
	"sync"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]any
}

func New() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) error {
	if key == "" {
		return errors.New("Empty string as keyvalue is not allowed.")
	}

	c.mu.Lock() // exclusive lock
	defer c.mu.Unlock()

	c.data[key] = value
	return nil
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock() // shared read lock
	defer c.mu.RUnlock()

	v, ok := c.data[key]
	return v, ok
}
