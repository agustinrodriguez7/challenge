package handler

import (
	"encoding/json"
	"errors"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/dbModel"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/response"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/handler"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/test/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Get Order Handler Tests", func() {
	var (
		getOrderHandler             handler.GetOrderHandler
		orderRepositoryMockInstance OrderRepositoryImplMock
		dbResponse                  []dbModel.Product
	)
	Context("Handler tests", func() {
		It("Happy Path", func() {
			dbResponse = []dbModel.Product{{
				Id:          0,
				CategoryId:  1,
				Label:       "a label",
				Type:        2,
				DownloadUrl: "",
				Weight:      3,
			}}
			orderRepositoryMockInstance.On("GetOrder").
				Return(&dbResponse, nil)
			getOrderHandler = handler.NewGetOrderHandlerWithParams(orderRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, map[string]string{"orderId": "1"})
			getOrderHandler.HandleGetOrderRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusOK))
			bodyResponse := response.GetCartResponse{}
			_ = json.Unmarshal([]byte(recorder.Body.String()), &bodyResponse)
			Expect(len(bodyResponse.Products)).To(Equal(1))
			productResponse := bodyResponse.Products[0]
			Expect(productResponse.Id).To(Equal(0))
			Expect(productResponse.CategoryId).To(Equal(1))
			Expect(productResponse.Label).To(Equal("a label"))
			Expect(productResponse.Type).To(Equal(2))
			Expect(productResponse.DownloadUrl).To(Equal(""))
			Expect(productResponse.Weight).To(Equal(float64(3)))
		})

		It("With no clientId header should return http bad request status", func() {
			dbResponse = []dbModel.Product{{
				Id:          0,
				CategoryId:  1,
				Label:       "a label",
				Type:        2,
				DownloadUrl: "",
				Weight:      3,
			}}
			orderRepositoryMockInstance.On("GetOrder").
				Return(&dbResponse, nil)
			getOrderHandler = handler.NewGetOrderHandlerWithParams(orderRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, nil, map[string]string{"orderId": "1"})
			getOrderHandler.HandleGetOrderRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("With no orderId path param should return http bad request status", func() {
			dbResponse = []dbModel.Product{{
				Id:          0,
				CategoryId:  1,
				Label:       "a label",
				Type:        2,
				DownloadUrl: "",
				Weight:      3,
			}}
			orderRepositoryMockInstance.On("GetOrder").
				Return(&dbResponse, nil)
			getOrderHandler = handler.NewGetOrderHandlerWithParams(orderRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, nil)
			getOrderHandler.HandleGetOrderRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("If query return error != nil should return http internal server error status", func() {
			orderRepositoryMockInstance.On("GetOrder").
				Return(&dbResponse, errors.New("Something went wrong"))
			getOrderHandler = handler.NewGetOrderHandlerWithParams(orderRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, map[string]string{"orderId": "1"})
			getOrderHandler.HandleGetOrderRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusInternalServerError))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("If there is no database records should return http not found status and empty array in response", func() {
			dbResponse = []dbModel.Product{}
			orderRepositoryMockInstance.On("GetOrder").
				Return(&dbResponse, nil)
			getOrderHandler = handler.NewGetOrderHandlerWithParams(orderRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, map[string]string{"orderId": "1"})
			getOrderHandler.HandleGetOrderRequest(context)


			Expect(context.Writer.Status()).To(Equal(http.StatusNotFound))
			Expect(recorder.Body.String()).To(Equal(""))
		})
	})

	BeforeEach(func() {
		orderRepositoryMockInstance = OrderRepositoryImplMock{}
	})

})
