package usecases

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/check"
)

type CheckService interface {
	CreateKitchenCheck(ctx context.Context, check entity.OrderCheck) error
	CreateGuestCheck(ctx context.Context, check entity.OrderCheck) error
	GetGeneratedChecks(ctx context.Context) ([]entity.OrderCheck, error)
	UpdateChecksStatus(ctx context.Context, checkIDs []int) error
	GetCheckFilePath(ctx context.Context, checkId int) (string, error)
}

type PrinterService interface {
	GetPrintersByPoint(ctx context.Context, pointID int) ([]entity.IPrinter, error)
}

type checkUseCase struct {
	checkService   CheckService
	printerService PrinterService
}

func NewCheckUseCase(checkService CheckService, printerService PrinterService) *checkUseCase {
	return &checkUseCase{checkService: checkService, printerService: printerService}
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

func (c checkUseCase) GetCheckFilePath(ctx context.Context, checkID int) (string, error) {
	return c.checkService.GetCheckFilePath(ctx, checkID)
}

func (c checkUseCase) SetChecksStatusPrinted(ctx context.Context, checkIDs []int) error {
	return c.checkService.UpdateChecksStatus(ctx, checkIDs)
}

func (c checkUseCase) CreateChecks(ctx context.Context, order string) error {
	printers, err := c.printerService.GetPrintersByPoint(ctx, 1)
	if err != nil {
		return err
	}
	for _, printer := range printers {
		if printer.Type() == entity.Kitchen {
			err := c.checkService.CreateKitchenCheck(ctx, entity.OrderCheck{})
			if err != nil {
				return err
			}
		} else if printer.Type() == entity.Guest {
			err := c.checkService.CreateGuestCheck(ctx, entity.OrderCheck{})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
