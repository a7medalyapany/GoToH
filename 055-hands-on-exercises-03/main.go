package main

import (
	"fmt"
)

func main() {
	x := []int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53,54, 55)
	fmt.Println(x)

	y := []int {56, 57, 58, 59, 60}

	x = append(x, y...)
	fmt.Println(x)
	
	
	fmt.Println("------------------")
	// deleting and slicing
	x = append(x[:3], x[6:10]...)
	fmt.Println(x)


	fmt.Println("------------------")
	states := make([]string, 0, 20)
	states = append(states,` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	
	fmt.Printf("the len - %v\nthe cap - %v\nand the slice holds -  %v\n", len(states), cap(states), states)

	
	fmt.Println("------------------")
}
