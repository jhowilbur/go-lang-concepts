package main

import "fmt"

type pessoa struct {
	nome string
	idade int8
}

// heranca
type estudante struct {
	pessoa pessoa // struct dentro de struct
	curso string
	faculdade string
}

func main() {
	fmt.Println("Heranca structs")

	p1 := pessoa{"João", 30}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "UFMG"}
	fmt.Println(e1)
}