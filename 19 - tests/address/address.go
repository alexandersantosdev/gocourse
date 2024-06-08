package address

import "strings"

var TypeAddress = []string{"Rua", "Avenida", "Estrada", "Travessa", "Praça", "Alameda", "Viela", "Rodovia", "Trilha", "Parque", "Jardim"}

func AddressType(address string) string {

	firstWord := strings.Split(address, " ")[0]
	isValid := false

	for _, t := range TypeAddress {
		if strings.Contains(strings.ToLower(firstWord), strings.ToLower(t)) {
			isValid = true
			break
		}
	}

	if isValid {
		return firstWord
	}

	return "Tipo Inválido"
}
