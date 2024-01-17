package common

func OnlySql(toSql ToSql, config ToSqlConfig) string {

	sql, _ := toSql.ToSql(config)
	return sql
}
