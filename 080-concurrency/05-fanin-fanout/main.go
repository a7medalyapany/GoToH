package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan string)

	// FAN-OUT: start 3 workers
	for i := range 3 {
		go worker(i, out)
	}

	// FAN-IN: collect their results
	for i := 1; i <= 3; i++ {
		fmt.Println(<-out)
	}
}

func worker(id int, out chan<- string) {
	time.Sleep(300 * time.Millisecond)
	out <- fmt.Sprintf("worker %d done", id)
}
