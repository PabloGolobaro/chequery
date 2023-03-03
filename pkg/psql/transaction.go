package psql

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

type SqlQuery struct {
	sql  string
	args []interface{}
}

func NewSqlQuery() *SqlQuery {
	return &SqlQuery{}
}

func (s *SqlQuery) SetSql(sql string) *SqlQuery {
	s.sql = sql
	return s
}

func (s *SqlQuery) SetArgs(args []interface{}) *SqlQuery {
	s.args = args
	return s
}

func Transaction(dbClient *sqlx.DB, queries []SqlQuery) error {
	tx, err := dbClient.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, query := range queries {
		_, err = tx.Exec(query.sql, query.args)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("update drivers: unable to rollback: %v", rollbackErr)
			}
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
