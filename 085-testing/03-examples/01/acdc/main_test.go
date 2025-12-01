package acdc

import "fmt"

func ExampleSum() {
	fmt.Println(Sum(6, 3, 5, -7, 1))
	fmt.Println(Sum(61, 8))
	// Output:
	// 9
	// 69
}