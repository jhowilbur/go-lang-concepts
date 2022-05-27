package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile("sites")
	fmt.Println("-----------------------")
	writeFile("write", "write2")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(file string) {
	// read full file
	dat, err := os.ReadFile(file)
	check(err)
	fmt.Print(string(dat), "\n")

	// if you need to define a limit bytes quantity
	f, err := os.Open(file)
	check(err)
	fmt.Println(f)        // print memory pointer
	b1 := make([]byte, 5) // in this case with 5 bytes
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}

func writeFile(file string, file2 string) {
	d1 := []byte("hello\ngo\n") // text to write
	err := os.WriteFile(file, d1, 0644)
	check(err)
	readFile(file) // the func WriteFile close automatically os

	// if you need a granular writes, follow the steps below
	f, err := os.Create(file2)
	check(err)
	defer f.Close() // waiting a Sync command to execute the close

	d2 := []byte{115, 111, 109, 101, 10} // write from bytes to string ASCII table
	n2, err := f.Write(d2)
	check(err)
	fmt.Println("====================")
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n") // write directly a string
	check(err)
	fmt.Println("====================")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() // closing the "defer" func

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Println("====================")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush() // to conclude writing in file without this you can't send to file the information

	fmt.Println("*******************")
	readFile(file2)
}
