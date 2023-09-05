package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:xjv123..@tcp(127.0.0.1:3306)/day09")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("db:::", db)

}
