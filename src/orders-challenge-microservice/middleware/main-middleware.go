package middleware

import "github.com/gin-gonic/gin"

func MainMiddleware(context *gin.Context) {
	corsMiddleware(context)
	loggerMiddleware(context)
}
