package main

import (
	"fmt"
)

func main() {
	for i := range 100 {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v is an odd number\n", i)
	}

	xi := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, v := range xi {
		fmt.Printf("the index of the slice is %v and the value is %v\n", i, v)
	}

	m := map[string]int{
		"Ahmed": 3,
		"&":     8,
		"Menna": 1,
	}

	for k, v := range m {
		fmt.Printf("the string is %v and the value is %v\n", k, v)
	}
}
