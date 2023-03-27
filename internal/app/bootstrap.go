package app

import (
	"github.com/pablogolobaro/chequery/internal/adapters/db/postgres"
	"github.com/pablogolobaro/chequery/internal/adapters/pdf"
	"github.com/pablogolobaro/chequery/internal/config"
	"github.com/pablogolobaro/chequery/internal/domain/services"
	"github.com/pablogolobaro/chequery/internal/domain/usecases"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/check"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/health"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/order"
	"github.com/pablogolobaro/chequery/internal/handlers/web"
	"github.com/pablogolobaro/chequery/pkg/htmltopdf"
	"github.com/pablogolobaro/chequery/pkg/psql"
	"github.com/pablogolobaro/chequery/pkg/renderer"
)

const templateDir = "./static/templates"
const exePath = "bin"

func (a *Application) Bootstrap(conf config.Config) error {
	repository, err := psql.New(conf.DSN())
	if err != nil {
		return err
	}

	postgresStorages := postgres.New(repository)

	a.renderer = renderer.New()

	err = a.renderer.LoadTemplates(templateDir)
	if err != nil {
		return err
	}

	err = htmltopdf.FindWKHTMLTOPDF(exePath)
	if err != nil {
		return err
	}

	pdfStorage := pdf.NewPdfStorage(a.renderer)

	printerService := services.NewPrinterService(a.log, postgresStorages.PrinterStorage)

	checkService := services.NewCheckService(a.log, postgresStorages.CheckStorage, pdfStorage)

	useCases := usecases.NewCheckUseCase(a.log, checkService, printerService)

	a.checkHandler = check.NewCheckHandler(a.log, useCases)

	a.orderHandler = order.NewOrderHandler(a.log, useCases)

	a.healthHandler = health.NewHealthCheckHandler()

	a.uiHandler = web.NewUiHandler(a.log, useCases)

	return nil
}
