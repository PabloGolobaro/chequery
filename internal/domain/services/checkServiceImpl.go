package services

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

func (c checkService) GeneratePDFFile(ctx context.Context, check entity.OrderCheck) error {

	filePath, err := c.pdfStorage.GenerateCheckPDF(check)
	if err != nil {
		return err
	}

	err = c.checkStorage.UpdateStatusGeneratedAndFilePath(check.Id(), filePath)
	if err != nil {
		return err
	}

	return nil
}

func (c checkService) CreateCheck(ctx context.Context, check entity.OrderCheck) (entity.OrderCheck, error) {
	type resultStruct struct {
		createdId int
		err       error
	}

	res := make(chan resultStruct)
	quit := make(chan struct{})
	go func() {
		select {
		case <-quit:
			return
		default:
			create, err := c.checkStorage.Create(check)
			res <- resultStruct{createdId: create, err: err}
		}
	}()

	select {
	case result := <-res:
		check.SetId(result.createdId)
		return check, result.err

	case <-ctx.Done():
		close(quit)
		return check, ctx.Err()
	}
}

func (c checkService) GetGeneratedChecks(ctx context.Context) ([]entity.OrderCheck, error) {
	errs := make(chan error)
	quit := make(chan struct{})
	res := make(chan []entity.OrderCheck)

	go func() {
		select {
		case <-quit:
			return
		default:
			generatedChecks, err := c.checkStorage.GetAllGeneratedChecks()
			if err != nil {
				errs <- err
			}
			res <- generatedChecks
		}
	}()

	select {
	case err := <-errs:
		return nil, err
	case result := <-res:
		return result, nil
	case <-ctx.Done():
		close(quit)
		return nil, ctx.Err()
	}
}

func (c checkService) UpdateChecksStatus(ctx context.Context, checkIDs []int) error {
	errs := make(chan error)
	quit := make(chan struct{})
	go func() {
		select {
		case errs <- c.checkStorage.UpdateStatusPrinted(checkIDs):
		case <-quit:
			return
		}
	}()

	select {
	case err := <-errs:
		return err
	case <-ctx.Done():
		close(quit)
		return ctx.Err()
	}
}

func (c checkService) GetCheckFilePath(ctx context.Context, checkId int) (string, error) {
	errs := make(chan error)
	quit := make(chan struct{})
	res := make(chan string)

	go func() {
		select {
		case <-quit:
			return
		default:
			generatedChecks, err := c.checkStorage.Get(checkId)
			if err != nil {
				errs <- err
			}
			res <- generatedChecks.FilePath()
		}
	}()

	select {
	case err := <-errs:
		return "", err
	case result := <-res:
		return result, nil
	case <-ctx.Done():
		close(quit)
		return "", ctx.Err()
	}
}
