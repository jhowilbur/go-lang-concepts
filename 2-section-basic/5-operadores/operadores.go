package main

import "fmt"

func main() {
	// operadores aritméticos
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 - 1 =", 1-1)
	fmt.Println("1 * 1 =", 1*1)
	fmt.Println("1 / 1 =", 1/1)
	fmt.Println("1 % 1 =", 1%1)

	// operadores de atribuição
	var x int = 1
	x += 1
	fmt.Println("x =", x)

	// operadores relacionais
	fmt.Println("1 == 1 =", 1 == 1)
	fmt.Println("1 != 1 =", 1 != 1)
	fmt.Println("1 < 1 =", 1 < 1)
	fmt.Println("1 > 1 =", 1 > 1)
	fmt.Println("1 <= 1 =", 1 <= 1)
	fmt.Println("1 >= 1 =", 1 >= 1)

	// operadores lógicos
	fmt.Println("true && true =", true && true)
	fmt.Println("true || true =", true || true)
	fmt.Println("!true =", !true)
}