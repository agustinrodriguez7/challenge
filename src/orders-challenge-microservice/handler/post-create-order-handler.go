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

type (
	PostCreateOrderHandler interface {
		HandlePostCreateOrderRequest(engine *gin.Context)
	}
	PostCreateOrderHandlerImpl struct {
		orderRepository repository.OrderRepository
		logger          *logrus.Logger
	}
)

func NewPostCreateOrderHandler() PostCreateOrderHandler {
	return PostCreateOrderHandlerImpl{
		orderRepository: repository.NewOrderRepository(),
		logger:          utils.GetLogger(),
	}
}

func (pcohi PostCreateOrderHandlerImpl) HandlePostCreateOrderRequest(context *gin.Context) {
	pcohi.logger.Info("Starting to create order.")
	clientId, err := pcohi.getClientIdFromInputData(context)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	orderId, err := pcohi.orderRepository.CreateOrder(*clientId)
	if err != nil {
		pcohi.logger.Errorf("Failed trying to create order to the client: %+v with error: %+v, and message: %+v ",
			*clientId, err, err.Error())
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, response.CreateOrderResponse{OrderId: *orderId})
	pcohi.logger.Infof("Success creating order to client: %+v, resulting with orderId: %+v", *clientId, *orderId)
}

func (pcohi PostCreateOrderHandlerImpl) getClientIdFromInputData(context *gin.Context) (*int, error) {
	clientId := context.GetHeader(CLIENT_ID)
	formattedClientId, err := strconv.Atoi(clientId)
	if err != nil {
		pcohi.logger.Errorf("Error trying to format clientId: %+v, with error: %+v", clientId, err)
		return nil, err
	}
	return &formattedClientId, err
}
