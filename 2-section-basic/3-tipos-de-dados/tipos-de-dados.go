package main

import "fmt"

func main() {
	// Tipos de int
	// int8, int16, int32, int64

	var numeroint16 int16 = 100
	fmt.Println(numeroint16)

	// se não for declarado, o compilador irá inferir o tipo de acordo com architetura do sistema (int32 ou int64)
	var numeroint int = 1000000000000000
	fmt.Println(numeroint)
	numeroint2 := 1000000000000000
	fmt.Println(numeroint2)

	// Tipos de float
	// float32, float64
	var numerofloat32 float32 = 1.5
	fmt.Println(numerofloat32)
	numerofloat64 := 1.5
	fmt.Println(numerofloat64)

	// char retornar tabela ascii
	var char rune = 'a'
	fmt.Println(char)
	char2 := 'a'
	fmt.Println(char2)

	// tipo erro
	var erro error = nil
	fmt.Println(erro)
}