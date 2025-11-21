package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)

	gen(c)
	fmt.Println("Goroutines after gen():", runtime.NumGoroutine())

	receive(c)
	fmt.Println("Goroutines after receive():", runtime.NumGoroutine())

	fmt.Println("about to exit")
}

func gen(c chan<- int) {
	for range 10 {
		go func() {
			for i := range 10 {
				c <- i
			}
		}()
	}
}

func receive(cr <-chan int) {
	for i := range 100 {
		fmt.Printf("%d \t%v\n", i+1, <-cr)
	}
}
