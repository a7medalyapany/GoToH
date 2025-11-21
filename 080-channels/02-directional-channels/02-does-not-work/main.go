package main

import "fmt"

func main() {
	cr := make(<-chan int)

	// The following code is commented out because it will not compile.
	// go func() {
	// 	cr <- 42
	// }()

	fmt.Println("cr", <-cr)
}
