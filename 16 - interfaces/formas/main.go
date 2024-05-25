package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("Area da forma: %.2f\n", f.area())
}

func main() {

	r := retangulo{
		largura: 10,
		altura:  15,
	}

	c := circulo{
		raio: 10,
	}

	escreverArea(r)
	escreverArea(c)

}
