package check

import (
	"github.com/jmoiron/sqlx"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/internal/domain/services"
	"github.com/pablogolobaro/chequery/pkg/psql"
)

type storage struct {
	dbClient *sqlx.DB
}

func NewStorage(repository psql.Repository) services.CheckStorage {
	return &storage{dbClient: repository.GetConnection()}
}

func (s storage) UpdateStatusGeneratedAndFilePath(checkId int, filePath string) error {
	sqlQuery := []psql.SqlQuery{}

	sql, args, err := prepareUpdateStatusGenerated(checkId)
	if err != nil {
		return err
	}

	sqlQuery = append(sqlQuery, *(psql.NewSqlQuery().SetSql(sql).SetArgs(args)))

	sql, args, err = prepareUpdateFilePath(checkId, filePath)
	if err != nil {
		return err
	}

	sqlQuery = append(sqlQuery, *(psql.NewSqlQuery().SetSql(sql).SetArgs(args)))

	err = psql.Transaction(s.dbClient, sqlQuery)
	if err != nil {
		return err
	}

	return nil
}

func (s storage) Get(id int) (entity.OrderCheck, error) {
	var check = entity.OrderCheck{}
	sql, args, err := prepareGet(id)
	if err != nil {
		return check, err
	}

	err = s.dbClient.Select(&check, sql, args)
	if err != nil {
		return check, err
	}
	return check, err

}

func (s storage) GetAll() []entity.OrderCheck {
	return nil
}

func (s storage) Create(check entity.OrderCheck) (int, error) {
	sql, args, err := prepareCreate(check)
	if err != nil {
		return 0, nil
	}

	result, err := s.dbClient.Exec(sql, args)
	if err != nil {
		return 0, nil
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}

func (s storage) GetAllGeneratedChecks() ([]entity.OrderCheck, error) {
	var checks = []entity.OrderCheck{}
	sql, args, err := prepareGetAllGenerated()
	if err != nil {
		return nil, err
	}

	err = s.dbClient.Select(&checks, sql, args)
	if err != nil {
		return nil, err
	}
	return checks, err
}

func (s storage) UpdateStatusPrinted(checkIds []int) error {
	sql, args, err := prepareUpdateStatusPrinted(checkIds)
	if err != nil {
		return err
	}

	_, err = s.dbClient.Exec(sql, args)
	if err != nil {
		return err
	}
	return nil
}
