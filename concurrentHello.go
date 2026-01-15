package main

import (
	"fmt"
	"sync"
	"time"
)

func concurrentWorld(msgChan chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	msgChan <- "This is hello world from concurrency "

}

func testConcurrentWorld() {

	now := time.Now()

	fmt.Printf(" this is the start , it starts at %v ", now)

	fmt.Println()

	var wg sync.WaitGroup

	msgchan := make(chan string, 1100)

	wg.Add(1000)

	i := 0

	for i = 0; i < 1000; i++ {

		go concurrentWorld(msgchan, &wg)

	}

	go func() {
		wg.Wait()
		close(msgchan)
	}()

	for msg := range msgchan {
		fmt.Println(msg)

	}

	fmt.Printf(" ending time %v", time.Since(now))
}
