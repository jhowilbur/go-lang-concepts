package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)

	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	fmt.Println(texto, numeros)
}

func main() {
	total := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	escrever("Valores: ", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}