package main

import (
	"fmt"
	"time"
)

func pingpong() {
	pingchan := make(chan string)
	pongchan := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			pingchan <- "ping"             // send ping
			fmt.Println("   ", <-pongchan) //recieve pong
			time.Sleep(500 * time.Millisecond)

		}

		close(pingchan)
	}()

	go func() {
		for msg := range pingchan {
			fmt.Println(msg)
			pongchan <- "pong"
		}

		done <- true

	}()

	<-done

} // pingpong moved into main.go to allow `go run main.go` without extra files.
