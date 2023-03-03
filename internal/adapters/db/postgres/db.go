package postgres

import (
	"github.com/pablogolobaro/chequery/internal/adapters/db/postgres/check"
	"github.com/pablogolobaro/chequery/internal/adapters/db/postgres/printer"
	"github.com/pablogolobaro/chequery/internal/domain/services"
	"github.com/pablogolobaro/chequery/pkg/psql"
)

type Storages struct {
	CheckStorage   services.CheckStorage
	PrinterStorage services.PrinterStorage
}

func New(repository psql.Repository) *Storages {
	return &Storages{
		CheckStorage:   check.NewStorage(repository),
		PrinterStorage: printer.NewStorage(repository),
	}
}
