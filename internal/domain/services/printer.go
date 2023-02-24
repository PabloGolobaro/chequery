package services

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

type PrinterStorage interface {
	Get(id string) (entity.IPrinter, error)
	GetByPoint(pointID int) ([]entity.IPrinter, error)
	GetAll() ([]entity.IPrinter, error)
	Create(printer entity.IPrinter) error
}

type printerService struct {
	printerStorage PrinterStorage
}

func NewPrinterService(printerStorage PrinterStorage) *printerService {
	return &printerService{printerStorage: printerStorage}
}

func (p printerService) GetPrinters(ctx context.Context) ([]entity.IPrinter, error) {
	return p.printerStorage.GetAll()
}

func (p printerService) GetPrintersByPoint(ctx context.Context, pointID int) ([]entity.IPrinter, error) {
	return p.printerStorage.GetByPoint(pointID)
}
