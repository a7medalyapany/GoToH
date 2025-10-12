package main

import "fmt"

func main() {
	a := [3]int{11, 22, 33}
	b := [...]string{"ahmed", "menna"}
	fmt.Println(a, b)

	var c [3]int

	fmt.Println(c)
	c[0] = 23
	c[1] = 20
	fmt.Println(c)

	fmt.Println(len(c))

	// vs code keyboard shortcuts to place multiple cursors
	// alt				click
	// alt+shift 		drag
	// alt+shift+i		cursors at the end of each line for selection
}
