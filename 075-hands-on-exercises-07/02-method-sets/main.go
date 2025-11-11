package main

import "fmt"

type Dog struct{
	name string
}

func (d Dog) walk() {
	fmt.Printf("The dog %v is walking\n", d.name)
}

func (d *Dog) setName() {
	d.name = "Fido"
	fmt.Printf("The dog's name is now %v\n", d.name)
}

type younging interface {
	walk()
	setName()
}

func youngRun(y younging) {
	y.walk()
	y.setName()
}

func main() {

	d1 := Dog{name: "Buddy"}
	d2 := &Dog{name: "Max"}

	d1.walk()
	d1.setName()
	// youngRun(d1) // it doesnt work bc d1 is not a pointer

	d2.walk()
	d2.setName()
	youngRun(d2)

}