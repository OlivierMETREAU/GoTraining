package cache

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetWithEmptyString(t *testing.T) {
	c := New()
	err := c.Set("", 0)
	assert.NotEqual(t, nil, err)
	assert.NotContains(t, c.data, "")
	assert.Empty(t, c.data)
}

func TestSetWithNonEmptyString(t *testing.T) {
	c := New()
	err := c.Set("John", "Doe")
	assert.Equal(t, nil, err)
	assert.Contains(t, c.data, "John")
	assert.NotEmpty(t, c.data)
}

func TestGetWithUnknownKey(t *testing.T) {
	c := New()
	v, ok := c.Get("John")
	assert.Equal(t, nil, v)
	assert.Equal(t, false, ok)
}

func TestGetWithKnownKey(t *testing.T) {
	c := New()
	c.Set("John", "Doe")
	v, ok := c.Get("John")
	assert.Equal(t, "Doe", v)
	assert.Equal(t, true, ok)
}

func TestCacheConcurrentAccess(t *testing.T) {
	c := New()

	var wg sync.WaitGroup

	// Start 100 writers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			c.Set(key, i)
		}(i)
	}

	// Start 100 readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			c.Get(key)
		}(i)
	}

	wg.Wait()
}
