package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

func main() {

	p1 := person{
		firstName:       "Bob",
		lastName:        "Brown",
		iceCreamFlavors: []string{"Vanilla", "Chocolate", "Strawberry"},
	}

	fmt.Println("the person struct is - ", p1)
	fmt.Println("Ice Cream Flavors:", p1.iceCreamFlavors)
	fmt.Printf("the type of the p1 is %T & the type of the iceCreamFlavors variable is %T\n", p1, p1.iceCreamFlavors)

	for i, flavor := range p1.iceCreamFlavors {
		fmt.Printf("the ice cream flavor number %v is %v\n", i+1, flavor)
	}

	fmt.Println("------------------------------")

	mp := map[string]person{
		p1.lastName: p1,
	}

	for k, v := range mp {
		fmt.Println("Key:", k)
		fmt.Println("Value:", v)
	}
}
