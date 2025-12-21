package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := range 15 {
			c <- i
		}
		close(c)
	}()

	for range n {
		go func() {
			for v := range c {
				fmt.Println(v)
			}
			done <- true
		}()
	}

	for range n {
		<-done
	}
}
