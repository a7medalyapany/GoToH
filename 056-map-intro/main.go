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
}
