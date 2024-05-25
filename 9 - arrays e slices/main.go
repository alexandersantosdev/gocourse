package main

import "fmt"

func main() {

	//arrays devem ter o tamanho definido na sua declaração
	var array [5]int
	fmt.Println(array)

	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50
	fmt.Println(array)

	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	//slices

	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	slice = append(slice, 60)
	fmt.Println(slice)

	slice2 := array[1:3] // pegando os elementos do indice 1 (inclusivo) ao 3 (exclusivo)
	fmt.Println(slice2)

	fmt.Println("--------------------")

	slice3 := make([]int, 5, 25)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho do slice
	fmt.Println(cap(slice3)) //capacidade maxima
}
