package main

import (
	"fmt"
	"time"
)

func pongping() {
	table := make(chan string)
	done := make(chan bool)

	// Player goroutine
	go func() {
		count := 0
		for {
			msg := <-table

			if msg == "ping" {
				fmt.Println("ping")
				time.Sleep(500 * time.Millisecond)
				table <- "pong"
			} else if msg == "pong" {
				fmt.Println("pong")
				count++
				if count >= 5 {
					done <- true
					return
				}
				time.Sleep(500 * time.Millisecond)
				table <- "ping"
			}
		}
	}()

	// Start the game
	table <- "ping"

	// Wait for done signal
	<-done
}
