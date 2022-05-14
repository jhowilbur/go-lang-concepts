package main

import "fmt"

func main() {

	func () {
		fmt.Println("Olá mundo!")
	}() // com () executa a função

	// ou entao eh possivel passar parametros assim
	func (nome string) {
		fmt.Println("Olá", nome)
	}("João")

	// posso guardar retornos da variavel
	retorno := func (nome string) string {
		return fmt.Sprintf("Olá mundo - %s", nome)
	}("Joao")
	fmt.Println(retorno)
	// ou entao posso retornar uma função assim com o parametro sendo enviado pela variavel
	retorno2 := func (nome string) string {
		return fmt.Sprintf("Olá mundo - %s", nome)
	}
	fmt.Println(retorno2("Joao"))

	
	
}