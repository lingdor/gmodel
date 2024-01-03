package orm

import (
	"bytes"
	"fmt"
	"github.com/lingdor/magicarray/array"
)

type SetInfo struct {
	k, v any
}

func (s *SetInfo) ToSql() (string, []any) {

	var fieldName string
	switch v := s.k.(type) {
	case Field:
		fieldName = v.Name()
	case ToSql:
		fieldName, _ = v.ToSql()
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
	table string
	where ToSql
	set   map[string]any
	last  []ToSql
}
type FieldMap map[Field]any

func (f FieldMap) ToSql() (string, []any) {
	panic("Field Map is invoked")
}

func (d *updateSqlBuilder) Where(sql ToSql) {
	d.where = sql
}

func (d *updateSqlBuilder) Set(sql ToSql) {
	d.set = append(d.set, sql)
}

func (d *updateSqlBuilder) SetKV(k, v any) {
	d.Set(&SetInfo{k: k, v: v})
}
func (d *updateSqlBuilder) SetArr(arr array.MagicArray) {
	iter := arr.Iter()
	for k, v := iter.NextKV(); k != nil; k, v = iter.NextKV() {
		d.SetKV(k, v.Interface())
	}
}
func (d *updateSqlBuilder) SetMap(vals map[Field]any) {
	for k, v := range vals {
		d.SetKV(k, v)
	}
}

func (d *updateSqlBuilder) Last(sql ToSql) {
	d.last = append(d.last, sql)
}

func Update(table ToSql) *updateSqlBuilder {
	return &updateSqlBuilder{
		table: table,
		set:   make([]ToSql, 0),
		last:  make([]ToSql, 0),
	}
}

func (d *updateSqlBuilder) ToSql() (string, any) {

	if d.table == nil {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'from'"))
	}
	if d.set == nil {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'set'"))
	}

	buf := bytes.Buffer{}
	parameters := make([]any, 0, 10)
	buf.Write([]byte("update "))
	sqlStr, _ := d.table.ToSql()
	buf.Write([]byte(sqlStr))

	buf.Write([]byte(" set "))

	first := true
	for _, sql := range d.set {
		if !first {
			buf.Write([]byte(byte(',')))
		}
		sqlStr, pms := sql.ToSql()
		if pms != nil {
			parameters = append(parameters, pms...)
		}
		buf.Write([]byte(sqlStr))
		if first && len(sqlStr) > 0 {
			first = false
		}
	}

	if d.where != nil {
		buf.Write([]byte(" where "))
		sqlStr, pms := d.where.ToSql()
		if pms != nil {
			parameters = append(parameters, pms...)
		}
		buf.Write([]byte(sqlStr))
	}
	if d.last != nil {
		buf.Write([]byte(byte(' ')))
		for _, sql := range d.last {
			sqlStr, pms := sql.ToSql()
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.Write([]byte(sqlStr))
		}
	}
	return buf.String(), parameters
}
