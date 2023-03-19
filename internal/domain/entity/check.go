package entity

type OrderCheck struct {
	Id        int    `db:"id"`
	PrinterId string `db:"printer_id"`
	Order     string `db:"check_order"`
	Status    string `db:"status"`
	FilePath  string `db:"file_path"`
	CheckType int    `db:"check_type"`
}

func (c OrderCheck) GetId() int {
	return c.Id
}

func (c *OrderCheck) SetId(id int) {
	c.Id = id
}

func (c OrderCheck) GetPrinterId() string {
	return c.PrinterId
}

func (c OrderCheck) GetOrder() string {
	return c.Order
}

func (c OrderCheck) GetStatus() string {
	return c.Status
}

func (c *OrderCheck) SetStatus(status string) {
	c.Status = status
}

func (c OrderCheck) GetFilePath() string {
	return c.FilePath
}

func (c OrderCheck) GetCheckType() int {
	return c.CheckType
}
