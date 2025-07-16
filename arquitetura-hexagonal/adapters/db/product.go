package db

import (
	"database/sql"
	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db * sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (productDb *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := productDb.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (productDb *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	productDb.db.QueryRow("select count(*) from products where id = ?", product.GetID()).Scan(&rows)

	if rows == 0 {
		_, err := productDb.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := productDb.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (productDb *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := productDb.db.Prepare(`insert into products(id, name, price, status) values(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productDb *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := productDb.db.Exec(
		`update products set name = ?, price = ?, status = ? where id = ?`,
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID(),
	)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return product, nil
}