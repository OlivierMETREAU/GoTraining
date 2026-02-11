package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterEvenNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedEvenNumbers := []int{2, 4, 6, 8}
	evenNumbers := FilterEvenNumbers(numbers)
	assert.Equal(t, expectedEvenNumbers, evenNumbers)
}
