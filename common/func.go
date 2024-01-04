package common

func OnlySql(toSql ToSql) string {

	sql, _ := toSql.ToSql()
	return sql
}
