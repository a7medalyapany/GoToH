package main

import "fmt"

func main() {
	func() {
		for i := range 5 {
			fmt.Println(i)
		}
	}()
}
