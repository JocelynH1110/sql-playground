package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/jocelynh1110/sql-playground/models"
	_ "github.com/lib/pq"
)

func main() {
	// 連接localhost
	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@127.0.0.1:5432/playground")
	if err != nil {
		log.Fatalln(err)
	}

	rows := db.QueryRowx("select extract(year from now())")
	var result int
	err = rows.Scan(&result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)

	products, err := models.ListProducts(db)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(products)

	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))

}
