package gmodel

import (
	"bytes"
	"reflect"
)

type TypeSchemaFields []ToSql

func (t TypeSchemaFields) ToSql(config ToSqlConfig) (string, []any) {
	sqlBuf := &bytes.Buffer{}
	pms := make([]any, 0)
	for i, field := range t {
		sql, fieldPms := field.ToSql(config)
		if fieldPms != nil {
			pms = append(pms, fieldPms...)
		}
		if i > 0 {
			sqlBuf.WriteString(",")
		}
		sqlBuf.WriteString(sql)
	}
	return sqlBuf.String(), pms
}

func AllSchemaFields(schema any) ToSql {

	rval := reflect.ValueOf(schema)
	if rval.Kind() != reflect.Pointer {
		panic("assert faild AllSchemaFields type")
	}
	rval = rval.Elem()

	if rval.Kind() != reflect.Struct {
		panic("assert type faild, schema must be a struct")
	}

	var num = rval.NumField()
	var ret = make([]ToSql, 0, num)
	for i := 0; i < num; i++ {
		fieldval := rval.Field(i).Interface()
		if field, ok := fieldval.(ToSql); ok {
			ret = append(ret, field)
		}
	}
	return TypeSchemaFields(ret)
}
