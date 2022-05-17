package main

import "fmt"

func main() {
	fmt.Println("Go Routines")
	go print("Go")
	print("loop")
}

func print(text string) {
	for {
		fmt.Println(text)
	}
}
