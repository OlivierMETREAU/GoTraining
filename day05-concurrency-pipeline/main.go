package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("day05-concurrency-pipeline")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	input := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	step1 := FilterEvenNumbers(ctx, input)
	step2 := MultiplyByThree(ctx, step1)
	step3 := SumValues(ctx, step2)

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

func FilterEvenNumbers(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case n, ok := <-in:
				if !ok {
					return
				}
				if n%2 == 0 {
					out <- n
				}
			}
		}
	}()
	return out
}

func MultiplyByThree(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case n, ok := <-in:
				if !ok {
					return
				}
				out <- n * 3
			}
		}
	}()
	return out
}

func SumValues(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		total := 0
		for {
			select {
			case <-ctx.Done():
				return
			case n, ok := <-in:
				if !ok {
					out <- total
					return
				}
				total += n
			}
		}
	}()
	return out
}
