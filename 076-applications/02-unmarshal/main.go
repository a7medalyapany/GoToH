package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int	 `json:"Age"`
}

func main() {
	s := `[{"FirstName":"John","LastName":"Doe","Age":30},{"FirstName":"Jane","LastName":"Smith","Age":25}]`
	byteSlice := []byte(s)

	fmt.Printf("type of s is %T\ntype of byteSlice is %T\n", s, byteSlice)

	// people := []Person{}
	var people []Person
	err := json.Unmarshal(byteSlice, &people)

	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Printf("Unmarshaled data: %+v\n", people)

	for i, person := range people {
		fmt.Printf("Person %d: %+v\n", i+1, person)

	}
}