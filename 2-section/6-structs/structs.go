package main

import 	"fmt"

// struct eh um tipo personalizado
type usuario struct {
	nome string
	idade int8
}

// struct usuario com endereco
type usuarioEndereco struct {
	nome string
	idade int8
	endereco endereco // struct dentro de struct
}

type endereco struct {
	rua string
	numero int8
}

func main() {
	fmt.Println("Arquivo structs")

	var usuario1 usuario
	usuario1.nome = "João"
	usuario1.idade = 30
	fmt.Println("Nome:", usuario1.nome)
	fmt.Println("Idade:", usuario1.idade)

	usuario2 := usuario{"Maria", 25}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Pedro", idade: 35}
	fmt.Println(usuario3)

	// struct dentro de struct
	usuarioStruct := usuarioEndereco{nome: "João", idade: 30, endereco: endereco{rua: "Rua ABC", numero: 123}}
	fmt.Println(usuarioStruct)
}