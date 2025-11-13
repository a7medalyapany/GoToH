package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"Name"`
	Age int   `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	u1 := User{
		Name: "Alice",
		Age:  30,
		Sayings: []string{
			"Hello, world!",
			"Go is awesome.",
		},
	}
	
	u2 := User{
		Name: "Bob",
		Age:  25,
		Sayings: []string{
			"Good morning!",
			"Have a great day!",
		},
	}

	users := []User{u1, u2}

	fmt.Printf("Users: %+v\n", users)

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	
}