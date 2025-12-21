package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Number   int
	Factorial int
}

func main() {
	nums := generateNumbers()
	fac := factorial(nums)

	for result := range fac {
		fmt.Printf("Factorial of %d: %d\n", result.Number, result.Factorial)
	}
}

func generateNumbers() <-chan int {
	nums := make(chan int)
	go func() {
		for i := range 10 {
			nums <- i
		}
		close(nums)
	}()
	return nums
}

func factorial(in <-chan int) <-chan Result {
	out := make(chan Result)
	go func() {
		var wg sync.WaitGroup
		for n := range	 in {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				out <- Result{Number: n, Factorial: fact(n)}
			}(n)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := range n {
		total *= i + 1
	}
	return total
}