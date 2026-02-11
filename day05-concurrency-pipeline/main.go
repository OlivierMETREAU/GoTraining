package main

import "fmt"

func main() {
	fmt.Println("day05-concurrency-pipeline")
	FilterEvenNumbers([]int{1, 2, 3, 4, 5})
}

func FilterEvenNumbers(l []int) []int {
	evenNumbers := []int{}
	for _, n := range l {
		if n%2 == 0 {
			evenNumbers = append(evenNumbers, n)
		}
	}
	return evenNumbers
}

func MultiplyByThree(numbers []int) {
	for i := range numbers {
		numbers[i] *= 3
	}
}

func SumValues(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}
