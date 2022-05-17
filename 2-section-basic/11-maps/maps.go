package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// usuario := make(map[string]int)
	usuario := map[string] string{ // as chaves e os valores das chaves tem q ser do mesmo tipo no caso string e string
		"nome": "João",
		"idade": "23",
	}
	fmt.Println(usuario)

	// para pegar um dos valores do map
	fmt.Println(usuario["nome"])

	// para modificar o valor de um map
	usuario["nome"] = "Maria"
	fmt.Println(usuario)

	// para deletar um valor do map
	delete(usuario, "nome")
	fmt.Println(usuario)

	// para verificar se um valor existe no map
	_, existe := usuario["nome"]
	fmt.Println(existe)

	// eh possivel ter map dentro de map
	usuario2 := map[string] map[string] string{
		"nome": {
			"nome": "João",
			"idade": "23",
		}, // ou mais um map dentro de outro map semelhante a nome
	}
	fmt.Println(usuario2)
}