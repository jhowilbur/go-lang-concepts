package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	urlConnection := "golang:golang@/wilbur?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConnection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	fmt.Println("Connected")

	return db, erro
}
