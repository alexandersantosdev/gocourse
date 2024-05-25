package main

func main() {
	numero := 0
	if numero >= 10 {
		println("Maior ou igualque 10")
	} else {
		println("Menor que 10")
	}

	if outroNumero := numero; outroNumero > 0 { // outronumero é uma variável de escopo do if/else
		println("Maior que 0")
	} else {
		println("Menor ou igaul que 0")
	}

	if numero > 0 {
		println("Maior que 0")
	} else if numero < 0 {
		println("Menor que 0")
	} else {
		println("Igual a 0")
	}

}
