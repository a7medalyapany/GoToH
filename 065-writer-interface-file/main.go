package main

import (
	"bytes"
	"fmt"
)

func main() {

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
