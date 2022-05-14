package main

import "fmt"

func main() {
	// ponteiro eh referencia de memoria, ou seja endereco de memoria sera mesmo para varaiveis de ponteiro
	var variavel1 int
	var ponteiro1 *int // ponteiro

	variavel1 = 10
	fmt.Println(&variavel1) // print endereco de memoria da variavel1

	ponteiro1 = &variavel1 // ponteiro aponta para variavel1

	fmt.Println(ponteiro1)  // printa o endereco de memoria da variavel1
	fmt.Println(*ponteiro1) // print o valor, * ponteiro1 pega o valor da variavel que o ponteiro aponta

	*ponteiro1 = 20 // altera o valor da variavel que o ponteiro aponta
	fmt.Println(variavel1)
}