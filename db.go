package gmodel

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

type DBHandler interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	Driver() driver.Driver
}
