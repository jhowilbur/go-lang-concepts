package main

import "fmt"

func tentativaRecuperarExecucao() {
	fmt.Println("Tentativa de recuperar execucao")
}

// para usar recovery, usar da func abaixo
// no caso abaixo foi pego pelo "defer" para recuperar execucao
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando execucao")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer tentativaRecuperarExecucao()
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media >= 7 {
		return true
	} else if media < 6 {
		return false
	}

	//panic caso media seja 6
	panic("Media invalida") // quando chama essa func ele para de fazer oque estava e entrara em panico (Bem parecido com Throw do java pra explodir exceptions)
}

func main() {
	// o erro avisa mas segue o programa, o panic mata o programa
	fmt.Println(alunoAprovado(6, 7))
	fmt.Println("Pos execucao")
}