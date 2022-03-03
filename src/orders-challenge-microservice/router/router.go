package router

import (
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	getApiOrderHandler  handler.GetOrderHandler
	getApiCardHandler   handler.GetCartHandler
	postApiCartProducts handler.PostAddProductsCartHandler
	postApiCreateOrder  handler.PostCreateOrderHandler
)

func ConfigureRouter(engine *gin.Engine) {
	initHandlers()
	getRouter(engine)
	postRouter(engine)

}

func getRouter(engine *gin.Engine) {
	engine.Handle(http.MethodGet, getApiCart, getApiCardHandler.HandleGetCartRequest)
	engine.Handle(http.MethodGet, getOrder, getApiOrderHandler.HandleGetOrderRequest)
}

func postRouter(engine *gin.Engine) {
	engine.Handle(http.MethodPost, postCartAddProduct, postApiCartProducts.HandlePostAddProductsCartRequest)
	engine.Handle(http.MethodPost, postCreateOrder, postApiCreateOrder.HandlePostCreateOrderRequest)
}

func initHandlers() {
	getApiOrderHandler = handler.NewGetOrderHandler()
	getApiCardHandler = handler.NewGetApiCartHandler()
	postApiCartProducts = handler.NewPostApiCartProductsHandler()
	postApiCreateOrder = handler.NewPostCreateOrderHandler()
}
