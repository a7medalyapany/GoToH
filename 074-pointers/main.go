package main

import "fmt"

func main() {
	x := 42
	s := "Hello, Go pointers!"
	fmt.Println(x)
	fmt.Println(&x)

	// Pointer to string variable and its type
	fmt.Printf("%v\t%T", &s, &s)
}