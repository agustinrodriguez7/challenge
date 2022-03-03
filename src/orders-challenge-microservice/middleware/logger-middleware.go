package middleware

import (
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/utils"
	"github.com/gin-gonic/gin"
)

func loggerMiddleware(context *gin.Context) {
	utils.GetLogger().Debugf("Request to path: %+v", context.Request.RequestURI) //mode info can be added
}
