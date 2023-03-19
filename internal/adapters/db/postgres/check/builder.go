package check

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
)

const checkTable = "checks"

func prepareCreate(check entity.OrderCheck) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Insert(checkTable).Columns("printer_id", "check_order", "status", "check_type").
		Values(check.GetPrinterId(), check.GetOrder(), check.GetStatus(), check.GetCheckType())

	return rawQuery.ToSql()
}

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

func prepareUpdateStatusGenerated(id int) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Update(checkTable).Set("status", "generated").Where(sq.Eq{"id": id})

	return rawQuery.ToSql()
}

func prepareUpdateFilePath(id int, filePath string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	rawQuery := psqlSq.Update(checkTable).Set("file_path", filePath).Where(sq.Eq{"id": id})

	return rawQuery.ToSql()
}
