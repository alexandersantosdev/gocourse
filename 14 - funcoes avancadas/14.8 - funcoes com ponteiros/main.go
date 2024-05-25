package main

import "fmt"

func main() {
	num := 20
	nInv := inverter(num)

	fmt.Println(num)
	fmt.Println(nInv)

	num2 := 20
	inverter2(&num2)
	fmt.Println(num2)

}

func inverter(n int) int {
	return n * -1
}

func inverter2(n *int) {
	*n = *n * -1
}
