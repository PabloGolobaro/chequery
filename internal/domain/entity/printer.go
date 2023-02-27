package entity

const (
	Kitchen = iota + 1
	Guest
)

type Printer struct {
	apiKey      string `json:"api_key,omitempty"`
	name        string `json:"name,omitempty"`
	pointId     int    `json:"point_id,omitempty"`
	printerType int    `json:"printer_type,omitempty"`
}

func (p Printer) ApiKey() string {
	return p.apiKey
}

func (p Printer) Name() string {
	return p.name
}

func (p Printer) PointId() int {
	return p.pointId
}

func (p Printer) PrinterType() int {
	return p.printerType
}
