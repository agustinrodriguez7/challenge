package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware(context *gin.Context) { //allow to receipt api calls
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	if context.Request.Method == "OPTIONS" { //Before to perform requests, consumers use to hit API with OPTIONS method,
		context.AbortWithStatus(http.StatusOK) // to get API information and show allowed methods, and more.
	}
}
