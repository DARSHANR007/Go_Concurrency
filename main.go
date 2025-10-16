package main

import (
	"fmt"
)

func main() {
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
