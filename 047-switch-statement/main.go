package main

import "fmt"

func main() {
	x := 1
	switch {
	case x == 1:
		fmt.Println("x is 1")
	case x == 2:
		fmt.Println("x is 2")
	case x == 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("idk about x")
	}

	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("idk about x")
	}

	switch x {
	case 1:
		fmt.Println("x is 1")
		fallthrough
	case 2:
		fmt.Println("printed because of fallthrough: x is 2")
	case 3:
		fmt.Println("printed because of fallthrough: x is 3")
	default:
		fmt.Println("printed because of fallthrough: idk about x")
	}

	fmt.Println("the final value of x is: ", x)

	switch x {
	case 1:
		fmt.Println("x is 1")
		fallthrough
	case 2:
		fmt.Println("printed because of fallthrough: x is 2")
		fallthrough
	case 3:
		fmt.Println("printed because of fallthrough: x is 3")
	default:
		fmt.Println("printed because of fallthrough: idk about x")
	}

	fmt.Println("the final value of x is: ", x)

}
