package entity

const (
	Kitchen = iota + 1
	Guest
)

type IPrinter interface {
	SetName(name string)
	SetApiKey(api_key string)
	SetPointID(point_id int)
	ApiKey() string
	Name() string
	PointID() int
	Type() int
}

type Printer struct {
	name     string
	api_key  string
	point_id int
}

func (p Printer) SetName(name string) {
	p.name = name
}

func (p Printer) SetApiKey(api_key string) {
	p.api_key = api_key
}

func (p Printer) SetPointID(point_id int) {
	p.point_id = point_id
}

func (p Printer) ApiKey() string {
	return p.api_key
}

func (p Printer) Name() string {
	return p.name
}

func (p Printer) PointID() int {
	return p.point_id
}

type KitchenPrinter struct {
	Printer
}

func (k KitchenPrinter) Type() int {
	return Kitchen
}

type GuestPrinter struct {
	Printer
}

func (g GuestPrinter) Type() int {
	return Guest
}
