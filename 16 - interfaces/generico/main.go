package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	generic(10)
	generic("Alexander")
	generic(true)

}
