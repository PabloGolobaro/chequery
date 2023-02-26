package postgres

import (
	"github.com/pablogolobaro/chequery/internal/adapters/db/postgres/check"
	"github.com/pablogolobaro/chequery/internal/domain/services"
	"github.com/pablogolobaro/chequery/pkg/psql"
)

type Storages struct {
	checkStorage services.CheckStorage
}

func New(repository psql.Repository) *Storages {
	return &Storages{
		checkStorage: check.NewStorage(repository),
	}
}
