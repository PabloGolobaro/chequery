package entity

type ICheck interface {
}

type OrderCheck struct {
	id          int    `json:"id,omitempty"`
	printerId   string `json:"printer_id,omitempty"`
	order       string `json:"order"`
	status      string `json:"status,omitempty"`
	pdfFileName string `json:"pdf_file_name,omitempty"`
	checkType   string `json:"check_type,omitempty"`
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

func (c OrderCheck) PdfFileName() string {
	return c.pdfFileName
}

func (c OrderCheck) SetPdfFileName(pdfFileName string) {
	c.pdfFileName = pdfFileName
}

func (c OrderCheck) CheckType() string {
	return c.checkType
}

func (c OrderCheck) SetCheckType(checkType string) {
	c.checkType = checkType
}
