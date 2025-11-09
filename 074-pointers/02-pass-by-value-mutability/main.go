package main

import "fmt"

// What we just did was:
// 1. Take the address of a (using &a)
// 2. Pass that address into the intDelta function
// 3. Inside intDelta, we dereferenced the pointer (using *n) to set the value at that address to 43
func intDelta(n *int) { //*int means a pointer to an int TYPE
	*n = 43
}

func sliceDelta(s []int) { // slices are reference types by default
	s[0] = 100
}

func mapDelta(m map[string]int,s string) {
	m[s] = 100
}

func main() {
	a := 42
	intDelta(&a)
	fmt.Println(a)


	xi := []int{1,2,3,4,5}
	fmt.Println("Before:", xi)
	sliceDelta(xi)
	fmt.Println("After:", xi)

	ms := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("Before:", ms)
	mapDelta(ms,"foo")
	fmt.Println("After:", ms)

}