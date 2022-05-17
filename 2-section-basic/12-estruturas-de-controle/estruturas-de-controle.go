package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// durante if atribui o valor (cria variavel)
	if outroNumero := 16; outroNumero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}
	// fmt.Println(outroNumero) // nao consigo acessar outroNumero pois foi definido dentro do if

	// switch
	switch numero {
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	case 10:
		fmt.Println("Dez")
	default:
		fmt.Println("Não é 1, 2 ou 10")
	}
	
	// loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	
	nomes:= [3]string {"João", "Maria", "José"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	// caso nao queira indice pode colocar _
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// pegar range string como char
	for indice, letra := range "João" {
		fmt.Println(indice, letra)
		fmt.Printf("\n----------------\n")
		fmt.Printf("%#U\n", letra) // %#U mostra o codigo unicode
		fmt.Printf("\n----------------\n")
		fmt.Printf("%c\n", letra) // %c mostra o caracter
	}


}