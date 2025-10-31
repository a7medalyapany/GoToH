package main

import "fmt"

func main() {

	x := foo()
	y := bar()

	fmt.Printf("the type of x is %T & the value of it is %v\n", x, x)
	fmt.Printf("the type of y is %T & the value of it is %v\n", y, y())
	fmt.Printf("the type of bar is %T\n", bar)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 8
	}
}
