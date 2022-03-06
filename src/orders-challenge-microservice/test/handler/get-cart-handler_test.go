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

var _ = Describe("Get Cart Handler Tests", func() {
	var (
		getCartHandler             handler.GetCartHandler
		cartRepositoryMockInstance CartRepositoryImplMock
		dbResponse                 []dbModel.Product
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
			cartRepositoryMockInstance.On("GetCart").
				Return(&dbResponse, nil)
			getCartHandler = handler.NewGetApiCartHandlerWithParams(cartRepositoryMockInstance)
			context, recorder := helper.CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, nil)
			getCartHandler.HandleGetCartRequest(context)

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
			getCartHandler = handler.NewGetApiCartHandlerWithParams(cartRepositoryMockInstance)
			context, recorder := helper.CreateGinContextWithRequest("", "", nil, nil, nil)
			getCartHandler.HandleGetCartRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("If query return error != nil should return http internal server error status", func() {
			cartRepositoryMockInstance.On("GetCart").
				Return(&dbResponse, errors.New("Something went wrong"))
			getCartHandler = handler.NewGetApiCartHandlerWithParams(cartRepositoryMockInstance)
			context, recorder := helper.CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, nil)
			getCartHandler.HandleGetCartRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusInternalServerError))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("If there is no database records should return http ok status and empty array in response", func() {
			dbResponse = []dbModel.Product{}
			cartRepositoryMockInstance.On("GetCart").
				Return(&dbResponse, nil)
			getCartHandler = handler.NewGetApiCartHandlerWithParams(cartRepositoryMockInstance)
			context, recorder := helper.CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, nil)
			getCartHandler.HandleGetCartRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusOK))
			bodyResponse := response.GetCartResponse{}
			_ = json.Unmarshal([]byte(recorder.Body.String()), &bodyResponse)
			Expect(len(bodyResponse.Products)).To(Equal(0))
		})

	})

	BeforeEach(func() {
		cartRepositoryMockInstance = CartRepositoryImplMock{}
	})

})