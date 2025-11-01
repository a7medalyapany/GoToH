package main

import "fmt"

func main() {

	defer foo()
	bar()

	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)

}

func foo() {
	fmt.Println("foo is supposed to run at first")
}

func bar() {
	fmt.Println("bar is supposed to run at last")
}

// deferred functions run in LIFO order
