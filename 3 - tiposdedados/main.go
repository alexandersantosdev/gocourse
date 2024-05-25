package main

import (
	"errors"
	"fmt"
)

func main() {

	//inteiros baseado na arquitetura do processador do computador
	var numero int = 10
	fmt.Println(numero)

	//int 8 bits
	var numero2 int8 = 10
	fmt.Println(numero2)

	//int 16 bits
	var numero3 int16 = 10
	fmt.Println(numero3)

	//int 32 bits
	var numero4 int32 = 10
	fmt.Println(numero4)

	//int 64 bits
	var numero5 int64 = 10
	fmt.Println(numero5)

	//uint baseado na arquitetura do processador do computador
	var numero6 uint = 10
	fmt.Println(numero6)

	//uint 8 bits
	var numero7 uint8 = 10
	fmt.Println(numero7)

	//uint 16 bits
	var numero8 uint16 = 10
	fmt.Println(numero8)

	//uint 32 bits
	var numero9 uint32 = 10
	fmt.Println(numero9)

	//uint 64 bits
	var numero10 uint64 = 10
	fmt.Println(numero10)

	//alias uint32
	var numero11 rune = 10
	fmt.Println(numero11)

	//alis uint8
	var numero12 byte = 10
	fmt.Println(numero12)

	//float32
	var numero13 float32 = 10.5
	fmt.Println(numero13)

	//float64
	var numero14 float64 = 10.5
	fmt.Println(numero14)

	//string
	var nome string = "Alexander"
	fmt.Println(nome)

	//representação do caracter na tabela ASCII
	char := 'A'
	fmt.Println(char)

	//boolean
	var booleano bool = true
	fmt.Println(booleano)

	//valor zero de todos os tipos de dados
	var texto string //valor zero = "" (VAZIO) para tipos string
	fmt.Println(texto)

	var num int //valor zero = 0 (ZERO) para tipos inteiros
	fmt.Println(num)

	var n2 float32 //valor zero = 0 (ZERO) para tipos float
	fmt.Println(n2)

	var booleano1 bool //valor zero = false (FALSO) para tipos boolean
	fmt.Println(booleano1)

	var erro error //valor zero = nil (VAZIO) para tipos error
	fmt.Println(erro)

	//ERROS
	var err error = errors.New("Erro interno")
	fmt.Println(err)
}
