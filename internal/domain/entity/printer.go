package entity

const (
	Kitchen = iota + 1
	Guest
)

type Printer struct {
	apiKey      string
	name        string
	pointId     int
	printerType int
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
