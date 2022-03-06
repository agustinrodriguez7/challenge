package handler

import (
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/dbModel"
	"github.com/stretchr/testify/mock"
)

type CartRepositoryImplMock struct {
	mock.Mock
}

func (crim CartRepositoryImplMock) GetCart(clientInd int) (*[]dbModel.Product, error) {
	args := crim.Called()
	return args.Get(0).(*[]dbModel.Product), args.Error(1)
}

func (crim CartRepositoryImplMock) AddProductToCart(productId, clientId int) error {
	args := crim.Called()
	return args.Error(0)
}

type OrderRepositoryImplMock struct {
	mock.Mock
}

func (crim OrderRepositoryImplMock) CreateOrder(clientId int) (*int64, error) {
	args := crim.Called()
	return args.Get(0).(*int64), args.Error(1)
}

func (crim OrderRepositoryImplMock) GetOrder(clientId, orderId int) (*[]dbModel.Product, error) {
	args := crim.Called()
	return args.Get(0).(*[]dbModel.Product), args.Error(1)
}

