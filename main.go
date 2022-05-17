package main

import (
	"database/sql"
	"fmt"

	db2 "github.com/Daniel-Vinicius/go-ports-adapters/adapters/db"
	"github.com/Daniel-Vinicius/go-ports-adapters/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductPersistenceSQLite(db)

	productService := application.NewProductService(productDbAdapter)

	productCreated, _ := productService.Create("Shirt", 30)

	productService.Enable(productCreated)

	fmt.Print(productCreated)
}
