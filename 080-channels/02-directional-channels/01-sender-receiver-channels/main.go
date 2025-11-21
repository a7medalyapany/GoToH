package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Printf("c  type \t%T\n", c)
	fmt.Printf("cr type \t%T\n", cr)
	fmt.Printf("cs type \t%T\n", cs)

	// specific to general doesnt work
	// c = cr
	// c = cs
	// fmt.Printf("cr\t%T\n", (chan int)(cr))
	// fmt.Printf("cs\t%T\n", (chan int)(cs))

	// general to specific assigns
	cr = c
	cs = c
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))

}
