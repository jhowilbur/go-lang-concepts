package main

import "fmt"

var numero int

// eh executada antes de qualquer outra funcao, incluindo a main
func init() {
	fmt.Println("init")
	numero = 10 // init pode inicializar variaveis globais
}

func main() {
	fmt.Println("main")
	fmt.Println(numero)
}