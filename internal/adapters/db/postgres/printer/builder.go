package printer

import (
	sq "github.com/Masterminds/squirrel"
)

const printerTable = "printers"

func prepareGetAll() (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Select("api_key,name, printer_type, point_id").From(printerTable)

	return rawQuery.ToSql()
}

func prepareGetByPoint(pointID int) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Select("api_key,name, printer_type, point_id").From(printerTable).Where(sq.Eq{"point_id": pointID})

	return rawQuery.ToSql()
}
