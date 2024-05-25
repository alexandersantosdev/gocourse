package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1 //passagem por cópia

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 *int = &variavel2    //passagem por referência
	fmt.Println(variavel2, *variavel3) //desreferenciando a variável para pegar o valor e não o endereço de memória

	*variavel3++
	fmt.Println(variavel2, *variavel3)
}
