package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	fmt.Printf(" this is the start , it starts at %v ", now)

	fmt.Println()

	var wg sync.WaitGroup

	msgchan := make(chan string, 15)

	wg.Add(10)

	i := 0

	for i = 0; i < 10; i++ {

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
