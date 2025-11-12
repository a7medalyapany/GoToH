package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age 	 int
}
func main() {

	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:      30,
	}
	p2 := Person{
		FirstName: "Jane",
		LastName:  "Smith",
		Age:      25,
	}

	people := []Person{p1, p2}

	fmt.Println(people)

	byteSlice, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	
	fmt.Println(string(byteSlice))

}