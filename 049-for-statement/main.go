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

}
