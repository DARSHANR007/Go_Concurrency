package main

import (
	"fmt"
	"time"
)

func pingpong() {
	pingchan := make(chan string)
	pongchan := make(chan string)

	go func() {
		for {
			pingchan <- "ping"
			fmt.Println("   ", <-pongchan)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			fmt.Println("   ", <-pingchan)
			pongchan <- "pong"
		}
	}()

	select {}
} // pingpong moved into main.go to allow `go run main.go` without extra files.
