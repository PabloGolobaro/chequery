package printer

import (
	"github.com/jmoiron/sqlx"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/pkg/psql"
)

type printerStorage struct {
	dbClient *sqlx.DB
}

func NewStorage(repository psql.Repository) *printerStorage {
	return &printerStorage{dbClient: repository.GetConnection()}
}

func (p printerStorage) Get(id string) (entity.Printer, error) {
	return entity.Printer{}, nil
}

func (p printerStorage) GetAll() (printers []entity.Printer, err error) {
	sql, args, err := prepareGetAll()
	if err != nil {
		return nil, err
	}

	err = p.dbClient.Select(&printers, sql, args)
	if err != nil {
		return nil, err
	}

	return
}

func (p printerStorage) GetByPoint(pointID int) (printers []entity.Printer, err error) {
	sql, args, err := prepareGetByPoint(pointID)
	if err != nil {
		return nil, err
	}

	err = p.dbClient.Select(&printers, sql, args)
	if err != nil {
		return nil, err
	}

	return
}

func (p printerStorage) Create(printer entity.Printer) error {
	return nil
}
