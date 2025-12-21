package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := range 10 {
		fmt.Println("foo -", i)
	}
	wg.Done()
}

func bar() {
	for i := 10; i < 20; i++ {
		fmt.Println("bar -", i)
	}
	wg.Done()
}
