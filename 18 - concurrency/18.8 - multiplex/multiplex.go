package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	channel := multiplex(write("Ola mundo"), write("Ola mundo 2"))

	for i := 0; i < 10; i++ {
		message := <-channel
		fmt.Println(message)
	}

}

func multiplex(c1, c2 <-chan string) <-chan string {

	channel := make(chan string)

	go func() {
		for {
			select {
			case m := <-c1:
				channel <- m
			case m := <-c2:
				channel <- m
			}
		}
	}()
	return channel
}

func write(text string) <-chan string {

	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("%s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
