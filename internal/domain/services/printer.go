package services

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

type PrinterStorage interface {
	Get(id string) (entity.Printer, error)
	GetByPoint(pointID int) ([]entity.Printer, error)
	GetAll() ([]entity.Printer, error)
	Create(printer entity.Printer) error
}

type printerService struct {
	printerStorage PrinterStorage
}

func NewPrinterService(printerStorage PrinterStorage) *printerService {
	return &printerService{printerStorage: printerStorage}
}

func (p printerService) GetPrinters(ctx context.Context) ([]entity.Printer, error) {
	return p.printerStorage.GetAll()
}

func (p printerService) GetPrintersByPoint(ctx context.Context, pointID int) ([]entity.Printer, error) {
	return p.printerStorage.GetByPoint(pointID)
}
