package main;

import "fmt";

func main() {

	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	variavel2 := "variavel2"
	fmt.Println(variavel2)

	variavel3, variavel4 := "variavel3", "variavel4"
	fmt.Println(variavel3, variavel4)
//------------------------------------------------------------------
	const constante1 = "constante1"
	fmt.Println(constante1)

	const constante2 string = "constante2"
	fmt.Println(constante2)

	const constante3, constante4 = "constante3", "constante4"
	fmt.Println(constante3, constante4)
}