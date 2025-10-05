package main

import (
	"fmt"

	"github.com/a7medalyapany/puppy"
)

func main() {
	barks := puppy.Barks()
	fmt.Println(puppy.Bark())
	fmt.Println(barks)

	fmt.Println("Big Bark: " + puppy.BigBark())

	fmt.Println("Big Barks: " + puppy.BigBarks())
}
