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

func test() {
	numUsers := 1000

	fmt.Println("=== Performance Test: 1000 Concurrent User Fetches ===\n")

	// Test 1: Unbuffered
	fmt.Println("Test 1: UNBUFFERED CHANNEL")
	unbufferedTime := testUnbuffered(numUsers)

	fmt.Println()

	// Test 2: Small buffer
	fmt.Println("Test 2: SMALL BUFFER (size 10)")
	smallBufferTime := testBuffered(numUsers, 10)

	fmt.Println()

	// Test 3: Medium buffer
	fmt.Println("Test 3: MEDIUM BUFFER (size 100)")
	mediumBufferTime := testBuffered(numUsers, 100)

	fmt.Println()

	// Test 4: Large buffer
	fmt.Println("Test 4: LARGE BUFFER (size 1000)")
	largeBufferTime := testBuffered(numUsers, 1000)

	fmt.Println()
	fmt.Println("=== Performance Comparison ===")
	fmt.Printf("Unbuffered:        %v (baseline)\n", unbufferedTime)
	fmt.Printf("Small Buffer:      %v (%.2fx faster)\n", smallBufferTime, float64(unbufferedTime)/float64(smallBufferTime))
	fmt.Printf("Medium Buffer:     %v (%.2fx faster)\n", mediumBufferTime, float64(unbufferedTime)/float64(mediumBufferTime))
	fmt.Printf("Large Buffer:      %v (%.2fx faster)\n", largeBufferTime, float64(unbufferedTime)/float64(largeBufferTime))

	// Calculate improvement
	improvement := float64(unbufferedTime-largeBufferTime) / float64(unbufferedTime) * 100
	fmt.Printf("\nPerformance improvement: %.1f%% faster with buffer\n", improvement)
}
