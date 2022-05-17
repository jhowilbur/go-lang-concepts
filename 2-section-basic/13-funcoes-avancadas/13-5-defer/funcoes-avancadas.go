package main

import "fmt"

func funcao1() {
	fmt.Println("Funcao 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func main() {
	defer funcao1() // defer funcao1() executa a funcao1() depois que a funcao2() termina
	// ou seja uma func defer adia a execucao ate o ultimo momento possivel
	// DEFER = ADIAR
	funcao2()
}