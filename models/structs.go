package models

import "github.com/jmoiron/sqlx"

type Product struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Price int32  `db:"price" json:"price"`
}

func ListProducts(db *sqlx.DB) ([]Product, error) {
	result := []Product{}
	err := db.Select(&result, "select id,name,price from products")
	return result, err
}
