package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	fmt.Println("day05-concurrency-pipeline")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	input := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	verbose := true
	step1 := FilterEvenNumbers(ctx, input, verbose)
	step2 := MultiplyByThree(ctx, 4, step1, verbose)
	step3 := SumValues(ctx, step2, verbose)

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

func FilterEvenNumbers(ctx context.Context, in <-chan int, verbose bool) <-chan int {
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
					if verbose {
						log.Printf("filterEven: %d is even", n)
					}
					out <- n
				}
			}
		}
	}()
	return out
}

func MultiplyByThree(ctx context.Context, workers int, in <-chan int, verbose bool) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case n, ok := <-in:
					if !ok {
						return
					}
					if verbose {
						log.Printf("worker %d: multiply %d → %d", id, n, n*3)
					}
					out <- n * 3
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func SumValues(ctx context.Context, in <-chan int, verbose bool) <-chan int {
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
					if verbose {
						log.Printf("sum: final total = %d", total)
					}
					out <- total
					return
				}
				total += n
				if verbose {
					log.Printf("sum: added %d → total %d", n, total)
				}
			}
		}
	}()
	return out
}
