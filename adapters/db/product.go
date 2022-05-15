package db

import (
	"database/sql"

	"github.com/Daniel-Vinicius/go-ports-adapters/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductPersistenceSQLite struct {
	db *sql.DB
}

func NewProductPersistenceSQLite(db *sql.DB) *ProductPersistenceSQLite {
	return &ProductPersistenceSQLite{db: db}
}

func (productPersistenceSQLite *ProductPersistenceSQLite) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	statement, err := productPersistenceSQLite.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (productPersistenceSQLite *ProductPersistenceSQLite) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	productPersistenceSQLite.db.QueryRow("SELECT id from products where id = ?", product.GetID()).Scan(&rows)

	if rows == 0 {
		_, err := productPersistenceSQLite.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := productPersistenceSQLite.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (productPersistenceSQLite *ProductPersistenceSQLite) create(product application.ProductInterface) (application.ProductInterface, error) {
	statement, err := productPersistenceSQLite.db.Prepare("INSERT INTO products(id, name, price, status) values (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	statement.Close()
	return product, nil
}

func (productPersistenceSQLite *ProductPersistenceSQLite) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := productPersistenceSQLite.db.Exec(
		"UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

	if err != nil {
		return nil, err
	}

	return product, nil
}
