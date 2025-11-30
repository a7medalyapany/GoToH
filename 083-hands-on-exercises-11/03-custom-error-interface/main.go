package main

import "fmt"

type customError struct {
	info string
}

func (e *customError) Error() string {
	return fmt.Sprintf("Custom Error Interface: %s", e.info)
}

func main() {
	err := &customError{info: "Something went wrong!"}
	foo(err)
}

func foo(e error) {
	// Type assertion to access customError fields -> e.(*customError).info
	fmt.Printf("foo ran - %v\n%v\n", e, e.(*customError).info)
}