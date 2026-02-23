package cache

import "errors"

type Cache struct {
	data map[string]any
}

func New() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) error {
	if key != "" {
		c.data[key] = value
		return nil

	}
	return errors.New("Empty string as keyvalue is not allowed.")
}

func (c *Cache) Get(key string) (any, bool) {
	v, ok := c.data[key]
	return v, ok
}
