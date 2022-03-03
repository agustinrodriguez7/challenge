package config

import (
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/middleware"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/router"
	"github.com/gin-gonic/gin"
)

func InitApp() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.MainMiddleware)
	router.ConfigureRouter(engine)
	engine.Run()
}
