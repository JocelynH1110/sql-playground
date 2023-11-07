package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Product struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Price int32  `db:"price"`
}

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

	products := []Product{}
	err = db.Select(&products, "select id,name,price from products")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(products)
}
