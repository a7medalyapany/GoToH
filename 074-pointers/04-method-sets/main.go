package main

import "fmt"

type Car struct {
	brand string
}

func (c Car) Drive() {
	fmt.Println("Driving a", c.brand)
}

func (c *Car) Refuel() {
	fmt.Println("Refueling a", c.brand)
}
func main() {

	car1 := Car{brand: "Toyota"}
	car2 := &Car{brand: "Honda"}

	car1.Drive()   // ok
    // car1.Refuel() // ❌ can't call, needs pointer => under the hood Go does: (&car1).Refuel()
	// Car{brand: "BMW"}.Refuel() // ❌ can't call, needs pointer

    car2.Drive()   // ok
    car2.Refuel()  // ok
}

// So in summary:
// - Value receivers can be called on both values and pointers.
// - Pointer receivers can only be called on pointers (but Go will automatically take the address of a value if needed).
// - This distinction affects method sets and interface implementations.
// Reference: https://go.dev/ref/spec#Method_sets