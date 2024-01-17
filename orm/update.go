package orm

import (
	"bytes"
	"fmt"
	"github.com/lingdor/gmodel/common"
	"github.com/lingdor/magicarray/array"
)

type SetInfo struct {
	k, v any
}

func (s *SetInfo) ToSql(config common.ToSqlConfig) (string, []any) {

	var fieldName string
	switch v := s.k.(type) {
	case Field:
		fieldName = v.Name()
	case ToSql:
		fieldName, _ = v.ToSql(config)
	case string:
		fieldName = v
	case array.ZVal:
		fieldName = v.String()
	default:
		panic(fmt.Errorf("no found type:%t", s.k))
	}
	return fmt.Sprintf("\"%s\" = ?", fieldName), []any{s.v}
}

type updateSqlBuilder struct {
	table ToSql
	where ToSql
	set   map[Field]any
	last  []ToSql
}
type FieldMap map[Field]any

func (f FieldMap) ToSql(config common.ToSqlConfig) (string, []any) {
	panic("Field Map is invoked")
}

func (d *updateSqlBuilder) Where(sql ToSql) {
	d.where = sql
}

func (d *updateSqlBuilder) Set(field Field, val any) {
	d.set[field] = val
}

func (d *updateSqlBuilder) SetArr(arr array.MagicArray) {
	iter := arr.Iter()
	for k, v := iter.NextKV(); k != nil; k, v = iter.NextKV() {
		d.Set(NewNameField(k.String()), v.Interface())
	}
}
func (d *updateSqlBuilder) SetMap(vals map[Field]any) {
	for k, v := range vals {
		d.Set(k, v)
	}
}

func (d *updateSqlBuilder) Last(sql ToSql) {
	d.last = append(d.last, sql)
}

func Update(table ToSql) *updateSqlBuilder {
	return &updateSqlBuilder{
		table: table,
		set:   make(map[Field]any, 0),
		last:  make([]ToSql, 0),
	}
}

func (d *updateSqlBuilder) ToSql(config common.ToSqlConfig) (string, any) {

	if d.table == nil {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'from'"))
	}
	if d.set == nil {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'set'"))
	}

	buf := bytes.Buffer{}
	parameters := make([]any, 0, 10)
	buf.WriteString("update ")
	sqlStr, _ := d.table.ToSql(config)
	buf.WriteString(sqlStr)

	buf.WriteString(" set ")

	first := true
	for field, fieldv := range d.set {
		if !first {
			buf.Write([]byte{byte(',')})
		} else {
			first = true
		}
		sqlStr := fmt.Sprintf("\"%s\"=?", field.Name())
		parameters = append(parameters, fieldv)
		buf.WriteString(sqlStr)
	}

	if d.where != nil {
		buf.WriteString(" where ")
		sqlStr, pms := d.where.ToSql(config)
		if pms != nil {
			parameters = append(parameters, pms...)
		}
		buf.WriteString(sqlStr)
	}
	if d.last != nil {
		buf.Write([]byte{byte(' ')})
		for _, sql := range d.last {
			sqlStr, pms := sql.ToSql(config)
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.WriteString(sqlStr)
		}
	}
	return buf.String(), parameters
}
