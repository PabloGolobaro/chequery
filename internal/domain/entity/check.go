package entity

type ICheck interface {
}

type OrderCheck struct {
	id        int    `json:"id,omitempty"`
	printerId string `json:"printer_id,omitempty"`
	order     string `json:"order"`
	status    string `json:"status,omitempty"`
	filePath  string `json:"file_path,omitempty"`
	checkType int    `json:"check_type,omitempty"`
}

func NewOrderCheck() *OrderCheck {
	return &OrderCheck{}
}

func (c *OrderCheck) OrderDetails(details OrderDetails) *OrderCheck {
	c.order = details.details
	return c
}

func (c *OrderCheck) Printer(printer Printer) *OrderCheck {
	c.printerId = printer.ApiKey()
	c.checkType = printer.PrinterType()
	c.status = "created"
	return c
}

func (c OrderCheck) Id() int {
	return c.id
}

func (c OrderCheck) SetId(id int) {
	c.id = id
}

func (c OrderCheck) PrinterId() string {
	return c.printerId
}

func (c OrderCheck) SetPrinterId(printerId string) {
	c.printerId = printerId
}

func (c OrderCheck) Order() string {
	return c.order
}

func (c OrderCheck) SetOrder(order string) {
	c.order = order
}

func (c OrderCheck) Status() string {
	return c.status
}

func (c OrderCheck) SetStatus(status string) {
	c.status = status
}

func (c OrderCheck) FilePath() string {
	return c.filePath
}

func (c OrderCheck) SetFilePath(pdfFileName string) {
	c.filePath = pdfFileName
}

func (c OrderCheck) CheckType() int {
	return c.checkType
}

func (c OrderCheck) SetCheckType(checkType int) {
	c.checkType = checkType
}
