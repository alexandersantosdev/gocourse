package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa
	matricula uint16
	curso     string
}

func main() {

	p1 := pessoa{"Alexander", 20}
	fmt.Println(p1)

	e1 := estudante{p1, 123, "Engenharia"}
	fmt.Println(e1)

	e2 := estudante{pessoa{"Alex", 20}, 123, "Engenharia"}
	fmt.Println(e2)

}
