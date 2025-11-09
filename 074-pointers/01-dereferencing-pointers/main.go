package main

import "fmt"

func main() {
	s := "Hello, Go pointers!"

	p := &s

	fmt.Printf("the value of the pointer is %v\nthe value stored in it = %v\n", p, *p)
	fmt.Println(*&s) // dereferencing pointer so it prints s


	*p = s + " with more text."
	fmt.Println(s)
	fmt.Println(*p)
}