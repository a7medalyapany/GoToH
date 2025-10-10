package main

import (
	"fmt"
	"math/rand"
)

func main() {
	age := map[string]int{
		"Ahmed": 23,
		"Menna": 20,
		"Ali":   17,
	}

	// comma idiom
	aAge := age["Ahmed"]
	fmt.Printf("Ahmed's age is: %d\n", aAge)

	rAge := age["Mariam"]
	fmt.Printf("Mariam's age is: %d\n", rAge) // so when it doesnt find the key it returns the zero value of the value type which is int here so 0

	if age, ok := age["Menna"]; ok {
		fmt.Printf("Menna's age is: %d\n", age)
	} else {
		fmt.Println("Menna's age not found")
	}

	if age, ok := age["Mariam"]; ok {
		fmt.Printf("Mariam's age is: %d\n", age)
	} else {
		fmt.Println("Mariam's age not found")
	}

	// --------------------------------------------

	for i := range 100 {
		// statement statement idiom
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("iteration is number %v and the value of x is 3\n", i)
		}
	}
}
