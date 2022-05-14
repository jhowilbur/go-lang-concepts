package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")
	
	var array1[5] int8
	fmt.Println(array1)

	array2 := [5] int8 {1, 2, 3, 4, 5}
	fmt.Println(array2)

	// os tres ... nao transforma em array dinamico, so pega valor entre {} pra saber tamanho final do array
	array3 := [...] int8 {10, 20, 30, 40, 50}
	fmt.Println(array3)

	// slice eh um array dinamico, nao precisa definir tamanho, apenas o tipo
	slice1 := [] int16 {100, 200, 300, 400, 500}
	fmt.Println(slice1)
	// para adicionar como nao temos indice no slice, usamos append, ele cria um novo slice com o novo valor
	slice1 = append(slice1, 600)
	fmt.Println(slice1)

	// pegar dados do array, pegando do indice 0 ate o indice 2 (parando no 3)
	slice2 := array2[0:3]
	fmt.Println(slice2)

	slice3 := array2[:] // pegando todos os dados do array
	fmt.Println(slice3)
}