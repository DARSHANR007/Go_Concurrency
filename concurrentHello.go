package main

import (
	"fmt"
	"sync"
)

func concurrentWorld(msgChan chan string, wg *sync.WaitGroup) {

	defer wg.Done()

	msgChan <- fmt.Sprintf("This is hello world from concurrency ")

}
