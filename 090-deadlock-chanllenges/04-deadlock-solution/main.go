package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := range 5 {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
