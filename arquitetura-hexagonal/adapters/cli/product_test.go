package cli_test

import (
	"testing"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/codeedu/go-hexagonal/adapters/cli"
	"github.com/stretchr/testify/require"
	mock_application "github.com/codeedu/go-hexagonal/application/mocks"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf(
		"Product ID %s with name %s has been created with price %f and status %s", 
		productId, 
		productName,
		productPrice,
		productStatus,
	)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, result, resultExpected)

	resultExpected = fmt.Sprintf(
		"Product %s has been enabled", 
		productName,
	)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, result, resultExpected)

	resultExpected = fmt.Sprintf(
		"Product %s has been disabled", 
		productName,
	)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, result, resultExpected)

	resultExpected = fmt.Sprintf(
			"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			productId, 
			productName,
			productPrice,
			productStatus,
		)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, result, resultExpected)
}