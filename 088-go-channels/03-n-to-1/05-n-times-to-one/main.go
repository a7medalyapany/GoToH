package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for range n {
		go func() {
			for i := range 10 {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for range n {
			<-done
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}
