package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		for i := range 10 {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := range 10 {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done // this is gonna wait till it runs this line`done <- true`
		<-done
		close(c)
	}()
	for n := range c {
		fmt.Println(n)
	}
}
