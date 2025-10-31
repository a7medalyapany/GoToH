package main

import "fmt"

func main() {
	foo()

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(name string) {
		fmt.Println("My name is, " + name)
	}

	x()
	y("Alyapany")
}

func foo() {
	fmt.Println("foo ran")
}
