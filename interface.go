package gmodel

type ToSql interface {
	ToSql() (string, []any)
}

type Field interface {
	Name() string
	Type() string
	Size() int
	ToSql
}
