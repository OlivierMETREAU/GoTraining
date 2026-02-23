package cache

import (
	"errors"
	"sync"
	"time"
)

type item struct {
	value     any
	expiresAt time.Time // zero means "no expiration"
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]item
}

func New() *Cache {
	return &Cache{
		data: make(map[string]item),
	}
}

func (c *Cache) SetWithTTL(key string, value any, ttl time.Duration) error {
	if key == "" {
		return errors.New("Empty string as keyvalue is not allowed.")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	var expires time.Time
	if ttl > 0 {
		expires = time.Now().Add(ttl)
	}

	c.data[key] = item{
		value:     value,
		expiresAt: expires,
	}

	return nil
}

func (c *Cache) Set(key string, value any) error {
	return c.SetWithTTL(key, value, 0)
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock() // shared read lock
	defer c.mu.RUnlock()

	v, ok := c.data[key]
	return v.value, ok
}
