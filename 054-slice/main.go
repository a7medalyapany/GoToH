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
}
