package gsql

import (
	"bytes"
	"fmt"
	"github.com/lingdor/gmodel/common"
	"github.com/lingdor/magicarray/array"
)

type SetInfo struct {
	k, v any
}

type updateSqlBuilder struct {
	table string
	where ToSql
	joins []ToSql
	set   map[string]any
	last  []ToSql
}

func (d *updateSqlBuilder) Where(sql ToSql) {
	d.where = sql
}

func (d *updateSqlBuilder) Set(field string, val any) {
	d.set[field] = val
}
func (d *updateSqlBuilder) SetArr(arr array.MagicArray) {
	iter := arr.Iter()
	for k, v := iter.NextKV(); k != nil; k, v = iter.NextKV() {
		fieldtag, _ := array.ZValTagGet(v, common.TagName)
		fieldname := ToDbName(k.String(), fieldtag)
		d.set[fieldname] = v.Interface()
	}
}
func (d *updateSqlBuilder) SetMap(vals map[string]any) {
	for k, v := range vals {
		d.set[k] = v
	}
}

func (d *updateSqlBuilder) Last(sql ToSql) {
	d.last = append(d.last, sql)
}

func Update(table string) *updateSqlBuilder {
	return &updateSqlBuilder{
		table: table,
		set:   make(map[string]any, 0),
		last:  make([]ToSql, 0),
	}
}

func (d *updateSqlBuilder) ToSql(config common.ToSqlConfig) (string, any) {

	if d.table == "" {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'from'"))
	}
	if d.set == nil {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'set'"))
	}

	buf := bytes.Buffer{}
	parameters := make([]any, 0, 10)
	buf.WriteString("update ")
	buf.WriteString(d.table)

	buf.WriteString(" set ")

	first := true
	for k, v := range d.set {
		if !first {
			buf.Write([]byte{byte(',')})
		} else {
			first = false
		}
		buf.WriteString(fmt.Sprintf("\"%s\"=?", k))
		parameters = append(parameters, v)

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
