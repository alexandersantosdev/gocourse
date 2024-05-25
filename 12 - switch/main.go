package main

import "fmt"

func main() {

	dia := diaDaSemana(3)

	fmt.Println(dia)

	dia2 := diaDaSemana2(3)

	fmt.Println(dia2)

}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Invalido"
	}
}

func diaDaSemana2(numero int) string {
	var dia string
	switch {
	case numero == 1:
		dia = "Domingo"
	case numero == 2:
		dia = "Segunda"
	case numero == 3:
		dia = "Terça"
	case numero == 4:
		dia = "Quarta"
	case numero == 5:
		dia = "Quinta"
	case numero == 6:
		dia = "Sexta"
	case numero == 7:
		dia = "Sabado"
	default:
		dia = "Invalido"
	}

	return dia
}
