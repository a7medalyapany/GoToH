package main

import (
	"fmt"
	"sort"
)

type User struct {
	FirstName string
	Age       int
}

type ByAge []User
type ByName []User

func (ba ByAge) Len() int {
	return len(ba)
}

func (ba ByAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

func (ba ByAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Less(i, j int) bool {
	return bn[i].FirstName < bn[j].FirstName
}

func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}



func main() {

	u1 := User{
		FirstName: "Mariam",
		Age:       30,
	}

	u2 := User{
		FirstName: "Bob",
		Age:       60,
	}

	u3 := User{
		FirstName: "Menna",
		Age:       21,
	}

	u4 := User{
		FirstName: "Alyapany",
		Age:       23,
	}

	users := []User{u1, u2, u3, u4}

	sort.Sort(ByAge(users))
	fmt.Printf("After Sorting people by age: %+v\n", users)

	sort.Sort(ByName(users))
	fmt.Printf("\nAfter Sorting people by name: %+v", users)

}