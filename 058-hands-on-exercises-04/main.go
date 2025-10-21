package main

import (
	"fmt"
)

func main() {
	ahmed := []string{"Alex", "Cairo", "Alamain", "Marsa-Matrouh", "Ismailia", "Porto-Said", "Giza", "Behaira"}
	omar := []string{"Mansoura", "Zagazig", "Ismailia"}
	mona := []string{"Aswan", "Luxor", "Qena"}


	cityMap := map[string][]string {
		"Ahmed": ahmed,
		"Omar":  omar,
		"Mona":  mona,
	}

	for name, cities := range cityMap {
		fmt.Printf("%s has been in:\n", name)
		for _, city := range cities {
			fmt.Printf(" - %s\n", city)
		}
	}

	fmt.Println("-------- Adding a record ---------")
	mariam := []string{"Luxor", "Aswan", "Cairo"}
	cityMap["Mariam"] = mariam
	
	fmt.Println("\nAfter adding Mariam:")
	for name, cities := range cityMap {
		fmt.Printf("%s has been in:\n", name)
		for _, city := range cities {
			fmt.Printf(" - %s\n", city)
		}
	}


	fmt.Println("-------- Deleting a record ---------")
	delete(cityMap, "Mona")
	for name, cities := range cityMap {
		fmt.Printf("%s has been in:\n", name)
		for _, city := range cities {
			fmt.Printf(" - %s\n", city)
		}
	}
}
