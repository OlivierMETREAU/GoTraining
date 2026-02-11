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

func MultiplyByThree(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 3
		}
	}()
	return out
}

func SumValues(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		total := 0
		for i := range in {
			total += i
		}
		out <- total
	}()
	return out
}
