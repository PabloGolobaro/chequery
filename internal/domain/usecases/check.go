package usecases

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/check"
)

type CheckService interface {
	CreateGuestCheck(ctx context.Context) error
	CreateKitchenCheck(ctx context.Context) error
	GetGeneratedChecks(ctx context.Context) ([]entity.ICheck, error)
	UpdateChecksStatus(ctx context.Context, checkIDs []string) error
	GetCheckFilePath(ctx context.Context, checkId string) (string, error)
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

func (c checkUseCase) GetGeneratedCheckIDs(ctx context.Context) (check.IDs, error) {
	var ids = check.IDs{}

	generatedChecks, err := c.checkService.GetGeneratedChecks(ctx)
	if err != nil {
		return ids, err
	}

	for _, generatedCheck := range generatedChecks {
		ids = append(ids, generatedCheck.Id())
	}

	return ids, nil

}

func (c checkUseCase) GetCheckFilePath(ctx context.Context, checkID string) (string, error) {
	return c.checkService.GetCheckFilePath(ctx, checkID)
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
