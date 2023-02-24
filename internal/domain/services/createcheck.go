package services

import (
	"context"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

func (c checkService) CreateGuestCheck(ctx context.Context) error {
	errs := make(chan error)
	quit := make(chan struct{})
	go func() {
		select {
		case errs <- c.checkStorage.Create(&entity.GuestCheck{}):
		case <-quit:
			return
		}
	}()

	select {
	case err := <-errs:

		go func() {
			err := c.pdfStorage.GenerateCheckPDF(&entity.GuestCheck{})
			if err != nil {
				return
			}
		}()

		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c checkService) CreateKitchenCheck(ctx context.Context) error {
	errs := make(chan error)
	quit := make(chan struct{})
	go func() {
		select {
		case errs <- c.checkStorage.Create(&entity.KitchenCheck{}):
		case <-quit:
			return
		}
	}()

	select {
	case err := <-errs:
		go func() {
			err := c.pdfStorage.GenerateCheckPDF(&entity.KitchenCheck{})
			if err != nil {
				return
			}
		}()

		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
