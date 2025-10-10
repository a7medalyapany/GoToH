package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)

	fmt.Printf("the value of x is %v \t", x)

	switch {
	case x <= 100:
		fmt.Println("x is btwn 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("x is btwn 101 and 200")
	case x > 200:
		fmt.Println("x is btwn 201 and 250")
	}
}
