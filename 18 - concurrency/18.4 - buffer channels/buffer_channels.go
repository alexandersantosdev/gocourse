package main

import "fmt"

func main() {

	channel := make(chan string, 2) //removendo o buffer do canal (valor 2) daria deadlock

	channel <- "Ola mundo"
	message := <-channel
	fmt.Println(message)

	channel <- "Ola mundo 2"
	message2 := <-channel
	fmt.Println(message2)

}
