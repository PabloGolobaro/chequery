package entity

type OrderDetails struct {
	pointID int    `json:"pointID,omitempty"`
	details string `json:"details,omitempty"`
}

func (o OrderDetails) Details() string {
	return o.details
}

func (o OrderDetails) PointId() int {
	return o.pointID
}
