package main

import (
	"fmt"
	"net/http"
)

func main() {

	// NET
	response, _ := http.Get("https://google.com")
	fmt.Println(response)
}
