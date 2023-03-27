package entity

import (
	"fmt"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type Order struct {
	PointID  int       `json:"point_id"`
	Products []Product `json:"products"`
}

func (o Order) Details() string {
	result := ""
	for i, product := range o.Products {
		s := fmt.Sprintf("%v. Name: %v, price: %v, quantity: %v\n", i, product.Name, product.Price, product.Quantity)
		result += s
	}
	return result
}

func (o Order) PointId() int {
	return o.PointID
}
