package gmodel

type ToSql interface {
	ToSql() (string, []any)
}

type Field interface {
	Name() string
	As(name string) ToSql
	ToSql
}
