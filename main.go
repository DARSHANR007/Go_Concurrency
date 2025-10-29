package main

import "fmt"

func main() {
	// start := time.Now() // record start time

	// for i := 1; i <= 1000; i++ {
	// 	fmt.Printf("Hello, World %d\n", i)
	// }

	// elapsed := time.Since(start) // calculate time taken
	// fmt.Sprintf("Time taken: %v\n", elapsed)

	// testConcurrentWorld()

	array := []int{1, 2, 3, 4, 5, 6}

	result := splitSum(array)

	fmt.Println(result)

	fmt.Println(*result)
}
