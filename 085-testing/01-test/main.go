package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5)
	fmt.Println(x)
}

func sum(xi ...int) int {
	total := 0 
	for i := range xi {
		total += xi[i]
	}
	return total
}