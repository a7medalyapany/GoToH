package main

import "fmt"

func main() {
	func() {
		for i := range 5 {
			fmt.Println(i)
		}
	}()

	x := func() {
		for i := range 3 {
			fmt.Println("Line number:", i)
		}
	}

	x()

	fmt.Println("-----------Function return a function-----------")

	y := addOne()
	fmt.Println(y(5))
}

func addOne() func(int) int {
	return func(x int) int {
		return x + 1
	}
}
