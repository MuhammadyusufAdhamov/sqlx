package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	PostgresUser     = "postgres"
	PostgresDatabase = "crud"
	PostgresPassword = "7"
	PostgresHost     = "localhost"
	PostgresPort     = 5432
)

type Product struct {
	Id          int32     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Color       string    `db:"color"`
	Price       float64   `db:"price"`
	Category    string    `db:"category"`
	CreatedAt   time.Time `db:"created_at"`
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s database=%s password=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresDatabase,
		PostgresPassword,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("failed connect to database: %v", err)
	}

	query := `
		select 
		    id,
		    name,
		    description,
		    color,
		    price,
		    category,
		    created_at
from products 
where id = $1

`

	var product Product
	err = db.Get(&product, query, 28)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(product)

	//	queryProducts := `
	//		select
	//		    id,
	//		    name,
	//		    description,
	//		    color,
	//		    price,
	//		    category,
	//		    created_at
	//from products
	//limit 10
	//
	//`
	//
	//	var products []Product
	//	err = db.Select(&products, queryProducts)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	fmt.Println("Products")
	//	fmt.Println(products)
}
