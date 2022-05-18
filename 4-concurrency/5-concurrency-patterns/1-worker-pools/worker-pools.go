package main

import (
	"fmt"
	"time"
)

/*
	Worker pools
imagine uma fila de tarefas a ser executada
com esse pattern vc possui varios workers pegando dessa fila e executando as tarefas
*/

func main() {
	timeStart := time.Now().Second()
	// canal contendo sequencia de numeros a executar
	tasks := make(chan int, 45) //canal com 45 de buffer

	// canal para armazenar os resultados
	results := make(chan int, 45)

	go worker(tasks, results) // antes mesmo de popular canal de tasks
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		println(result)
	}

	fmt.Printf("final : %d", time.Now().Second()-timeStart)
}

// posso usar das setas para definir se recebe ou armazena dados no canal
// worker eh quem chama a func necessaria
func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fibonacci(task)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
