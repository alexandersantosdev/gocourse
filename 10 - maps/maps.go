package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Alexander",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario["nome"], usuario["sobrenome"])
}
