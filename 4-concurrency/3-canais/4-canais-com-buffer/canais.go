package main

/*
	Canais com buffer

*/
func main() {
	/*
		channel := make(chan string)
		// channel sempre vai ficar recebendo, bloqueando assim que message receba ele
		// por isso geralmente usamos isso em func separadas
		channel <- "Ola channel" // receber
		message := <-channel     // e enviar dados
		// sao operacoes bloqueantes, de tal forma que somente assim no print gera deadlock
		println(message)
	*/

	// para resolver a problematica acima, colocamos um BUFFER no canal, especificando a capacidade dele
	channel := make(chan string, 2) // canal com capacidade 2
	channel <- "Ola channel 1"      // receber 1 valor de 2 disponiveis
	message := <-channel            // envia os dados
	println(message)                // printa um valor que recebeu

	// eh possivel colocarmos mais valor nesse canal ate atingir o limite de 2
	channel <- "Ola channel 2"
	message = <-channel
	println(message)

	// mas se inserir mais um valor, o deadlock nao
	// ocorre pois estou separando e zerando/limpando o limite do buffer
	channel <- "Ola channel 3"
	message = <-channel
	println(message)

	// mas se inserir mais tres valores de uma vez, o deadlock ocorre pois
	// estou forcando o limite do buffer que eh dois
	channel <- "Ola channel 1"
	channel <- "Ola channel 2"
	channel <- "Ola channel 3" // vai travar aqui e portanto nao chega a imprimir ultima msg
	message = <-channel
	println(message)
}
