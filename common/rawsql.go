package common

var CountAll = Raw("count(*)", nil)
var Null = Raw("NULL", nil)

type rawSql struct {
	sql string
	pms []any
}

func (r *rawSql) ToSql(config ToSqlConfig) (string, []any) {
	return r.sql, r.pms
}

func Raw(sql string, pms ...any) ToSql {
	return &rawSql{
		sql: sql,
		pms: pms,
	}
}

type SqlStr string

func (s SqlStr) ToSql(config ToSqlConfig) (string, []any) {
	return string(s), nil
}

func Sql(sql string) ToSql {
	return SqlStr(sql)
}
