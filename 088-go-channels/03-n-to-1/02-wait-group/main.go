package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		for i := range 10 {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := range 10 {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}
}
