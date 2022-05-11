package main

// import "fmt"
import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Escrever()

	// to call checkmail just insert after last slash bar
	erro := checkmail.ValidateFormat("wilbur@test.com")
	fmt.Println(erro)
}