package main

import "fmt"

func main() {
	fmt.Println("day05-concurrency-pipeline")
}

func FilterEvenNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
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
