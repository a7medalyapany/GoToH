package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", "https://example.com", nil)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed due to:", err)
		return
	}

	fmt.Println("Response status:", resp.Status)
}

/*
What’s happening?
 - If the server is slow → context cancels → HTTP request is aborted.
 - No goroutine leaks.
 - No hanging calls.
 - Your service stays fast and safe.
*/
