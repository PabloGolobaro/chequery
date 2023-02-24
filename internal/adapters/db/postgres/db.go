package postgres

import "github.com/pablogolobaro/chequery/pkg/psql"

type Storages struct {
}

func New(repository psql.Repository) *Storages {
	return &Storages{}
}
