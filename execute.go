package gmodel

import (
	"context"
	"database/sql"
)

// ExecuteContext some sql like update,insert with context
func ExecuteContext(ctx context.Context, db DBHandler, toSql ToSql) (sql.Result, error) {
	tosqlConfig := GetToSqlConfig(db)
	sqlStr, ps := toSqlCall(ctx, toSql, tosqlConfig)
	return db.ExecContext(ctx, sqlStr, ps)
}

// Execute some sql like update,insert
func Execute(handler DBHandler, toSql ToSql) (sql.Result, error) {
	return ExecuteContext(context.Background(), handler, toSql)
}
