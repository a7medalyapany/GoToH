package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := strings.NewReader("Hello Nati")

	io.Copy(f, r)

}
