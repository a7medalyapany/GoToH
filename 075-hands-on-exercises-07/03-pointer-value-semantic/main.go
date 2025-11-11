package main

import "fmt"

type Person struct {
	name  string
}

func main() {
	p := Person{name: "Alice"}
	fmt.Println("Person's name is:", p.name)

	p = updateName(p, "Bob")
	fmt.Println("Person's name after updateName:", p.name)

	updateNamePointer(&p, "Charlie")
	fmt.Println("Person's name after updateNamePointer:", p.name)
}

func updateName(p Person, newName string) Person{
	p.name = newName
	return p
}

func updateNamePointer(p *Person, newName string) {
	p.name = newName
}