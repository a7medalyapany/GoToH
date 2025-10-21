package main

import "fmt"

func main() {
	am := map[string]int {
		"Ahmed": 23,
		"Ali":  17,
		"Omar": 45,
	}

	for k, v := range am {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	fmt.Printf("MAP : %#v\n", am)

	pm := make(map[string]int)
	pm["Apple"] = 100
	pm["Banana"] = 200
	pm["Grapes"] = 300

	fmt.Printf("MAP2 : %#v - len: %v\n", pm, len(pm))

	fmt.Println("-----------------")
	// Deleting an element
	delete(am, "omar") // wont delete as key is case-sensitive also wont panic
	fmt.Printf("MAP : %#v - len: %v\n", am, len(am))
	delete(am, "Omar")
	fmt.Printf("MAP : %#v - len: %v\n", am, len(am))

	fmt.Printf("It will return the ZERO value - %v\n", am["Menna"])


	fmt.Println("-----------------")
	v, ok := am["Mariam"]
	if ok {
		fmt.Printf("Found: %v\n", v)
	} else {
		fmt.Printf("Not Found\n")
	}

	if v, ok := am["Ahmed"]; !ok {
		fmt.Printf("Not Found\n")
	} else {
		fmt.Printf("Found: %v\n", v)
	}
}
