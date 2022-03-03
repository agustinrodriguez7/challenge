package handler

import (
	"errors"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/handler"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/test/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Post Add Cart Products Handler Tests", func() {
	var (
		postAddProductsCartHandler handler.PostAddProductsCartHandler
		cartRepositoryMockInstance CartRepositoryImplMock
	)
	Context("Handler tests", func() {
		It("Happy Path", func() {
			cartRepositoryMockInstance.On("AddProductToCart").
				Return(nil)
			postAddProductsCartHandler = handler.NewPostApiCartProductsHandlerWithParams(cartRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, map[string]string{"productId": "1"})
			postAddProductsCartHandler.HandlePostAddProductsCartRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusOK))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("With no clientId header should return http bad request status", func() {
			cartRepositoryMockInstance.On("AddProductToCart").
				Return(nil)
			postAddProductsCartHandler = handler.NewPostApiCartProductsHandlerWithParams(cartRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, nil, map[string]string{"productId": "1"})
			postAddProductsCartHandler.HandlePostAddProductsCartRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("With no productId path param should return http bad request status", func() {
			cartRepositoryMockInstance.On("AddProductToCart").
				Return(nil)
			postAddProductsCartHandler = handler.NewPostApiCartProductsHandlerWithParams(cartRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, nil)
			postAddProductsCartHandler.HandlePostAddProductsCartRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusBadRequest))
			Expect(recorder.Body.String()).To(Equal(""))
		})

		It("If query return error != nil should return http internal server error status", func() {
			cartRepositoryMockInstance.On("AddProductToCart").
				Return(errors.New("Something went wrong"))
			postAddProductsCartHandler = handler.NewPostApiCartProductsHandlerWithParams(cartRepositoryMockInstance)
			context, recorder := helper.
				CreateGinContextWithRequest("", "", nil, map[string]string{"clientId": "1"}, map[string]string{"productId": "1"})
			postAddProductsCartHandler.HandlePostAddProductsCartRequest(context)

			Expect(context.Writer.Status()).To(Equal(http.StatusInternalServerError))
			Expect(recorder.Body.String()).To(Equal(""))
		})

	})

	BeforeEach(func() {
		cartRepositoryMockInstance = CartRepositoryImplMock{}
	})

})
