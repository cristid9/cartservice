package productservice

import (
	"errors"
	"lotterProject/cart"
)

// ProductService To save the trouble of connecting this simple API to a database and fetch real products from
// that database, we made this mock service with some hard codeed products
type ProductService struct {
	products map[string]cart.Product
}

var mockProducts = map[string]cart.Product {
	"1" : {"1", "Pepsi 2.5 L", 5.6 },
	"2" : {"2", "Coca Cola 2.5 L", 5.4},
	"3" : {"3", "Giusto 2L", 4.5},
}

func NewProductService() ProductService {
	return ProductService{products: mockProducts}
}

func (p ProductService) GetProduct(id string) (cart.Product, error) {
	product, ok := p.products[id]
	if !ok {
		return cart.Product{}, errors.New("couldn't find product")
	}

	return product, nil
}
