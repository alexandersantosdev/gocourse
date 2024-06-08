package main

import (
	"fmt"
	"time"
)

func main() {

	channel := write("Ola mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

// returns an channel
func write(text string) <-chan string {

	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("%s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
