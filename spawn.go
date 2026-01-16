package main

import (
	"fmt"
	"sync"
	"time"
)

func spawnrandom() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {

		wg.Add(1)

		go func(id int) {

			defer wg.Done()

			fmt.Print("worker %d is starting", id)
			time.Sleep(time.Second)
			fmt.Println("worker %d has finished", id)

		}(i)

	}

	fmt.Print("main : start")
	wg.Done()
	fmt.Print("main:received")

}
