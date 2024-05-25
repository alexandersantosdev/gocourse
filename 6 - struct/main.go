package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	pessoa := usuario{
		nome:  "Alexander",
		idade: 20,
		endereco: endereco{
			logradouro: "Rua dos bobos",
			numero:     0,
		},
	}

	fmt.Println(pessoa, pessoa.nome, pessoa.idade)

}
