package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)
	go writeLine("Ola mundo ", channel)

	for line := range channel {
		fmt.Println(line)
	}
}

func writeLine(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- fmt.Sprintf("%s %d", text, i+1)
		time.Sleep(time.Second)
	}
	close(channel)
}
