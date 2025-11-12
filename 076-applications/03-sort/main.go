package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 2, 4, 3, 1, 16, 32, 23, 69, 11}
	xs := []string{"banana", "apple", "orange", "kiwi", "grape"}

	fmt.Println("Before sorting:", xi)
	sort.Ints(xi)
	fmt.Println("After sorting:", xi)

	fmt.Println("Before sorting:", xs)
	sort.Strings(xs)
	fmt.Println("After sorting:", xs)
}