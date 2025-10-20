package main

import "fmt"

func main() {
	a := []int{11, 22, 33}
	b := []string{"ahmed", "menna"}
	fmt.Println(a, b)

	var s []int                                 // s is nil, len(s) == 0
	fmt.Println("nil?", s == nil)               // true
	fmt.Println("len:", len(s), "cap:", cap(s)) // 0, 0

	// Uncommenting the next line will panic:
	// fmt.Println(s[0]) // panic: runtime error: index out of range [0] with length 0

	if len(s) > 0 {
		fmt.Println("first element:", s[0])
	} else {
		fmt.Println("cannot index: slice is empty")
	}

	// Safe ways to add/read
	s = append(s, 42)                  // append works on nil slices
	fmt.Println("after append:", s[0]) // prints 42

	t := make([]int, 3)               // allocated with len 3
	fmt.Println("t[0] exists:", t[0]) // safe, zero value 0

	for i, v := range []int{20, 54, 23, 67, 89} {
		fmt.Println("at index", i, "value", v)
	}
	
	xs := []string{"ahmed", "menna", "ali", "sara", "mariam"}
	for i := 0; i < len(xs); i++ {
		fmt.Println("at index", i, "value", xs[i])
	}

	fmt.Println("--------------")
	// Alternative using range
	for i, v := range xs {
		fmt.Println("at index", i, "value", v)
	}
	
	si := []int{10, 20, 30, 40, 50}
	fmt.Println("si before appends:", si)
	
	si = append(si, 60, 70)
	fmt.Println("si after appends:", si)
	
	
	fmt.Println("--------------")
	// Slicing operations

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// [inclusive:exclusive]
	fmt.Printf("slice - %v\n", slice[2:5]) // 2,3,4
	// [:exclusive]
	fmt.Printf("slice - %v\n", slice[:6]) // slice[:0] -> empty slice
	// [inclusive:]
	fmt.Printf("slice - %#v\n", slice[7:]) // from index 7 to end
	// [:]
	fmt.Printf("slice - %#v\n", slice[:]) // entire slice
}
