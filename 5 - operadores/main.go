package main

import "fmt"

func main() {
	//operadores aritmeticos
	soma := 1 + 2
	fmt.Println(soma)

	subtracao := 1 - 2
	fmt.Println(subtracao)

	divisao := 10 / 2
	fmt.Println(divisao)

	multiplicacao := 10 * 2
	fmt.Println(multiplicacao)

	restoDaDivisao := 10 % 2
	fmt.Println(restoDaDivisao)

	//operadores de atribuição
	var variavel1 int = 10
	variavel1 += 5
	fmt.Println(variavel1)

	variavel2 := 10
	variavel2 -= 5
	fmt.Println(variavel2)

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//operadores logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

}
