package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
	Saying []string
}

func main() {
	p1 := person{
		Name: "James Bond",
		Age:  32,
		Saying: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In the world of espionage, nothing is certain but death",
		},
	}

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	// Simulating an error for demonstration purposes
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Error inside json.Marshal: %v", err)
	}
	return bs, nil
}
