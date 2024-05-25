package main

import "fmt"

func main() {
	defer recuperar()
	n := isNumberEqualsZero(0)

	fmt.Println(n)

}

func isNumberEqualsZero(n int) bool {
	if n < 0 {
		return false
	} else if n > 0 {
		return false
	}
	panic("Zero")
}

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Erro:", r)
	}
}
