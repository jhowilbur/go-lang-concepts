package main

import (
	"cli-interface/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando app")

	aplication := app.Gerar()
	erro := aplication.Run(os.Args)
	if erro != nil {
		log.Fatal(erro) // fatal faz parar execucao do programa
	}

	fmt.Println("Busca concluida")
}
