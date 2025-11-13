package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name  string
	Age int
}

func main() {
	u1 := user{Name: "Alice", Age: 30}
	u2 := user{Name: "Bob", Age: 25}
	fmt.Println(u1)
	fmt.Println(u2)	

	users := []user{u1, u2}
	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(data))
}
