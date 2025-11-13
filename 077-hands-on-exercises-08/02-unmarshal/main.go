package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"Name"`
	Age int `json:"Age"`
}

func main() {

	str := `[{"Name":"Alice","Age":30},{"Name":"Bob","Age":25}]`
	sliceByte := []byte(str)

	var users []User

	err := json.Unmarshal(sliceByte, &users)

	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Printf("Unmarshaled Users: %+v\n", users)
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)
	}
}