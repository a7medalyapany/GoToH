package main

import "fmt"

func main() {
	func() {
		for i := range 5 {
			fmt.Println(i)
		}
	}()

	x := func() {
		for i := range 3 {
			fmt.Println("Line number:", i)
		}
	}

	x()
}
