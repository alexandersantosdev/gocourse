package main

import "fmt"

func main() {

	a := 10
	b := 20

	soma := soma(a, b)

	fmt.Println(soma)

	subtracao := subtracao(a, b)

	fmt.Println(subtracao)

}

func soma(a int, b int) (soma int) {
	soma = a + b
	return
}

func subtracao(a int, b int) (subtracao int) {
	subtracao = a - b
	return
}
