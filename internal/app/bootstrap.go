package app

import (
	"github.com/pablogolobaro/chequery/internal/adapters/db/postgres"
	"github.com/pablogolobaro/chequery/internal/adapters/pdf"
	"github.com/pablogolobaro/chequery/internal/config"
	"github.com/pablogolobaro/chequery/internal/domain/services"
	"github.com/pablogolobaro/chequery/internal/domain/usecases"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/check"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/order"
	"github.com/pablogolobaro/chequery/pkg/psql"
	"github.com/pablogolobaro/chequery/pkg/templ"
)

func (a *Application) Bootstrap(conf config.Config) error {
	repository, err := psql.New(conf.DSN())
	if err != nil {
		return err
	}

	postgresStorages := postgres.New(repository)

	template, err := templ.ParseTemplate(templateDir)
	if err != nil {
		return err
	}

	pdfStorage := pdf.NewPdfStorage(template)

	printerService := services.NewPrinterService(a.log, postgresStorages.PrinterStorage)

	checkService := services.NewCheckService(a.log, postgresStorages.CheckStorage, pdfStorage)

	useCases := usecases.NewCheckUseCase(a.log, checkService, printerService)

	a.checkHandler = check.NewCheckHandler(a.log, useCases)

	a.orderHandler = order.NewOrderHandler(a.log, useCases)

	return nil
}
