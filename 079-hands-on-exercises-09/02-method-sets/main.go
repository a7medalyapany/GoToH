package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func (p *Person) speak() {
	fmt.Println("Im speaking inside speak method")
}

func main() {
	p1 := Person{
		Name: "Ahmed",
		Age:  23,
	}

	saySomething(&p1)
	// saySomething(p1) // This line would cause a compile-time error because p1 is not a pointer.
}
