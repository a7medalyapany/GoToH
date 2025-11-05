package main

import (
	"math"
)

func main() {
	x := powinator(3)
	println(x())
	println(x())
	println(x())
	println(x())
	println(x())
}

func powinator(a float64) func() int {
	c := 0

	return func() int {
		c++
		return int(math.Pow(a, float64(c)))
	}
}
