package main

import "fmt"

func main() {

	//tipo explícito
	var nome string = "Alexander"
	fmt.Println(nome, "Roberto dos Santos")

	//inferência de tipo
	nome2 := "Alex"
	fmt.Println(nome2)

	//declaração multipla
	var (
		nome3     = "Alex"
		sobrenome = "Santos"
	)

	fmt.Println(nome3, sobrenome)

	nome4, nome5 := "Alex", "Santos"
	fmt.Println(nome4, nome5)

	//contantes
	const imutable = 5
	fmt.Println(imutable)

	//troca de valores entre variáveis
	numero1, numero2 := 1, 2
	fmt.Println(numero1, numero2)

	numero1, numero2 = numero2, numero1
	fmt.Println(numero1, numero2)

}
