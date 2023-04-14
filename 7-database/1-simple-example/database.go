package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	urlConnection := "golang:golang@/wilbur?charset=utf8&parseTime=True&loc=Local"
	db, error := sql.Open("mysql", urlConnection)

	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	if error = db.Ping(); error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connected")

	queryExecute, error := db.Query("select * from users")
	if error != nil {
		log.Fatal(error)
	}
	defer queryExecute.Close()

	fmt.Println(queryExecute)
}
