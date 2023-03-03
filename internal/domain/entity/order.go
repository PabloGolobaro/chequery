package entity

import "fmt"

type OrderDetails struct {
	PointID int
	M       map[string]interface{}
}

func (o OrderDetails) Details() string {
	res := ""
	for key, value := range o.M {
		res += fmt.Sprintf("%v - %v\n", key, value)
	}
	return res
}

func (o OrderDetails) PointId() int {
	return o.PointID
}
