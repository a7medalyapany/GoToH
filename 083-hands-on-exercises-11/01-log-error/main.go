package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatal("Error inside json.Marshal:", err)
	}

	fmt.Println(string(bs))
}