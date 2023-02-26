package check

import (
	sq "github.com/Masterminds/squirrel"
)

const checkTable = "check"

func prepareGet(id int) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Select("*").From(checkTable).Where(sq.Eq{"id": id})

	return rawQuery.ToSql()
}

func prepareGetAllGenerated() (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Select("*").From(checkTable).Where(sq.Eq{"status": "generated"})

	return rawQuery.ToSql()
}

func prepareUpdateStatusPrinted(ids []int) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Update(checkTable).Set("status", "printed").Where(sq.Eq{"id": ids})

	return rawQuery.ToSql()
}
