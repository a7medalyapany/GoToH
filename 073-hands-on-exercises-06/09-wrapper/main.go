package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

func doWork() {
	for i := range 10 {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	startTime := time.Now()
	f()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Function took %s to execute\n", elapsedTime)
}
