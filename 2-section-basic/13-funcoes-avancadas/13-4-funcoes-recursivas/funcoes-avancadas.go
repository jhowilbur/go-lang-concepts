package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Recursivas")

	// sequencia fibonacci, proximo numero sempre soma dos outros dois numeros anteriores
	fmt.Println(fibonacci(10))

}