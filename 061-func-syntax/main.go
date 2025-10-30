package main

import "fmt"

/*
func syntax:
 -- no params, no returns
 -- no params, with returns
 -- with params, no returns
 -- with params, with returns
*/

func foo() {
	fmt.Println("foo called")
}

func bar() string {
	return "bar called"
}

func baz(s string) {
	fmt.Println("baz called with", s)
}

func qux(name string, age int) (string, int) {
	return fmt.Sprintf("qux called with name: %v", name), age
}
func main() {
	foo()

	resultBar := bar()
	fmt.Println(resultBar)

	baz("a parameter for baz")

	resultQux, age := qux("Alice", 30)
	fmt.Println(resultQux)
	fmt.Println("Age returned from qux:", age)
}
