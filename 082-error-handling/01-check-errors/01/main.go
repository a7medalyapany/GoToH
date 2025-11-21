package main

import "fmt"

func main() {
	n, err := fmt.Println("alyapany")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n) // 9 because of new-line character
}
