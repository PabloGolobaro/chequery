package services

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

func (c checkService) CreateGuestCheck(ctx context.Context, check entity.OrderCheck) error {
	errs := make(chan error)
	quit := make(chan struct{})
	go func() {
		select {
		case errs <- c.checkStorage.Create(check):
		case <-quit:
			return
		}
	}()

	select {
	case err := <-errs:

		go func() {
			err := c.pdfStorage.GenerateCheckPDF(check)
			if err != nil {
				return
			}
		}()

		return err
	case <-ctx.Done():
		close(quit)
		return ctx.Err()
	}
}

func (c checkService) CreateKitchenCheck(ctx context.Context, check entity.OrderCheck) error {
	errs := make(chan error)
	quit := make(chan struct{})
	go func() {
		select {
		case errs <- c.checkStorage.Create(check):
		case <-quit:
			return
		}
	}()

	select {
	case err := <-errs:
		go func() {
			err := c.pdfStorage.GenerateCheckPDF(check)
			if err != nil {
				return
			}
		}()

		return err
	case <-ctx.Done():
		close(quit)
		return ctx.Err()
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
			res <- generatedChecks.PdfFileName()
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
