package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

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
		time.Sleep(time.Millisecond * 2)
	}
	wg.Done()
}

func bar() {
	for i := 10; i < 20; i++ {
		fmt.Println("bar -", i)
		time.Sleep(time.Millisecond * 3)
	}
	wg.Done()
}
