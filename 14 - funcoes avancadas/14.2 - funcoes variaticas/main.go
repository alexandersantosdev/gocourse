package main

import "fmt"

func main() {

	soma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(soma)
}

func soma(numeros ...int) (soma int) {

	for _, num := range numeros {
		soma += num
	}

	return
}
