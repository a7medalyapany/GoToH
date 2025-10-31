package main

import "fmt"

func main() {
	x := factorial(5)
	fmt.Println("Factorial 5 = ", x)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
