package main

import (
	"fmt"
)

func addOne(n int) int {
	return n + 1
}

func addOnePtr(n *int){
	*n = *n + 1
}

func main() {
	// Demonstrating value semantics with an integer
	x := 10
	fmt.Println("Initial value of x: ", x)
	fmt.Println("After adding one to x: ", addOne(x))
	fmt.Println("Value of x after function call: ", x)

	// Demonstrating pointer semantics with an integer
	y := 20
	fmt.Println("Initial value of y: ", y)
	addOnePtr(&y)
	fmt.Println("Value of y after function call: ", y)	
}