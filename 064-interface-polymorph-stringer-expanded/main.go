package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("the title of the book is: ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("the count is: ", strconv.Itoa(int(c)))
}

func logInfo(v fmt.Stringer) {
	log.Println("Info from logInfo func", v.String())
}

func main() {

	b := book{title: "Go Programming"}
	c := count(42)

	logInfo(b)
	logInfo(c)

}
