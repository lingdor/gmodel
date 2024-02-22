package orm

import (
	"bytes"
	"fmt"
	"github.com/lingdor/gmodel/common"
	"github.com/lingdor/magicarray/array"
)

type insertSqlBuilder struct {
	table     ToSql
	fields    []Field
	values    [][]any
	selectSql ToSql
	last      []ToSql
}

func (d *insertSqlBuilder) Values(vals ...any) *insertSqlBuilder {
	d.values = append(d.values, vals)
	return d
}

func (d *insertSqlBuilder) Set(field Field, val any) *insertSqlBuilder {

	if len(d.values) < 2 {
		d.fields = append(d.fields, field)
	}
	if len(d.values) == 0 {
		d.values = append(d.values, []any{val})
	}
	last := len(d.values) - 1
	d.values[last] = append(d.values[last], val)
	return d
}

func (d *insertSqlBuilder) SetArr(arr array.MagicArray) *insertSqlBuilder {
	iter := arr.Iter()
	for k, v := iter.NextKV(); k != nil; k, v = iter.NextKV() {
		fieldtag, _ := array.ZValTagGet(v, common.TagName)
		fieldname := ToDbName(k.String(), fieldtag)
		d.Set(NewNameField(fieldname), v.Interface())
	}
	return d
}

func (d *insertSqlBuilder) SetEntity(entity any) *insertSqlBuilder {
	return d.SetArr(array.ValueofStruct(entity))
}
func (d *insertSqlBuilder) SetMap(vals map[Field]any) *insertSqlBuilder {
	for k, v := range vals {
		d.Set(k, v)
	}
	return d
}

func (d *insertSqlBuilder) Last(sql ToSql) *insertSqlBuilder {
	d.last = append(d.last, sql)
	return d
}
func (d *insertSqlBuilder) Select(fields ...ToSql) (ret *selectSqlBuilder) {
	ret = &selectSqlBuilder{
		fields: fields,
	}
	d.selectSql = ret
	return
}

func Insert(table ToSql) *insertSqlBuilder {

	return &insertSqlBuilder{
		table:  table,
		values: make([][]any, 0),
		last:   make([]ToSql, 0),
	}
}

func (d *insertSqlBuilder) ToSql(config common.ToSqlConfig) (string, []any) {

	if d.table == nil {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'from'"))
	}
	if len(d.values) < 1 && d.selectSql == nil {
		panic(fmt.Errorf("select sql generate faild, no found parts of:' values '"))
	}

	buf := bytes.Buffer{}
	parameters := make([]any, 0, 10)
	buf.WriteString("insert into ")
	tname, _ := d.table.ToSql(config)
	buf.WriteString(fmt.Sprintf(tname))
	if len(d.fields) > 0 {
		buf.WriteString("(")
		for i, field := range d.fields {
			if i > 0 {
				buf.WriteString(",")
			}
			buf.WriteString(common.OnlySql(field, config))
		}
		buf.WriteString(")")
	}

	for i, values := range d.values {
		if len(values) < 1 {
			panic(fmt.Errorf("empty values in the insert sql"))
		}
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString("(")
		for j, values := range d.values {
			if j > 0 {
				buf.WriteString(",")
			}
			buf.WriteString("?")
			parameters = append(parameters, values)
		}
		buf.WriteString(")")
	}
	if d.selectSql != nil {
		sql, pms := d.selectSql.ToSql(config)
		parameters = append(parameters, pms...)
		buf.WriteString(" ")
		buf.WriteString(sql)
	}
	if d.last != nil {
		buf.WriteString(" ")
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
