package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("ths is the start ")

	now := time.Now()

	userID := 10

	getUserId(userID)
	getUserLikes(userID)
	getUserPosts(userID)

	fmt.Println(time.Since(now))

}

func getUserId(id int) {

	time.Sleep(120 * time.Millisecond)

}

func getUserLikes(id int) {
	time.Sleep(80 * time.Millisecond)
}

func getUserPosts(id int) {
	time.Sleep(50 * time.Millisecond)
}
