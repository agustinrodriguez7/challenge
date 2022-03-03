package handler

import (
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/dbModel"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/response"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/repository"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type (
	GetCartHandler interface {
		HandleGetCartRequest(engine *gin.Context)
	}
	GetCartHandlerImpl struct {
		cartRepository repository.CartRepository
		logger         *logrus.Logger
	}
)

func NewGetApiCartHandler() GetCartHandler {
	return GetCartHandlerImpl{
		cartRepository: repository.NewCartRepositoryImpl(),
		logger:         utils.GetLogger(),
	}
}

func (gchi GetCartHandlerImpl) HandleGetCartRequest(context *gin.Context) {
	gchi.logger.Info("Starting to get cart.")
	resp := response.GetCartResponse{}
	clientId, err := gchi.getClientIdFromInputData(context)

	products, err := gchi.cartRepository.GetCart(*clientId)
	if err != nil {
		gchi.logger.
			Errorf("Error trying to get cart with error: %+v, and description: %+v for the client: %+v ",
				err, err.Error(), *clientId)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if products == nil || len(*products) == 0 {
		resp.Products = []dbModel.Product{}
	} else {
		resp.Products = *products
	}
	context.JSON(http.StatusOK, resp)
	gchi.logger.Infof("Success getting cart for client: %+v", *clientId)
}

func (gchi GetCartHandlerImpl) getClientIdFromInputData(context *gin.Context) (*int, error) {
	clientId := context.GetHeader(CLIENT_ID)
	formattedClientId, err := strconv.Atoi(clientId)
	if err != nil {
		gchi.logger.Errorf("Error trying to format clientId: %+v, with error: %+v", clientId, err)
		return nil, err
	}
	return &formattedClientId, nil
}
