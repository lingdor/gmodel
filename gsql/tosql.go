package gsql

type ToSql interface {
	ToSql() (string, []any)
}
