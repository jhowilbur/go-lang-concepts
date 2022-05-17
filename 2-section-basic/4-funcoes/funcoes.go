package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

									// defenido dois tipos de retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	// retornando dois valores
	return soma, subtracao
}

func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	// func inside func
	var f = func(txt string) string {
		return txt + " world"
	}
	fmt.Println(f("Hello"))

	// em GO funcoes podem ter retorno de mais de um valor/ podem ter mais de um retorno
	// retornando um valor
	// resultado := calculosMatematicos(1, 2)
	// fmt.Println(resultado)
	//causa erro o acima
	// retornando dois valores
	soma2, subtracao := calculosMatematicos(1, 2)
	fmt.Println(soma2, subtracao)
	// e caso vc nao queira usar o retorno, pode usar _
	_, subtracao2 := calculosMatematicos(1, 2)
	fmt.Println(subtracao2)

	
}