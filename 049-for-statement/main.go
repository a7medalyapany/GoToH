package main

import "fmt"

func main() {
	x := 5
	y := 10

	// for i := 0; i < 5; i++ -- is the same to the line below
	for i := range 5 {
		fmt.Printf("counting to five: %v \t first for loop\n", i)
	}

	for y < 13 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	for {
		fmt.Printf("x is %v \t\t third loop\n", x)
		if x > 8 {
			break
		}
		x++
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}

	// nested loops

	for i := range 3 {

		for j := range 5 {
			fmt.Printf("the outer loop is %v \t\t\t the inner loop is %v\n", i, j)
		}
	}

	// range loops

	xi := []int{30, 40, 50, 60, 70, 80} //slice
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	m := map[string]int{
		"Ahmed": 8,
		"Ali":   14,
	}

	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}
}
