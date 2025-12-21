package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := range 5 {
			c <- i
		}
	}()
	fmt.Println(<-c)
}
