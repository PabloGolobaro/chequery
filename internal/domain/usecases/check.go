package usecases

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/check"
	"go.uber.org/zap"
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
	log            *zap.SugaredLogger
	checkService   CheckService
	printerService PrinterService
}

func NewCheckUseCase(log *zap.SugaredLogger, checkService CheckService, printerService PrinterService) *checkUseCase {
	return &checkUseCase{log: log, checkService: checkService, printerService: printerService}
}

func (c checkUseCase) GetGeneratedCheckIDs(ctx context.Context) (check.GeneratedChecksResponse, error) {
	var ids []int

	generatedChecks, err := c.checkService.GetGeneratedChecks(ctx)
	if err != nil {
		c.log.Errorw("Check.Usecases.CreateChecks.checkService.GetGeneratedChecks", "error: ", err)
		return check.GeneratedChecksResponse{IDs: ids}, err
	}

	for _, generatedCheck := range generatedChecks {
		ids = append(ids, generatedCheck.GetId())
	}

	return check.GeneratedChecksResponse{IDs: ids}, nil

}

func (c checkUseCase) GetCheckFilePath(ctx context.Context, checkID int) (string, error) {
	return c.checkService.GetCheckFilePath(ctx, checkID)
}

func (c checkUseCase) SetChecksStatusPrinted(ctx context.Context, checkIDs []int) error {
	return c.checkService.UpdateChecksStatus(ctx, checkIDs)
}

func (c checkUseCase) CreateChecks(ctx context.Context, order entity.OrderDetails) (ids []int, err error) {
	printers, err := c.printerService.GetPrintersByPoint(ctx, order.PointId())
	if err != nil {
		c.log.Errorw("Order.Usecases.CreateChecks.printerService.GetPrintersByPoint", "error: ", err, "order", order)
		return nil, err
	}

	for _, printer := range printers {

		orderCheck := entity.NewCheckBuilder().SetPrinterId(printer.GetApiKey()).SetCheckType(printer.GetPrinterType()).SetOrder(order.Details()).Build()

		orderCheck, err = c.checkService.CreateCheck(ctx, orderCheck)
		if err != nil {
			c.log.Errorw("Order.Usecases.CreateChecks.checkService.CreateCheck", "error: ", err, "orderCheck", orderCheck)
			return nil, err
		}

		go func(orderCheck entity.OrderCheck) {
			err := c.checkService.GeneratePDFFile(context.Background(), orderCheck)
			if err != nil {
				c.log.Errorw("Order.Usecases.checkService.GeneratePDFFile", "error: ", err, "orderCheck", orderCheck)

				return
			}
		}(orderCheck)

		ids = append(ids, orderCheck.GetId())
	}

	return
}
