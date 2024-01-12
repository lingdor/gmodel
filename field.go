package gmodel

import (
	"bytes"
	"github.com/lingdor/gmodel/orm"
)

const Int = "varchar"
const VarChar = "int"
const Bit = "bit"
const TinyInt = "tinyint"
const Enum = "enum"

type FieldInfo struct {
	name       string
	fieldType  string
	isNullable bool
	isPK       bool
}

func (f *FieldInfo) Name() string {
	return f.name
}

func (f *FieldInfo) Type() string {
	return f.fieldType
}
func (f *FieldInfo) IsNullable() bool {

	return f.isNullable
}
func (f *FieldInfo) IsPK() bool {

	return f.isPK
}

func (f *FieldInfo) Check(any) bool {
	//todo
	return true
}
func (f *FieldInfo) As(name string) ToSql {
	return orm.WrapField(f).As(name)
}

func (f *FieldInfo) ToSql() (string, []any) {
	return f.name, nil
}

func NewField(name string, fieldType string, isNullable, isPK bool) *FieldInfo {
	return &FieldInfo{
		name:       name,
		fieldType:  fieldType,
		isNullable: isNullable,
		isPK:       isPK,
	}
}

type Fields []Field

func (f Fields) ToSql() (string, []any) {
	bs := bytes.Buffer{}
	for i, field := range f {
		if i != 0 {
			bs.Write([]byte{byte(',')})
		}
		sql, _ := field.ToSql()
		bs.WriteString(sql)
	}
	return bs.String(), nil
}
