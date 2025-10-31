package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

type person struct {
	firstName string
}

func (p person) writeOut(w io.Writer) error {
	_, err := io.WriteString(w, "Hello, "+p.firstName)
	return err
}

func main() {
	p := person{firstName: "Alyapany"}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}

	defer f.Close()

	s := []byte("Hello, World! ")

	_, err = f.Write(s)

	if err != nil {
		log.Fatalf("Could not write to file: %v", err)
	}

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)

	log.Println("From buffer:", b.String())
}
