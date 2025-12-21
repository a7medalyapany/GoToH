package main

import (
	"fmt"
	"sync"
	"time"
)

// processImage simulates processing an image (resize, compress, etc.)
func processImage(id int, image string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	
	// Simulate some work
	time.Sleep(time.Millisecond * time.Duration(100+id*50))
	
	result := fmt.Sprintf("Worker %d processed: %s", id, image)
	results <- result
}

func main() {
	images := []string{
		"photo1.jpg", "photo2.jpg", "photo3.jpg", 
		"photo4.jpg", "photo5.jpg", "photo6.jpg",
	}
	
	var wg sync.WaitGroup
	results := make(chan string, len(images))
	
	// Fan-out: Distribute work across multiple goroutines
	fmt.Println("Starting fan-out processing...")
	start := time.Now()
	
	for i, img := range images {
		wg.Add(1)
		go processImage(i+1, img, &wg, results)
	}
	
	// Wait for all workers to finish
	wg.Wait()
	close(results)
	
	// Collect results
	fmt.Println("\nResults:")
	for result := range results {
		fmt.Println(result)
	}
	
	fmt.Printf("\nCompleted in: %v\n", time.Since(start))
}