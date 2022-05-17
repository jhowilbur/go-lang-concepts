package main

import (
	"fmt"
	"time"
)

/*
	Canais
Naureza dos canais, ele possue duas operacoes, envio e receber dados.
Canal tem a ideia do wait group (waitGroup.wait()) ou seja ao fazer:
	message := <-channel
e esperado que o:
	channel := make(chan string)
termine sua execucao.
*/
func main() {
	// lembrando que make define tipo, quantidade daquele tipo
	// nesse caso o make vai permitir criar canal somente do tipo X
	channel := make(chan string)

	go print("text 1", channel) // para funcionar canal eh preciso ser uma go Rutine (comecar com palavra GO)
	// executa func acima dado o go routine e vai embora

	<-channel // sintaxe de utilizar o dado enviado ao canal da func print() abaixo
	// acima nada acontece, mas pode-se resolver isso colocando dentro de uma variavel
	message := <-channel
	fmt.Println(message)
}

func print(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		// fmt.Println(text) // agora vamos enviar esse dado atravez do canal
		channel <- text // sintaxe de enviar o dado via canal
		// quando o canal receber o valor o resto do for morre, nao existe i = 1, i = 2...

		time.Sleep(time.Second)
		fmt.Printf("dentro da func print for - %d \n", i) // so executa o for uma vez pq ja pegou valor no canal
	}
	fmt.Printf("dentro da func print for - fora for \n") // ele nunca print essa parte
}
