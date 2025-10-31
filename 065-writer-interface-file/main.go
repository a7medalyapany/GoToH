package main

import (
	"log"
	"os"
)

// type person struct {
// 	firstName string
// }

// func (p person) writeOut(w io.Writer) error {
// 	_, err := io.WriteString(w, "Hello, "+p.firstName)
// 	return err
// }

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}

	defer f.Close()

	s := []byte("Hello, World!")

	_, err = f.Write(s)

	if err != nil {
		log.Fatalf("Could not write to file: %v", err)
	}
}
