package main

import (
	"fmt"
	"sort"
)


func main() {
	xi := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(xi)
	fmt.Printf("Sorted integers: %v\n", xi)

	xs := []string{"banana", "apple", "cherry"}
	sort.Strings(xs)
	fmt.Printf("Sorted strings: %v\n", xs)
}