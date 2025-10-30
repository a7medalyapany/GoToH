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

// Variadic function example
func sum(numbers ...int) int {
	fmt.Printf("numbers - %v & type of it is %T", numbers, numbers)

	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	foo()

	resultBar := bar()
	fmt.Println(resultBar)

	baz("a parameter for baz")

	resultQux, age := qux("Alice", 30)
	fmt.Println(resultQux)
	fmt.Println("Age returned from qux:", age)

	fmt.Println("------------Variadic Parameters------------")

	total := sum(23, 9, 4, 13, 8, 1, 6, 5)

	fmt.Println("\nTotal sum is:", total)
}
