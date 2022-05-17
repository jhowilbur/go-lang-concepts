package main

import "fmt"

func main() {
	// para array interno precisa funcao make -> slice basicamente faz isso abaixo
	// tamanho de 5 capacidade maxima de 6
	slice1 := make([]int, 5, 6)
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // ver tamanho
	fmt.Println(cap(slice1)) // ver capacidade

	fmt.Println("\n----------------\n")

	// aumentando o tamanho da capacidade
	slice1 = slice1[0:6] // poderia ser
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // ver tamanho

	slice1 = append(slice1, 1) // ou poderia ser assim
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // ver tamanho
	fmt.Println(cap(slice1)) // ver capacidade, ele dobra capacidade maxima quando estoura a capacidade

	fmt.Println("\n----------------\n")
}