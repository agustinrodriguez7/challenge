package repository

import (
	"database/sql"
	"fmt"
	"github.com/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice/dto/dbModel"
	"os"
)

type (
	CartRepository interface {
		AddProductToCart(productId, clientId int) error
		GetCart(clientInd int) (*[]dbModel.Product, error)
	}
	CartRepositoryImpl struct {
		dbClient DBClient
	}
)

const (
	INSERT_PRODUCT_TO_CART_QUERY              = "INSERT_PRODUCT_TO_CART_QUERY"
	SELECT_PRODUCTS_FROM_PRODUCTS_CARTS_QUERY = "SELECT_PRODUCTS_FROM_PRODUCTS_CARTS_QUERY"
)

func NewCartRepositoryImpl() CartRepository {
	return CartRepositoryImpl{
		dbClient: NewDBClient(),
	}
}

func (cri CartRepositoryImpl) AddProductToCart(productId, clientId int) error {
	client, err := cri.dbClient.GetClient()
	if err != nil {
		return err
	}
	query := fmt.Sprintf(os.Getenv(INSERT_PRODUCT_TO_CART_QUERY), productId, clientId)
	_, err = client.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (cri CartRepositoryImpl) GetCart(clientId int) (*[]dbModel.Product, error) {
	client, err := cri.dbClient.GetClient()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(os.Getenv(SELECT_PRODUCTS_FROM_PRODUCTS_CARTS_QUERY), clientId)
	rows, err := client.Query(query)
	if err != nil {
		return nil, err
	}

	return cri.getProductsFromQueryRows(rows)

}

func (cri CartRepositoryImpl) getProductsFromQueryRows(rows *sql.Rows) (*[]dbModel.Product, error) {
	productsCart := new([]dbModel.Product)
	var id, categoryId, type_ int
	var label, downloadUrl sql.NullString
	var weight sql.NullFloat64
	var productCart dbModel.Product
	for rows.Next() {
		err := rows.Scan(&id, &categoryId, &label, &type_, &downloadUrl, &weight)
		if err != nil {
			return nil, err
		} else {
			productCart = dbModel.Product{
				Id:          id,
				CategoryId:  categoryId,
				Label:       label.String,
				Type:        type_,
				DownloadUrl: downloadUrl.String,
				Weight:      weight.Float64,
			}
			*productsCart = append(*productsCart, productCart)
		}
	}
	return productsCart, nil
}
