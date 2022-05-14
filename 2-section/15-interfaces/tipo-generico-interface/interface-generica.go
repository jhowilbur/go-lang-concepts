package main

import "fmt"

// tipo generico para func aceitar qualquer coisa // ou seja interface que aceite qualquer param
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("teste")
	generica(10)
	generica(10.5)
	generica(true)

	
}