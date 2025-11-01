package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 42
}

func bar() (x int, y string) {
	x = 23
	y = "alyapany"

	return
}
