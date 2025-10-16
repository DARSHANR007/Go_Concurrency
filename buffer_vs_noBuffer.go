package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUserProfile(userID int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate database query latency (1-5ms random variation)
	time.Sleep(time.Duration(1+userID%5) * time.Millisecond)

	// Send result
	ch <- fmt.Sprintf("User_%d_Profile", userID)
}

// Test WITHOUT buffer
func testUnbuffered(numUsers int) time.Duration {
	var wg sync.WaitGroup
	ch := make(chan string) // UNBUFFERED

	start := time.Now()

	wg.Add(numUsers)

	// Launch all goroutines
	for i := 0; i < numUsers; i++ {
		go fetchUserProfile(i, ch, &wg)
	}

	// Close channel when all done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive all results
	count := 0
	for range ch {
		count++
	}

	elapsed := time.Since(start)
	fmt.Printf("Unbuffered: %d users processed in %v\n", count, elapsed)
	return elapsed
}

// Test WITH buffer
func testBuffered(numUsers int, bufferSize int) time.Duration {
	var wg sync.WaitGroup
	ch := make(chan string, bufferSize) // BUFFERED

	start := time.Now()

	wg.Add(numUsers)

	// Launch all goroutines
	for i := 0; i < numUsers; i++ {
		go fetchUserProfile(i, ch, &wg)
	}

	// Close channel when all done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive all results
	count := 0
	for range ch {
		count++
	}

	elapsed := time.Since(start)
	fmt.Printf("Buffered (size %d): %d users processed in %v\n", bufferSize, count, elapsed)
	return elapsed
}
