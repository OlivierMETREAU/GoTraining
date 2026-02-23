package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"

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

func TestSetWithTTLWithEmptyKey(t *testing.T) {
	c := New()
	err := c.SetWithTTL("", 0, time.Duration(10*float64(time.Second)))
	assert.NotEqual(t, nil, err)
	assert.NotContains(t, c.data, "")
	assert.Empty(t, c.data)
}

func TestGetBeforeExpiration(t *testing.T) {
	c := New()
	c.SetWithTTL("John", "Doe", time.Duration(float64(time.Second)))
	time.Sleep(time.Duration(0.9 * float64(time.Second)))
	v, ok := c.Get("John")
	assert.Equal(t, "Doe", v)
	assert.Equal(t, true, ok)
}

func TestGetAfterExpiration(t *testing.T) {
	c := New()
	c.SetWithTTL("John", "Doe", time.Duration(float64(time.Second)))
	time.Sleep(time.Duration(1.1 * float64(time.Second)))
	v, ok := c.Get("John")
	assert.Equal(t, nil, v)
	assert.Equal(t, false, ok)
}
