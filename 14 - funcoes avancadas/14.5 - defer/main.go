package main

import "fmt"

func main() {
	//fluxo padrão de execução
	f1()
	f2()

	defer f1() //adia a execução de f1 para depois de todas as exucuções do código e antes de sair da função
	f2()

	if aprovado := calcAprovado(3.5, 6.5); aprovado {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}

func f1() {
	fmt.Println("Executando f1")
}

func f2() {
	fmt.Println("Executando f2")
}

func calcAprovado(n1, n2 float64) bool {
	defer fmt.Println("Media calculada. Resultado será retornado") //adia o print para antes de dar o return da função
	media := (n1 + n2) / 2
	return media > 6
}
