package main

import (
	"GoToH/085-testing/03-examples/01/acdc"
	"fmt"
)

func main() {
	x := acdc.Sum(1, 2, 3, 4, 5)
	fmt.Println(x)
	x = acdc.Sum(61, 8)
	fmt.Println(x)
}