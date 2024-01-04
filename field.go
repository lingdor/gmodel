package gmodel

import "bytes"

const Int = "varchar"
const VarChar = "int"
const Bit = "bit"
const TinyInt = "tinyint"
const Enum = "enum"

type FieldInfo struct {
	name       string
	fieldType  string
	size       int
	isNullable bool
}

func (f *FieldInfo) Name() string {
	return f.name
}

func (f *FieldInfo) Type() string {
	return f.fieldType
}

func (f *FieldInfo) Size() int {
	return f.size
}

func (f *FieldInfo) Check(any) bool {
	//todo
	return true
}

func (f *FieldInfo) ToSql() (string, []any) {
	return f.name, nil
}

func NewField(name string, fieldType string, size int) Field {
	return &FieldInfo{
		name:      name,
		fieldType: fieldType,
		size:      size,
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
