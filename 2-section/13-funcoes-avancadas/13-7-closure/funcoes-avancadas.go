package main

import "fmt"

func closure() func() {
	text := "dentro da closure"

	funcao := func() {
		fmt.Println(text) // busca texto dentro da closure
	}

	return funcao
}

func main() {
	text := "fora da closure"
	fmt.Println(text)

	funcaoNova := closure()
	funcaoNova()
}