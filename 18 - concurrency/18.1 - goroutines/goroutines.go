package main

import (
	"fmt"
	"time"
)

func main() {

	go writeLine("Ola")
	writeLine("Mundo")
}

func writeLine(text string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(text)
	}
}
