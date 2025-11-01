package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(xi...))
}

func foo(xs ...int) int {
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return sum
}
