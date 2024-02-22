package orm

import (
	"fmt"
	"github.com/lingdor/gmodel/common"
)

type nameField string

func (n nameField) ToSql(config common.ToSqlConfig) (string, []any) {
	return config.FieldFormat(string(n)), nil
}
func (n nameField) Name() string {
	return string(n)
}
func (n nameField) As(name string) ToSql {
	return WrapField(n).As(name)
}

func NewNameField(field string) Field {
	return nameField(field)
}

var Raw = common.Raw
var Sql = common.Sql

type fieldWrapper struct {
	field ToSql
	alias string
}

func (f *fieldWrapper) ToSql(config common.ToSqlConfig) (string, []any) {
	sql, pms := f.field.ToSql(config)
	return fmt.Sprintf("%s as %s", sql, config.FieldFormat(f.alias)), pms
}
func (f *fieldWrapper) As(alias string) *fieldWrapper {
	f.alias = alias
	return f
}
func WrapField(field ToSql) *fieldWrapper {
	return &fieldWrapper{
		field: field,
	}
}
