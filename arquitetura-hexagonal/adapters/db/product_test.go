package db_test

import (
	"database/sql"
	"log"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db * sql.DB) {
	table := `
		CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string
		);
	`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `
		INSERT INTO products VALUES (
			"abc",
			"Product Test",
			0,
			"disabled"
		)
	`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setup()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
}