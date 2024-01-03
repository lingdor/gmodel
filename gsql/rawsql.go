package gsql

var CountAll = Raw("count(*)", nil)
var Null = Raw("NULL", nil)

type rawSql struct {
	sql string
	pms []any
}

func (r *rawSql) ToSql() (string, []any) {
	return r.sql, r.pms
}

func Raw(sql string, pms ...any) ToSql {
	return &rawSql{
		sql: sql,
		pms: pms,
	}
}
