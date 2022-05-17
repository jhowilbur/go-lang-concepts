package main

import "time"

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"
		}
	}()

	/*
		for {
			messageCanal1 := <-channel1 // chego aqui e recebe a mensagem do canal um espero meio segundo
			println(messageCanal1)

			messageCanal2 := <-channel2 // so que o problema acontece aqui, pois espero 2 segundos antes
			// do for voltar e executar canal um com meio segundo, sendo que
			// poderia executar o channel 1 sem esperar 2 segundos
			println(messageCanal2)
		}
	*/

	// para resolver problematica acima, utilizamos o select (parecido com switch case)
	for {
		select {
		case messageCanal1 := <-channel1: // caso canal 1 esteja pronto pra receber, printa na tela
			println(messageCanal1)
		case messageCanal2 := <-channel2: // caso canal 2 esteja pronto pra receber, printa na tela
			println(messageCanal2)

		}
	}

}
