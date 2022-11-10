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

	args := map[string]interface{}{
		"limit":  10,
		"offset": 0,
		"color":  "black",
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
	where color=:color
	limit :limit offset :offset
	
	`

	rows, err := db.NamedQuery(query, args)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Color,
			&product.Price,
			&product.Category,
			&product.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		products = append(products, product)
	}
	fmt.Println(products)
}
