package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	GetConnection() *sqlx.DB
}

type repository struct {
	Client *sqlx.DB
}

func New(dsn string) (Repository, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	repo := repository{Client: db}
	if err = repo.Client.Ping(); err != nil {
		return &repo, err
	}

	return &repo, nil
}

func (c repository) GetConnection() *sqlx.DB {
	return c.Client
}
