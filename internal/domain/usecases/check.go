package usecases

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/check"
	"log"
)

type CheckService interface {
	CreateCheck(ctx context.Context, check entity.OrderCheck) (entity.OrderCheck, error)
	GetGeneratedChecks(ctx context.Context) ([]entity.OrderCheck, error)
	UpdateChecksStatus(ctx context.Context, checkIDs []int) error
	GetCheckFilePath(ctx context.Context, checkId int) (string, error)
	GeneratePDFFile(ctx context.Context, check entity.OrderCheck) error
}

type PrinterService interface {
	GetPrintersByPoint(ctx context.Context, pointID int) ([]entity.Printer, error)
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

func (c checkUseCase) CreateChecks(ctx context.Context, details entity.OrderDetails) error {
	printers, err := c.printerService.GetPrintersByPoint(ctx, details.PointId())
	if err != nil {
		return err
	}

	for _, printer := range printers {

		orderCheck := *(entity.NewOrderCheck().OrderDetails(details).Printer(printer))

		orderCheck, err = c.checkService.CreateCheck(ctx, orderCheck)
		if err != nil {
			return err
		}

		go func(orderCheck entity.OrderCheck) {
			err := c.checkService.GeneratePDFFile(context.Background(), orderCheck)
			if err != nil {
				log.Println("error generating pdf file: ", err)
				return
			}
		}(orderCheck)

	}

	return nil
}
