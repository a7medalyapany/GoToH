package main

import "fmt"

func main() {
	c := factorial(5)

	fmt.Println(<-c)
}

func factorial(n int) chan int {
	fac := make(chan int)

	go func() {
		total := 1
		for i := range n {
			total *= i + 1
		}
		fac <- total
		close(fac)
	}()

	return fac
}

// If i have a lot of factorial to calculate
// putting them in a go routine would make it a lot faster
