package main

import "fmt"

func main() {

	i := 0

	for i < 10 {

		fmt.Println(i)
		i++
	}

	fmt.Println("-----------")

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	fmt.Println("-----------")

	nomes := []string{"alex", "santos"}
	for i, nome := range nomes {
		fmt.Println(i, nome)
	}

	fmt.Println("-----------")

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("-----------")

	for _, letra := range "alex" {
		fmt.Println(string(letra))
	}

	fmt.Println("-----------")

	z := 0

	for { // loop infinito com condição de parada -> (WHILE true)
		fmt.Println("loop")
		if z >= 100 {
			break
		}

		z++
	}

}
