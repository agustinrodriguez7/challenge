package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/dbModel"
	"os"
)

type (
	OrderRepository interface {
		CreateOrder(clientId int) (*int64, error)
		GetOrder(clientId, orderId int) (*[]dbModel.Product, error)
	}
	OrderRepositoryImpl struct {
		dbClient DBClient
	}
)

const (
	SELECT_PRODUCTS_FROM_PRODUCTS_ORDERS_QUERY   = "SELECT_PRODUCTS_FROM_PRODUCTS_ORDERS_QUERY"
	INSERT_ORDERS_QUERY                          = "INSERT_ORDERS_QUERY"
	SELECT_ORDER_ID_QUERY                        = "SELECT_ORDER_ID_QUERY"
	SELECT_PRODUCTS_ID_FROM_PRODUCTS_CARTS_QUERY = "SELECT_PRODUCTS_ID_FROM_PRODUCTS_CARTS_QUERY"
	INSERT_PRODUCTS_ORDERS_QUERY                 = "INSERT_PRODUCTS_ORDERS_QUERY"
	DELETE_PRODUCTS_CARTS_QUERY                  = "DELETE_PRODUCTS_CARTS_QUERY"
)

func NewOrderRepository() OrderRepository {
	return OrderRepositoryImpl{dbClient: NewDBClient()}
}

func (ori OrderRepositoryImpl) CreateOrder(clientId int) (*int64, error) {
	orderId := new(int64)
	client, err := ori.dbClient.GetClient()
	if err != nil {
		return nil, err
	}

	insertQuery := fmt.Sprintf(os.Getenv(INSERT_ORDERS_QUERY), clientId)
	_, err = client.Exec(insertQuery)
	if err != nil {
		return nil, err
	}

	getOrderIdQuery := fmt.Sprintf(os.Getenv(SELECT_ORDER_ID_QUERY), clientId)

	err = client.QueryRow(getOrderIdQuery).Scan(orderId)
	if err != nil {
		return nil, err
	}

	getProductsCartsQuery := fmt.Sprintf(os.Getenv(SELECT_PRODUCTS_ID_FROM_PRODUCTS_CARTS_QUERY), clientId)

	productsCart, err := client.Query(getProductsCartsQuery)
	if err != nil {
		return nil, err
	}
	productsId := []int{}
	auxProductId := new(int)

	for productsCart.Next() {
		err = productsCart.Scan(auxProductId)
		if err != nil {
			return nil, err
		}
		productsId = append(productsId, *auxProductId)
	}

	if len(productsId) == 0 {
		return nil, errors.New("There is no products in the cart to be added to Order") //this could be better handled with custom error and different http status code
	}

	for _, productId := range productsId {
		_, err = client.Exec(fmt.Sprintf(os.Getenv(INSERT_PRODUCTS_ORDERS_QUERY), productId, *orderId))
		if err != nil {
			return nil, err
		}
	}

	deleteProductsCartsQuery := fmt.Sprintf(os.Getenv(DELETE_PRODUCTS_CARTS_QUERY), clientId)
	client.Exec(deleteProductsCartsQuery)

	return orderId, nil
}

func (ori OrderRepositoryImpl) GetOrder(clientId, orderId int) (*[]dbModel.Product, error) {
	client, err := ori.dbClient.GetClient()
	if err != nil {
		return nil, err
	}

	getProductsCartsQuery := fmt.Sprintf(os.Getenv(SELECT_PRODUCTS_FROM_PRODUCTS_ORDERS_QUERY), clientId, orderId)

	rows, err := client.Query(getProductsCartsQuery)
	if err != nil {
		return nil, err
	}

	return ori.getProductsFromQueryRow(rows, err)

}

func (ori OrderRepositoryImpl) getProductsFromQueryRow(rows *sql.Rows, err error) (*[]dbModel.Product, error) {
	products := new([]dbModel.Product)
	var id, categoryId, type_ int
	var label, downloadUrl sql.NullString
	var weight sql.NullFloat64
	var product dbModel.Product
	for rows.Next() {
		err = rows.Scan(&id, &categoryId, &label, &type_, &downloadUrl, &weight)
		if err != nil {
			return nil, err
		} else {
			product = dbModel.Product{
				Id:          id,
				CategoryId:  categoryId,
				Label:       label.String,
				Type:        type_,
				DownloadUrl: downloadUrl.String,
				Weight:      weight.Float64,
			}
			*products = append(*products, product)
		}
	}
	return products, nil
}
