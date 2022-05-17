package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // informa quantidade de Go Routines que vao estar rodando ao mesmo tempo

	// para rodar em paralelo
	// criar func anonima, mas dessa vez transformando ela em go routine (comecar palavra com go)
	go func() {
		print("Func 1")
		waitGroup.Done() // done tira um cara do -> waitGroup.Add(2) -> -1 -> waitGroup.Add(1)
	}()

	go func() {
		print("Func 2")
		waitGroup.Done()
	}()

	waitGroup.Wait() // espera executar todas as func ou seja, quando for -> waitGroup.Add(0)
}

func print(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
