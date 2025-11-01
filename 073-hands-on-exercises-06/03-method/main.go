package main

import "fmt"

type person struct {
	firstName string
	age       int
}

func (p person) speak() {
	fmt.Printf("my name is %v & im %v years old", p.firstName, p.age)
}

func main() {
	p := person{
		firstName: "Alyapany",
		age:       23,
	}

	p.speak()
}
