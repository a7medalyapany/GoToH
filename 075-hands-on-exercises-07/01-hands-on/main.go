package main

import "fmt"

var (
	a, b, c *string
	d 	  *int
)

func init() {
	p := "Hello, Go!"
	q := "Pointers"
	r := "are powerful"
	n:= 2002

	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)

	fmt.Println("Value of a:", *a)
	fmt.Println("Value of b:", *b)
	fmt.Println("Value of c:", *c)
	fmt.Println("Value of d:", *d)
}