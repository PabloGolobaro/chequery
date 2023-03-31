package entity

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOrderJson(t *testing.T) {
	order := Order{
		PointID: 1,
		Products: []Product{
			{Name: "Meat", Quantity: 3, Price: 145},
			{Name: "vegetables", Quantity: 2, Price: 32},
			{Name: "Juice", Quantity: 1, Price: 48}},
	}
	marshal, err := json.MarshalIndent(&order, "", "\t")
	if err != nil {
		return
	}
	fmt.Fprint(os.Stdout, string(marshal))

}

func TestOrderDetails_Details(t *testing.T) {
	order := Order{
		PointID: 1,
		Products: []Product{
			{Name: "Meat", Quantity: 3, Price: 145},
			{Name: "vegetables", Quantity: 2, Price: 32},
			{Name: "Juice", Quantity: 1, Price: 48}},
	}

	tests := []struct {
		name   string
		order  Order
		result string
	}{
		{name: "simple", order: order, result: "0. Name: Meat, price: 145, quantity: 3\n1. Name: vegetables, price: 32, quantity: 2\n2. Name: Juice, price: 48, quantity: 1\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.order.Details()
			if !assert.Equal(t, tt.result, got) {
				t.Errorf("Want %v got %v", tt.result, got)
			}
		})
	}
}
