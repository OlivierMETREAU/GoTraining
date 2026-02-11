package main

import "fmt"

func main() {
	fmt.Println("day05-concurrency-pipeline")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	input := make(chan int)
	step1 := FilterEvenNumbers(input)
	step2 := MultiplyByThree(step1)
	step3 := SumValues(step2)

	go func() {
		defer close(input)
		for _, n := range numbers {
			input <- n
		}
	}()

	total := <-step3
	fmt.Println("Input list : ")
	fmt.Println(numbers)
	fmt.Printf("The sum of the even numbers multiplied by 3 is %d\n", total)

	// {1, 2, 3, 4, 5, 6, 7, 8, 9}
	// step1: {2,4,6,8}
	// step2: {6,12,18,24}
	// step3: 60
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
