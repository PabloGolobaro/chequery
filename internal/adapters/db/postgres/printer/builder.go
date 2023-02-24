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

	rawQuery := psqlSq.Select("*").From(printerTable).Where(sq.Eq{"point_id": pointID})

	return rawQuery.ToSql()
}

//rawQuery := psqlSq.Insert(metricsTable).Columns("name, value, date")
//for _, metric := range metrics {
//	rawQuery = rawQuery.Values(metric.Name, metric.Value, metric.Date)
//}

//func prepareFindMetricsByFilter(filter models.Filter) (string, []interface{}, error) {
//	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
//
//	rowQuery := psqlSq.Select("name, value, date").From(metricsTable)
//	conj := sq.And{}
//
//	if !filter.Time_from.IsZero() {
//		conj = append(conj, sq.Gt{"time": filter.Time_from})
//	}
//
//	if !filter.Time_to.IsZero() {
//		conj = append(conj, sq.Lt{"time": filter.Time_to})
//	}
//
//	rowQuery = rowQuery.Where(conj)
//
//	if filter.Offset > 0 {
//		rowQuery = rowQuery.Offset(uint64(filter.Offset))
//	}
//
//	if filter.Limit > 0 {
//		rowQuery = rowQuery.Limit(uint64(filter.Limit))
//	}
//
//	return rowQuery.ToSql()
//}
