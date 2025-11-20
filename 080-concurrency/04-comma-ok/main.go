package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send
	go send(eve, odd, quit)
	//receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o chan<- int, q chan<- bool) {
	for i := range 10 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from eve channel:", v)
		case v := <-o:
			fmt.Println("from odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok:", i, ok)
				return
			} else {
				fmt.Println("from comma ok:", ok)

			}
		}
	}
}
