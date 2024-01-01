package gmodel

import (
	"context"
	"database/sql"
)

// ExecuteContext some sql like update,insert with context
func ExecuteContext(ctx context.Context, handler DBHandler, toSql ToSql) (sql.Result, error) {
	sqlStr, pms := toSql.ToSql()
	return handler.ExecContext(ctx, sqlStr, pms)
}

// Execute some sql like update,insert
func Execute(handler DBHandler, toSql ToSql) (sql.Result, error) {
	return ExecuteContext(context.Background(), handler, toSql)
}
