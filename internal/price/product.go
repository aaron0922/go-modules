package price

import (
	"errors"
)

type Product struct {
	ID int
	Name string
}

var products = map[int]Product{
	1: Product{ID: 1, Name: "Product 1"},
	2: Product{ID: 2, Name: "Product 2"},
	3: Product{ID: 3, Name: "Product 3"},
}


func GetProductByID(id int) (*Product, error) {
	if id > 3 {
		return nil, errors.New("Product not found")
	}
	
	product := products[id]
	return &product, nil
}