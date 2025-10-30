package main

import "fmt"

type person struct {
	name string
}

type secretAgent struct {
	person
	licenseToKill bool
}

// Simply an interface says anything calls my functions is my type
type Talker interface {
	talk() string
}

func saySomething(t Talker) {
	fmt.Println(t.talk())
}

func (p person) talk() string {
	return "Hello, my name is " + p.name
}

func (sa secretAgent) talk() string {
	return "Good evening, Agent " + sa.name
}

func main() {
	p := person{name: "John Doe"}
	sa := secretAgent{
		person:        person{name: "James Bond"},
		licenseToKill: true,
	}

	// fmt.Println(p.talk())
	// fmt.Println(sa.talk())

	saySomething(p)
	saySomething(sa)
}
