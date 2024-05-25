package main

import "fmt"

func main() {
	texto := "dentro da main"
	fmt.Println(texto)
	funcao := closure()
	funcao()
}

func closure() func() {
	texto := "dentro da closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}
