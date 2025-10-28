package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	iceCreamFlavors []string
}

type engine struct {
	electric bool
}

type vehicle struct {
	doors  int
	color  string
	model  string
	make   string
	engine engine
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

	fmt.Println("--------------Struct----------------")

	e1 := engine{
		electric: true,
	}

	v1 := vehicle{
		doors:  4,
		color:  "Red",
		model:  "Model S",
		make:   "Tesla",
		engine: e1,
	}

	v2 := vehicle{
		doors: 2,
		color: "Blue",
		model: "Model 6",
		make:  "Toyota",
		engine: engine{
			electric: false,
		},
	}

	fmt.Println("Vehicle 1:", v1)
	fmt.Println("Vehicle 2:", v2)

	fmt.Printf("Single field of v1 - %v & Single field of v2 - %v", v1.make, v2.model)

}
