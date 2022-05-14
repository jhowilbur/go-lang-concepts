package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

// func com ponteiro
func inverterSinalPonteiro(numero *int) {
	fmt.Println(*numero) // * é o operador de ponteiro que indica que o valor do ponteiro e nao o valor do endereço de memória
	fmt.Println(&numero) // & é o operador de ponteiro que indica o valor do endereço de memória
	*numero = *numero * -1 
	fmt.Println(*numero)
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	
	fmt.Println("\n--------------------------\n")

	inverterSinalPonteiro(&numero)
}