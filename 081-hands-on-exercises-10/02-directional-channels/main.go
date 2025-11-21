package main

import "fmt"

func main() {
	// cs := make(chan<- int) // commented out to avoid compile-time error
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c) // this will cause a compile-time error

	fmt.Printf("--------------\n")
	fmt.Printf("c\t%T\n", c)
}
