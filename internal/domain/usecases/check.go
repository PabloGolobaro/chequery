package usecases

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

type CheckService interface {
	CreateGuestCheck(ctx context.Context) error
	CreateKitchenCheck(ctx context.Context) error
	GetChecksByPrinterID(ctx context.Context, printerID string) ([]entity.ICheck, error)
	UpdateChecksStatus(ctx context.Context, checkIDs []string) error
}

type PrinterService interface {
	GetPrinters(ctx context.Context) ([]entity.IPrinter, error)
	GetPrintersByPoint(ctx context.Context, pointID int) ([]entity.IPrinter, error)
}

type checkUseCase struct {
	checkService   CheckService
	printerService PrinterService
}

func NewCheckUseCase(checkService CheckService) *checkUseCase {
	return &checkUseCase{checkService: checkService}
}

func (c checkUseCase) SetChecksStatusPrinted(ctx context.Context, checkID []string) error {
	return c.checkService.UpdateChecksStatus(ctx, checkID)
}

func (c checkUseCase) CreateChecks(ctx context.Context, order string) error {
	printers, err := c.printerService.GetPrintersByPoint(ctx, 1)
	if err != nil {
		return err
	}
	for _, printer := range printers {
		if printer.Type() == entity.Kitchen {
			err := c.checkService.CreateKitchenCheck(ctx)
			if err != nil {
				return err
			}
		} else if printer.Type() == entity.Guest {
			err := c.checkService.CreateGuestCheck(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c checkUseCase) GetChecks(ctx context.Context, printerID string) ([]string, error) {
	checks, err := c.checkService.GetChecksByPrinterID(ctx, printerID)
	if err != nil {
		return nil, err
	}

	filenames := make([]string, 0, len(checks))

	for _, check := range checks {
		filenames = append(filenames, check.FileName())
	}

	return filenames, err
}
