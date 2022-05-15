package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Daniel-Vinicius/go-ports-adapters/adapters/db"
	"github.com/Daniel-Vinicius/go-ports-adapters/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string, "name" string, "price" float, "status" string);`
	statement, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products values ("1", "Shirt", 0, "disabled");`
	statement, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
}

func TestProductDb_Get(testing *testing.T) {
	setUp()
	defer Db.Close()

	productPersistenceSQLite := db.NewProductPersistenceSQLite(Db)

	product, err := productPersistenceSQLite.Get("1")
	require.Nil(testing, err)
	require.Equal(testing, "Shirt", product.GetName())
	require.Equal(testing, 0.0, product.GetPrice())
	require.Equal(testing, "disabled", product.GetStatus())
}

func TestProductDb_Save(testing *testing.T) {
	setUp()
	defer Db.Close()

	productPersistenceSQLite := db.NewProductPersistenceSQLite(Db)

	product := application.NewProduct()
	product.Name = "Shirt"
	product.Price = 25

	productSaved, err := productPersistenceSQLite.Save(product)
	require.Nil(testing, err)

	require.Equal(testing, product.GetName(), productSaved.GetName())
	require.Equal(testing, product.GetPrice(), productSaved.GetPrice())
	require.Equal(testing, product.GetStatus(), productSaved.GetStatus())

	product.Status = "enabled"
	productUpdated, err := productPersistenceSQLite.Save(product)
	require.Nil(testing, err)
	require.Equal(testing, product.GetStatus(), productUpdated.GetStatus())
}
