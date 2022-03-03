package router

const (
	//get endpoints
	getApiCart = "/api/cart"
	getOrder = "/api/order/:orderId"

	//post endpoints
	postCartAddProduct = "/api/cart/products/:productId"
	postCreateOrder = "/api/order"
)
