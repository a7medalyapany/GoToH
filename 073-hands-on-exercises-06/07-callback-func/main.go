package main

import "fmt"

func main() {
	sq := printSquare(square, 6)
	fmt.Println(sq)
}

func square(x int) int {
	return x * x
}

func printSquare(sq func(int) int, x int) string {
	a := sq(x)
	return fmt.Sprintf("the squared number of %v is %v", x, a)
}
