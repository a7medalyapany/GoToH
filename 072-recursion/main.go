package main

import "fmt"

func main() {
	x := factorial(5)
	fmt.Println("Factorial 5 = ", x)

	y := factLoop(4)
	fmt.Println("Factorial 4 = ", y)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
