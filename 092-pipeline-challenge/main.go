package main

import "fmt"

func generate(nums ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, i := range nums {
			c <- i
		}
		close(c)
	}()
	return c
}


func double(in <-chan int) <-chan int {
	d := make(chan int)

	go func() {
		for n := range in {
			d <- 2 * n
		}
		close(d)
	}()

	return d
}


func main() {
	nums := []int{42, 43, 44, 45, 46}

	g := generate(nums...)
	d := double(g)

	for n := range d {
		fmt.Println(n)
	}
}