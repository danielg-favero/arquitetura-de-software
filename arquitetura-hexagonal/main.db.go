package main

import (
	"database/sql"
	db2 "github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func mainDb() {
	// Testando o adapter do Banco de Dados no service de produtos
	db, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)
}