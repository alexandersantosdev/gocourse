package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	pessoa := pessoa{"Alexander", "Santos", 20}
	pessoa.escrever()
	pessoa.fazerAniversario()
	pessoa.escrever()
}

func (p pessoa) escrever() {
	fmt.Println(p.nome, p.sobrenome, p.idade)
}

func (p *pessoa) fazerAniversario() {
	p.idade++
}
