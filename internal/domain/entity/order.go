package entity

import (
	"encoding/json"
	"fmt"
)

type OrderDetails struct {
	PointID int
	Order   string
}

func (o OrderDetails) Details() string {
	var M = map[string]interface{}{}

	err := json.Unmarshal([]byte(o.Order), &M)
	if err != nil {
		return ""
	}

	res := ""
	for key, value := range M {
		res += fmt.Sprintf("%v - %v\n", key, value)
	}
	return res
}

func (o OrderDetails) PointId() int {
	return o.PointID
}
