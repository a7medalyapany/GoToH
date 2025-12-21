package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for range 10 {
		fmt.Println(<-c)
	}

	fmt.Println("You're all boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)

	for _, ch := range channels {
		go func() {
			for msg := range ch {
				out <- msg
			}
		}()
	}

	return out
}