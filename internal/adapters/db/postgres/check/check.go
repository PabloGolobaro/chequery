package check

import (
	"github.com/jmoiron/sqlx"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/pkg/psql"
)

type checkStorage struct {
	dbClient *sqlx.DB
}

func NewStorage(repository psql.Repository) *checkStorage {
	return &checkStorage{dbClient: repository.GetConnection()}
}

func (c checkStorage) Get(id string) entity.ICheck {
	//TODO implement me
	panic("implement me")
}

func (c checkStorage) GetAll() []entity.ICheck {
	//TODO implement me
	panic("implement me")
}

func (c checkStorage) Create(check entity.ICheck) error {
	//TODO implement me
	panic("implement me")
}
