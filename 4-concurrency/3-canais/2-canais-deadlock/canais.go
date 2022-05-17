package main

import (
	"fmt"
	"time"
)

/*
	Canais com deadlock
programa vai esperar receber um valor no canal mas isso nunca acontece
entao estoura o deadlock
*/
func main() {
	channel := make(chan string)
	go print("text", channel)

	for { // loop infinito, ao final do for do print (5 vezes) estoura deadlock
		message, aberto := <-channel // canal permite receber mais um retorno, de saber se ta aberto
		fmt.Println(message)
		if !aberto {
			fmt.Printf("canal aberto ? %t \n", aberto)
			break
		}
	}

	println("fim do programa")
}

func print(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text // sintaxe de enviar o dado via canal
		time.Sleep(time.Second)
	}

	// no caso essa linha abaixo resolve problea do deadlock, fechando o canal
	close(channel)
}
