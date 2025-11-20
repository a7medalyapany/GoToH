package main

import (
	"fmt"
	"time"
)

// Each friend (goroutine) makes cookies
func friendMakeCookies(friendID int, bowl chan<- string) {
	// Pretend each friend spends time baking
	time.Sleep(time.Duration(friendID) * 300 * time.Millisecond)

	cookie := fmt.Sprintf("🍪 Cookie from friend %d", friendID)

	// Put cookie in the bowl
	bowl <- cookie
}

func main() {

	// The big bowl where everyone puts their cookies
	bowl := make(chan string)

	// FAN-OUT: Send 5 friends to go bake cookies
	for friend := 1; friend <= 5; friend++ {
		go friendMakeCookies(friend, bowl)
	}

	// FAN-IN: Collect all cookies back into the bowl
	for i := 1; i <= 5; i++ {
		cookie := <-bowl
		fmt.Println("Chef received:", cookie)
	}

	fmt.Println("All cookies collected! 🥣✨")
}
