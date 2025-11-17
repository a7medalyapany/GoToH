package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go foo()
	go bar()
	go func() {
		fmt.Println("anonymos func number 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("anonymos func number 2")
		wg.Done()
	}()
	wg.Wait()
}

func foo() {
	fmt.Println("foo goroutine")
	wg.Done()
}

func bar() {
	fmt.Println("bar goroutine")
	wg.Done()
}
