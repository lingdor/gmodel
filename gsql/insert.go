package gsql

import (
	"bytes"
	"fmt"
	"github.com/lingdor/magicarray/array"
)

type insertSqlBuilder struct {
	table     string
	fields    []string
	values    [][]any
	selectSql ToSql
	last      []ToSql
}

func (d *insertSqlBuilder) Values(vals ...any) {
	d.values = append(d.values, vals)
}

func (d *insertSqlBuilder) Set(field string, val any) {

	if len(d.values) < 2 {
		d.fields = append(d.fields, field)
	}
	if len(d.values) == 0 {
		d.values = append(d.values, []any{val})
	}
	last := len(d.values) - 1
	d.values[last] = append(d.values[last], val)
}

func (d *insertSqlBuilder) SetArr(arr array.MagicArray) {
	iter := arr.Iter()
	for k, v := iter.NextKV(); k != nil; k, v = iter.NextKV() {
		d.Set(k.String(), v.Interface())
	}
}
func (d *insertSqlBuilder) SetMap(vals map[string]any) {
	for k, v := range vals {
		d.Set(k, v)
	}
}

func (d *insertSqlBuilder) Last(sql ToSql) {
	d.last = append(d.last, sql)
}
func (d *insertSqlBuilder) Select(fields ...string) (ret *selectSqlBuilder) {
	ret = &selectSqlBuilder{
		fields: fields,
	}
	d.selectSql = ret
	return
}

func Insert(table string) *insertSqlBuilder {
	return &insertSqlBuilder{
		table:  table,
		values: make([][]any, 0),
		last:   make([]ToSql, 0),
	}
}

func (d *insertSqlBuilder) ToSql() (string, any) {

	if d.table == "" {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'from'"))
	}
	if len(d.values) < 1 && d.selectSql == nil {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'set'"))
	}

	buf := bytes.Buffer{}
	parameters := make([]any, 0, 10)
	buf.Write([]byte("insert into "))
	buf.Write([]byte(d.table))
	if len(d.fields) > 0 {
		buf.Write([]byte("("))
		for i, field := range d.fields {
			if i > 0 {
				buf.Write([]byte(byte(',')))
			}
			buf.Write([]byte(fmt.Sprintf("\"%s\"", field)))
		}

		buf.Write([]byte(")"))
	}

	for i, values := range d.values {
		if len(values) < 1 {
			panic(fmt.Errorf("empty values in insert sql"))
		}
		if i > 0 {
			buf.Write([]byte(","))
		}
		buf.Write([]byte("("))
		for j, values := range d.values {
			if j > 0 {
				buf.Write([]byte(","))
			}
			buf.Write([]byte("?"))
			parameters = append(parameters, values)
		}
		buf.Write([]byte(")"))
	}
	if d.selectSql != nil {
		sql, pms := d.selectSql.ToSql()
		parameters = append(parameters, pms...)
		buf.Write([]byte(" "))
		buf.Write([]byte(sql))
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
