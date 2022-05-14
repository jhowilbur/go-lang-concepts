package main

import "fmt"

type usuario struct {
	nome string
	idade int
}

// basicamente metodo eh uma func com strutura de outro tipo
func (user usuario) salvar() {
	fmt.Printf("salvando o user %s \n", user.nome)
}

// pode-se colocar retornos tambem
func (user usuario) maiorIdade() bool {
	return user.idade >= 18
}

func (user *usuario) fazerAniversario() {
	user.idade++ // o ponteiro eh usado para alterar o valor da variavel
}

func main () {
	fmt.Println("metodo")

	user := usuario{"Joao", 20}
	fmt.Println(user)
	user.salvar()
	fmt.Println(user.maiorIdade())

	user.fazerAniversario()
	fmt.Println(user.idade)
}