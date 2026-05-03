package main

import (
	"fmt"
	"sync"
)

func launch5() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			fmt.Println(i)
		}()

		wg.Wait()

		fmt.Println("received 5")

	}
}
