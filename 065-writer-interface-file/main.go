package main

import (
	"bytes"
	"fmt"
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

	fmt.Println("----------------Buffer-----------------")
	// Create a buffer and write to it
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("World")
	fmt.Println(b.String())

	b.Reset()
	b.WriteString("Alyapany")
	fmt.Println(b.String())
	b.Write([]byte(" Was Here"))
	fmt.Println(b.String())
}
