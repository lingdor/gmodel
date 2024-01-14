package common

type ToSql interface {
	ToSql() (string, []any)
}
type Field interface {
	Name() string
	ToSql
}
