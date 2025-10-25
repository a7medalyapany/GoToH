package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	agent := secretAgent{
		person:        p1,
		licenseToKill: true,
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

	fmt.Println("-------------Embedded struct---------------")

	fmt.Println("Agent First Name:", agent.firstName)
	fmt.Println("Agent Last Name:", agent.lastName)
	fmt.Println("Agent Age:", agent.age)
	fmt.Println("License to Kill:", agent.licenseToKill)
	fmt.Printf("Secret Agent: %+v\n", agent)

	fmt.Println("Accessing embedded struct directly:", agent.person)
}
