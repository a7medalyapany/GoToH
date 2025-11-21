package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c) // comment that and see what happen
	}()

	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}
