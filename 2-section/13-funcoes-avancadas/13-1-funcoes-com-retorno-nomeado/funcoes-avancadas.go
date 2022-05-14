package main

import "fmt"

func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2 // sem := pq ja definido no inicio
	return // retorno ja sabe os valores
}

func main() {
	soma, subtracao := calculos(10, 5)
	fmt.Println(soma, subtracao)
}