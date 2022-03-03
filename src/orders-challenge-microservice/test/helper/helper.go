package helper

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func CreateGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	return c, recorder
}

func CreateGinContextWithRequest(method string, url string, body []byte, headers map[string]string, params map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
	c, recorder := CreateGinContext()
	c.Request, _ = http.NewRequest(method, url, bytes.NewReader(body))
	c.Request.Header = map[string][]string{}
	for key, value := range headers {
		c.Request.Header.Add(key, value)
	}
	ginParams := []gin.Param{}
	for key, value := range params {
		ginParams = append(ginParams, gin.Param{Key: key, Value: value})
	}
	c.Params = ginParams
	return c, recorder
}
