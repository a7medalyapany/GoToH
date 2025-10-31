package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(name string) {
		fmt.Println("My name is, " + name)
	}("alyapany")
}

func foo() {
	fmt.Println("foo ran")
}
