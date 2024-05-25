package main

import "fmt"

func main() {

	func(t string) {
		fmt.Println(t, "Func anonima")
	}("Alexander")

}
