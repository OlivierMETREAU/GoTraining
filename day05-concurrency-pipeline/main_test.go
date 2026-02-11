package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterEvenNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	input := make(chan int)
	out := FilterEvenNumbers(input)

	go func() {
		defer close(input)
		for _, n := range numbers {
			input <- n
		}
	}()

	expectedEvenNumbers := []int{2, 4, 6, 8}
	var evenNumbers []int
	for n := range out {
		evenNumbers = append(evenNumbers, n)
	}
	assert.Equal(t, expectedEvenNumbers, evenNumbers)
}

func TestMultiplsByThree(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expectedEvenNumbers := []int{3, 6, 9, 12, 15}
	input := make(chan int)
	out := MultiplyByThree(input)

	go func() {
		defer close(input)
		for _, n := range numbers {
			input <- n
		}
	}()
	var multipliedOutput []int
	for n := range out {
		multipliedOutput = append(multipliedOutput, n)
	}
	assert.Equal(t, expectedEvenNumbers, multipliedOutput)
}

func TestSumValues(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expectedOutput := 15
	input := make(chan int)
	out := SumValues(input)

	go func() {
		defer close(input)
		for _, n := range numbers {
			input <- n
		}
	}()

	got := <-out
	assert.Equal(t, expectedOutput, got)
}
