package entity

type OrderCheck struct {
	id        int
	printerId string
	order     string
	status    string
	filePath  string
	checkType int
}

func (c OrderCheck) Id() int {
	return c.id
}

func (c *OrderCheck) SetId(id int) {
	c.id = id
}

func (c OrderCheck) PrinterId() string {
	return c.printerId
}

func (c OrderCheck) Order() string {
	return c.order
}

func (c OrderCheck) Status() string {
	return c.status
}

func (c *OrderCheck) SetStatus(status string) {
	c.status = status
}

func (c OrderCheck) FilePath() string {
	return c.filePath
}

func (c OrderCheck) CheckType() int {
	return c.checkType
}
