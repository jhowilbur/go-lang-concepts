package main

import (
	"time"
)

/*
	Canais sem for
sintaxe sem for e sem seta <-
*/
func main() {
	channel := make(chan string)
	go print("text", channel)

	// para cada mensagem recebida no canal enquanto ele estiver aberto
	// e feito o print na tela
	for message := range channel {
		println(message)
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
