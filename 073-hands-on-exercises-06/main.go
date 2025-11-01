package main

import "fmt"

func main() {
	fmt.Println(foo())
	i, s := bar()
	fmt.Println(i, s)
}

func foo() int {
	return 42
}

func bar() (x int, y string) {
	x = 23
	y = "alyapany"

	return
}
