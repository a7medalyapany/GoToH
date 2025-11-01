package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	fmt.Println(doMath(10, 5, add))
	fmt.Println(doMath(10, 5, subtract))
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func doMath(a, b int, operatoin func(int, int) int) int {
	return operatoin(a, b)
}
