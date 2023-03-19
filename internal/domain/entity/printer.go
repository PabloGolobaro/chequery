package entity

const (
	Kitchen = iota + 1
	Guest
)

type Printer struct {
	ApiKey      string `db:"api_key"`
	Name        string `db:"name"`
	PointId     int    `db:"point_id"`
	PrinterType int    `db:"printer_type"`
}

func (p Printer) GetApiKey() string {
	return p.ApiKey
}

func (p Printer) GetName() string {
	return p.Name
}

func (p Printer) GetPointId() int {
	return p.PointId
}

func (p Printer) GetPrinterType() int {
	return p.PrinterType
}
