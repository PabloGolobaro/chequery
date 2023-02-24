package entity

type ICheck interface {
	SetPrinterID(api_key string)
	SetOrder(order string)
	SetStatus(status string)
	SetFileName(filename string)
	Status() string
	Order() string
	PrinterID() string
	FileName() string
}

type Check struct {
	printer_id  string
	order       string
	status      string
	pdfFileName string
}

func (c *Check) Status() string {
	return c.status
}

func (c *Check) Order() string {
	return c.order
}

func (c *Check) PrinterID() string {
	return c.printer_id
}

func (c *Check) FileName() string {
	return c.pdfFileName
}

func (c *Check) SetFileName(filename string) {
	c.pdfFileName = filename
}

func (c *Check) SetPrinterID(api_key string) {
	c.printer_id = api_key
}

func (c *Check) SetOrder(order string) {
	c.order = order
}

func (c *Check) SetStatus(status string) {
	c.status = status
}

type KitchenCheck struct {
	Check
}

type GuestCheck struct {
	Check
}
