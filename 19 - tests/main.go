package main

import (
	"fmt"

	"github.com/alexandersantosdev/testsgo/address"
)

func main() {
	typeAddress := address.AddressType("Rua teste")
	fmt.Println(typeAddress)
}
