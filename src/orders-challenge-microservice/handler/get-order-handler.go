package handler

import (
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/response"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/repository"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const (
	ORDER_ID = "orderId"
)

type (
	GetOrderHandler interface {
		HandleGetOrderRequest(engine *gin.Context)
	}
	GetOrderHandlerImpl struct {
		orderRepository repository.OrderRepository
		logger          *logrus.Logger
	}
)

func NewGetOrderHandler() GetOrderHandler {
	return GetOrderHandlerImpl{
		orderRepository: repository.NewOrderRepository(),
		logger:          utils.GetLogger(),
	}
}

func NewGetOrderHandlerWithParams(orderRepository repository.OrderRepository) GetOrderHandler {
	return GetOrderHandlerImpl{
		orderRepository: orderRepository,
		logger:          utils.GetLogger(),
	}
}

func (gohi GetOrderHandlerImpl) HandleGetOrderRequest(context *gin.Context) {
	gohi.logger.Info("Starting to get order.")
	resp := response.GetOrderResponse{}

	orderId, clientId, err := gohi.getOrderIdAndClientIdFromInputData(context)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	products, err := gohi.orderRepository.GetOrder(*clientId, *orderId)
	if err != nil {
		gohi.logger.
			Errorf("Failed trying to get order with error: %+v, message: %+v , orderID: %+v and clientId: %+v",
				err, err.Error(), *orderId, *clientId)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if products == nil || len(*products) == 0 {
		gohi.logger.
			Errorf("There is no products in the order: %+v with the client: %+v",
				*orderId, *clientId)
		context.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		resp.Products = *products
	}
	context.JSON(http.StatusOK, resp)
	gohi.logger.Infof("Success getting order: %+v for client: %+v", *orderId, *clientId)
}

func (gohi GetOrderHandlerImpl) getOrderIdAndClientIdFromInputData(context *gin.Context) (*int, *int, error) {
	orderId := context.Param(ORDER_ID)
	formattedOrderId, err := strconv.Atoi(orderId)
	if err != nil {
		gohi.logger.Errorf("Error trying to format orderId: %+v, with error: %+v", orderId, err)
		return nil, nil, err
	}

	clientId := context.GetHeader(CLIENT_ID)
	formattedClientId, err := strconv.Atoi(clientId)
	if err != nil {
		gohi.logger.Errorf("Error trying to format clientId: %+v, with error: %+v", clientId, err)
		return nil, nil, err
	}
	return &formattedOrderId, &formattedClientId, nil
}
