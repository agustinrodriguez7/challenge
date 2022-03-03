package handler

import (
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/repository"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const (
	PRODUCT_ID = "productId"
	CLIENT_ID  = "clientId"
)

type (
	PostAddProductsCartHandler interface {
		HandlePostAddProductsCartRequest(engine *gin.Context)
	}
	PostAddProductsCartImpl struct {
		cartRepository repository.CartRepository
		logger         *logrus.Logger
	}
)

func NewPostApiCartProductsHandler() PostAddProductsCartHandler {
	return PostAddProductsCartImpl{
		cartRepository: repository.NewCartRepositoryImpl(),
		logger:         utils.GetLogger(),
	}
}

func NewPostApiCartProductsHandlerWithParams(cartRepository repository.CartRepository) PostAddProductsCartHandler {
	return PostAddProductsCartImpl{
		cartRepository: cartRepository,
		logger:         utils.GetLogger(),
	}
}

func (papci PostAddProductsCartImpl) HandlePostAddProductsCartRequest(context *gin.Context) {
	papci.logger.Info("Starting to add product.")
	productId, clientId, err := papci.getProductIdAndClientIdFromInputData(context)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = papci.cartRepository.AddProductToCart(*productId, *clientId)
	if err != nil {
		papci.logger.
			Errorf("Failed trying to add product: %+v to cart to the client: %+v with error: %+v, and message: %+v ",
				*productId, *clientId, err, err.Error())
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	context.Status(http.StatusOK)
	papci.logger.Infof("Success adding product: %+v to client: %+v", *productId, *clientId)
}

func (papci PostAddProductsCartImpl) getProductIdAndClientIdFromInputData(context *gin.Context) (*int, *int, error) {
	productId := context.Param(PRODUCT_ID)
	clientId := context.GetHeader(CLIENT_ID)
	formattedProductId, err := strconv.Atoi(productId)
	if err != nil {
		papci.logger.Errorf("Error trying to format productId: %+v, with error: %+v", productId, err)
		return nil, nil, err
	}
	formattedClientId, err := strconv.Atoi(clientId)
	if err != nil {
		papci.logger.Errorf("Error trying to format clientId: %+v, with error: %+v", clientId, err)
		return nil, nil, err
	}
	return &formattedProductId, &formattedClientId, nil
}
