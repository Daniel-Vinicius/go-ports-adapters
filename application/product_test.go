package application_test

import (
	"testing"

	"github.com/Daniel-Vinicius/go-ports-adapters/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(testing *testing.T) {
	product := application.Product{}
	product.Name = "Shirt"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Equal(testing, product.GetStatus(), application.ENABLED)
	require.Nil(testing, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(testing, "product price must be greater than 0 to enable the product", err.Error())
}

func TestProduct_Disable(testing *testing.T) {
	product := application.Product{}
	product.Name = "Shirt"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Equal(testing, product.GetStatus(), application.DISABLED)
	require.Nil(testing, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(testing, "product price must be 0 to disable the product", err.Error())
}

func TestProduct_IsValid(testing *testing.T) {
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Shirt"
	product.Status = application.DISABLED
	product.Price = 10

	valid, err := product.IsValid()

	require.Nil(testing, err)
	require.Equal(testing, true, valid)

	product.Status = "invalid"
	valid, err = product.IsValid()
	require.Equal(testing, "the product status must be enabled or disabled", err.Error())
	require.Equal(testing, false, valid)

	product.Status = application.ENABLED
	valid, err = product.IsValid()
	require.Nil(testing, err)
	require.Equal(testing, true, valid)

	product.Price = -10
	valid, err = product.IsValid()
	require.Equal(testing, "the product price must be greater or equal to 0", err.Error())
	require.Equal(testing, false, valid)
}

func TestProduct_GetID(testing *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()

	require.Equal(testing, product.ID, product.GetID())
}

func TestProduct_GetName(testing *testing.T) {
	product := application.Product{}
	product.Name = "Shirt"

	require.Equal(testing, product.Name, product.GetName())
}

func TestProduct_GetStatus(testing *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED

	require.Equal(testing, product.Status, product.GetStatus())
}

func TestProduct_GetPrice(testing *testing.T) {
	product := application.Product{}
	product.Price = 10

	require.Equal(testing, product.Price, product.GetPrice())
}
