package main

import "fmt"

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(nome string) {
		fmt.Println("Func anonima", nome)
	}

	f("Alexander")

	var f2 = func(n int) int {
		if n%2 == 0 {
			return n / 2
		}

		return n * 2
	}

	value := f2(4)
	fmt.Println(value)
	value2 := f2(5)

	fmt.Println(value2)

	soma, subtracao := calculos(20, 10)

	fmt.Println(soma, subtracao)

}

func somar(a int8, b int8) int8 {
	return a + b
}

//retorno de mais de um valor
//ambos os parametros do mesmo tipo poderm ser tipados somente no ultimo
func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
