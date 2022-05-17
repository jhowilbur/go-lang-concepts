package main

import (
	"fmt" 
	"math"
)

type retangulo struct {
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}

// interface vai permitir passar tanto forma de circulo quanto de retangulo
type forma interface {
	// so pode passar assinaturas de metodos
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("area %0.2f \n", f.area())
}

// Metodo area
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{altura: 10, largura: 20}
	escreverArea(r)

	c := circulo{raio: 10}
	escreverArea(c)
}