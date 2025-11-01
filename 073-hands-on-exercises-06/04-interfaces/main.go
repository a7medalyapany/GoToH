package main

import (
	"fmt"
	"math"
)

type square struct {
	length int
	width  int
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return float64(s.length * s.width)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	sq := square{length: 5, width: 10}
	ci := circle{radius: 7.5}

	sqArea := getArea(sq)
	ciArea := getArea(ci)

	fmt.Printf("square area = %v\ncircle area = %v", sqArea, ciArea)
}
