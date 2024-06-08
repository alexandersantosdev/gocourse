package main

import (
	"fmt"
	"time"
)

func main() {

	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"
		}
	}()

	for {
		select {
		case message := <-channel1:
			fmt.Println(message)
		case message := <-channel2:
			fmt.Println(message)
		}
	}
}
