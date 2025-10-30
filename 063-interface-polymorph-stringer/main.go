package main

import (
	"fmt"
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

func main() {

	b := book{title: "Go Programming"}
	c := count(42)

	fmt.Println(b)
	fmt.Println(c)

}
