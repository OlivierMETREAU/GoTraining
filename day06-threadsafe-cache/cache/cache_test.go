package cache

import (
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
