package entity

type CheckBuilder struct {
	id        int
	printerId string
	order     string
	status    string
	filePath  string
	checkType int
}

func NewCheckBuilder() *CheckBuilder {
	return &CheckBuilder{}
}

func (b *CheckBuilder) SetId(id int) *CheckBuilder {
	b.id = id
	return b
}

func (b *CheckBuilder) SetPrinterId(printerId string) *CheckBuilder {
	b.printerId = printerId
	return b
}

func (b *CheckBuilder) SetOrder(order string) *CheckBuilder {
	b.order = order
	return b

}

func (b *CheckBuilder) SetStatus(status string) *CheckBuilder {
	b.status = status
	return b

}

func (b *CheckBuilder) SetFilePath(filePath string) *CheckBuilder {
	b.filePath = filePath
	return b
}

func (b *CheckBuilder) SetCheckType(checkType int) *CheckBuilder {
	b.checkType = checkType
	return b
}

func (b CheckBuilder) Build() OrderCheck {
	return OrderCheck{
		id:        b.id,
		printerId: b.printerId,
		order:     b.order,
		status:    b.status,
		filePath:  b.filePath,
		checkType: b.checkType,
	}
}
