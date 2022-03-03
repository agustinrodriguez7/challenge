package response

import (
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/dbModel"
)

type GetOrderResponse struct {
	Products []dbModel.Product `json:"products"`
}
