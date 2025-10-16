package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("ths is the start ")

	channel1 := make(chan string)

	now := time.Now()

	userID := 10

	go getUserId(userID, channel1, &wg)
	go getUserLikes(userID, channel1, &wg)
	go getUserPosts(userID, channel1, &wg)

	go func() {
		wg.Wait()
		close(channel1)
	}()

	for curr := range channel1 {
		fmt.Println(curr)

	}

	fmt.Println(time.Since(now))

}

func getUserId(id int, response chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(120 * time.Millisecond)

	response <- "got the userID"

}

func getUserLikes(id int, response chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(80 * time.Millisecond)
	response <- "got the userLikes"

}

func getUserPosts(id int, response chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	time.Sleep(50 * time.Millisecond)
	response <- "got the userPosts"
}
