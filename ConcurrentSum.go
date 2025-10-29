package main

import "sync"

func splitSum(array []int) *int {

	var wg sync.WaitGroup

	n := len(array) / 2

	channel := make(chan int, 2)

	firstHalf := array[:n]
	secondHalf := array[n:]

	sumfunction := func(part []int) {

		defer wg.Done()

		sum := 0

		for _, val := range part {
			sum += val
		}

		channel <- sum
	}

	wg.Add(2)

	go sumfunction(firstHalf)
	go sumfunction(secondHalf)

	wg.Wait()
	close(channel)

	total := 0

	for i := range channel {
		total += i
	}

	return &total

}
