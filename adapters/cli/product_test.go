package cli_test

import (
	"fmt"
	"testing"

	"github.com/Daniel-Vinicius/go-ports-adapters/adapters/cli"
	mock_application "github.com/Daniel-Vinicius/go-ports-adapters/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(testing *testing.T) {
	goMockController := gomock.NewController(testing)
	defer goMockController.Finish()

	productId := "123"
	productName := "Shirt"
	productPrice := 25.99
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(goMockController)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(goMockController)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf("Product ID %s with the name %s has been created with price %f and status %s", productId, productName, productPrice, productStatus)
	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)

	require.Nil(testing, err)
	require.Equal(testing, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(serviceMock, "enable", productId, "", 0)

	require.Nil(testing, err)
	require.Equal(testing, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(serviceMock, "disable", productId, "", 0)

	require.Nil(testing, err)
	require.Equal(testing, expectedResult, result)

	expectedResult = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", productId, productName, productPrice, productStatus)
	result, err = cli.Run(serviceMock, "get", productId, "", 0)

	require.Nil(testing, err)
	require.Equal(testing, expectedResult, result)
}
