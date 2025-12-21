package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)

	fmt.Println(<-cSum)
}

func incrementor() chan int {
	out := make(chan int)

	go func() {
		for i := 10; i <= 15; i++ {
			out <- i
		}
		close(out)
	}()
	return out

}

func puller(c chan int) chan int {
	out := make(chan int)

	var sum int

	go func() {
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()

	return out
}
