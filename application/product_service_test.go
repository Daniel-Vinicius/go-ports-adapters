package application_test

import (
	"testing"

	"github.com/Daniel-Vinicius/go-ports-adapters/application"
	mock_application "github.com/Daniel-Vinicius/go-ports-adapters/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(testing *testing.T) {
	goMockController := gomock.NewController(testing)
	defer goMockController.Finish()

	productMock := mock_application.NewMockProductInterface(goMockController)
	persistenceMock := mock_application.NewMockProductPersistenceInterface(goMockController)

	persistenceMock.EXPECT().Get(gomock.Any()).Return(productMock, nil).AnyTimes()

	productService := application.ProductService{
		Persistence: persistenceMock,
	}

	product, err := productService.Get("id")

	require.Nil(testing, err)
	require.Equal(testing, productMock, product)
}

func TestProductService_Create(testing *testing.T) {
	goMockController := gomock.NewController(testing)
	defer goMockController.Finish()

	productMock := mock_application.NewMockProductInterface(goMockController)
	persistenceMock := mock_application.NewMockProductPersistenceInterface(goMockController)

	persistenceMock.EXPECT().Save(gomock.Any()).Return(productMock, nil).AnyTimes()

	productService := application.ProductService{
		Persistence: persistenceMock,
	}

	product, err := productService.Create("Shirt", 99.99)

	require.Nil(testing, err)
	require.Equal(testing, productMock, product)

	invalidProduct, invalidProductError := productService.Create("", -299.99)

	require.NotNil(testing, invalidProductError)
	require.Empty(testing, invalidProduct)
}

func TestProductService_EnableDisable(testing *testing.T) {
	goMockController := gomock.NewController(testing)
	defer goMockController.Finish()

	productMock := mock_application.NewMockProductInterface(goMockController)
	persistenceMock := mock_application.NewMockProductPersistenceInterface(goMockController)

	persistenceMock.EXPECT().Save(gomock.Any()).Return(productMock, nil).AnyTimes()

	productService := application.ProductService{
		Persistence: persistenceMock,
	}

	productMock.EXPECT().Enable().Return(nil).AnyTimes()
	productMock.EXPECT().Disable().Return(nil).AnyTimes()

	_, err := productService.Enable(productMock)
	require.Nil(testing, err)

	_, err = productService.Disable(productMock)
	require.Nil(testing, err)
}
