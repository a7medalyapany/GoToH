package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", doMath)

	x := add(5, 8)
	y := sub(11, 3)
	z := doMath(83, 14, sub)

	fmt.Printf("Add result = %v\nSub result = %v\ndoMath = %v", x, y, z)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
