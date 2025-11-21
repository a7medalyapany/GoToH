package main

import "fmt"

func main() {

	c := make(chan int)

	// put numbers
	gen(c)

	// get numbers
	receive(c)

	fmt.Println("about to exit")
}

func gen(cs chan<- int) {
	go func() {
		for i := range 10 {
			cs <- i
		}
		close(cs)
	}()
}

func receive(cr <-chan int) {
	for v := range cr {
		fmt.Println("cr =", v)
	}
}
