package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Mom gives you a context with timeout of 2 seconds:
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go bakeCookies(ctx, "Kid 1")
	go bakeCookies(ctx, "Kid 2")

	time.Sleep(3 * time.Second)
}

func bakeCookies(ctx context.Context, kid string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(kid, "STOPPED! Reason:", ctx.Err())
			return
		default:
			fmt.Println(kid, "🍪 baking...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
What happens?
 - Kids bake cookies.
 - After 2 seconds, context timeout fires.
 - Both kids stop instantly.
*/
