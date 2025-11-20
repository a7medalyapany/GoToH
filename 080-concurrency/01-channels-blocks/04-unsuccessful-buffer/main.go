package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	ch <- "Hello"
	ch <- "Gophers"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
