package orm

import (
	"fmt"
	"github.com/lingdor/gmodel/common"
)

type nameField string

func (n nameField) ToSql() (string, []any) {
	return string(n), nil
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

func (f *fieldWrapper) ToSql() (string, []any) {
	sql, _ := f.field.ToSql()
	return fmt.Sprintf("%s as \"%s\"", sql, f.alias), nil
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
