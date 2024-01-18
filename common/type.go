package common

type ToSqlConfig interface {
	FieldFormat(string) string
	TableFormat(string) string
}

type ToSql interface {
	ToSql(ToSqlConfig) (string, []any)
}
type Field interface {
	Name() string
	As(name string) ToSql
	ToSql
}

type EntityFieldHandler interface {
	GetFieldsHandler(fields []string) ([]any, bool)
}

type DefaultConfigLoader struct{}

func (m DefaultConfigLoader) FieldFormat(name string) string {
	return name
}

func (m DefaultConfigLoader) TableFormat(name string) string {
	return name
}

var All = Sql("*")
