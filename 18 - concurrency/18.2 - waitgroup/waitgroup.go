package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup) // create a new WaitGroup
	wg.Add(2)                 // added the quantity of go routines to wait

	/*
		for each goroutine, the WaitGroup will be decremented by 1 when call wg.Done().
	*/

	go func() {
		writeLine("go routine - 1")
		wg.Done()
	}()

	go func() {
		writeLine("go routine - 2")
		wg.Done() // decrement the wait group
	}()

	wg.Wait() // Wait for all go routines to finish
}

func writeLine(text string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(text)
	}
}
