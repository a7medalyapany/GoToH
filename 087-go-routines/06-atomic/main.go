package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go Incrementor("Foo:")
	go Incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func Incrementor(s string) {
	for i := range 10 {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
