package main

func main() {
	cr := make(chan<- int)

	go func() {
		cr <- 42
	}()

	// The following line is commented out because it will cause a deadlock at runtime.
	// fmt.Println("cr", <-cr)
}
