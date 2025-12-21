package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := range 10 {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := range 10 {
			c <- i
		}
		wg.Done()
	}()

	for n := range c {
		fmt.Println(n)
	}
}
