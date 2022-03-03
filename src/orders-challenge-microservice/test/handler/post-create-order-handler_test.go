package handler

import (
	"encoding/json"
	"errors"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/response"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/handler"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/test/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Post Create Order Tests", func() {
	var (
		postCreateOrderHandler      handler.PostCreateOrderHandler
		orderRepositoryMockInstance OrderRepositoryImplMock
		orderId                     int64
	)
	Context("Handler tests", func() {
		It("Happy Path", func() {
			orderRepositoryMockInstance.On("CreateOrder").
				Return(&orderId, nil)
			postCreateOrderHandler = handler.NewPostCreateOrderHandlerWithParams(orderRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, nil)
			postCreateOrderHandler.HandlePostCreateOrderRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusOK))
			bodyResponse := response.CreateOrderResponse{}
			_ = json.Unmarshal([]byte(recorder.Body.String()), &bodyResponse)
			Expect(bodyResponse.OrderId).To(Equal(int64(1)))
		})
		It("With no clientId header should return http bad request status", func() {
			orderRepositoryMockInstance.On("CreateOrder").
				Return(&orderId, nil)
			postCreateOrderHandler = handler.NewPostCreateOrderHandlerWithParams(orderRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, nil, nil)
			postCreateOrderHandler.HandlePostCreateOrderRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("If query return error != nil should return http internal server error status", func() {
			orderRepositoryMockInstance.On("CreateOrder").
				Return(&orderId, errors.New("Something went wrong"))
			postCreateOrderHandler = handler.NewPostCreateOrderHandlerWithParams(orderRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, nil)
			postCreateOrderHandler.HandlePostCreateOrderRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusInternalServerError))
			Expect(recorder.Body.String()).To(Equal(""))
		})

	})

	BeforeEach(func() {
		orderId = 1
		orderRepositoryMockInstance = OrderRepositoryImplMock{}
	})

})
