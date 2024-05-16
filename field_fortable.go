package gmodel

import (
	"fmt"
	"github.com/lingdor/gmodel/common"
	"github.com/lingdor/gmodel/orm"
)

type FieldForAlias struct {
	alias string
	field Field
}

func (f *FieldForAlias) Name() string {
	return f.Name()
}

func (f *FieldForAlias) As(name string) common.ToSql {
	return orm.WrapField(f).As(name)
}

func (f *FieldForAlias) ToSql(config common.ToSqlConfig) (string, []any) {
	return fmt.Sprintf("%s.%s", config.TableFormat(f.alias), config.FieldFormat(f.field.Name())), nil
}

func GetFieldForAlias(field Field, alias string) Field {

	return &FieldForAlias{
		alias: alias,
		field: field,
	}
}
