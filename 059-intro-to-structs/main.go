package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	fmt.Println("First Name:", p1.firstName)
	fmt.Println("Last Name:", p1.lastName)
	fmt.Println("Age:", p1.age)

	p2 := person{"Jane", "Smith", 25}

	fmt.Println("First Name:", p2.firstName)
	fmt.Println("Last Name:", p2.lastName)
	fmt.Println("Age:", p2.age)

	fmt.Printf("Person 1: %+v\n", p1)
	fmt.Printf("Person 2: %+v\n", p2)
}
