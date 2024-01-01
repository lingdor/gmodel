package gsql

import "bytes"

type selectSqlBuilder struct {
	selectSql string
	from      string
	where     string
	groupBy   string
	last      string
	paramers  []any
}

func (d *selectSqlBuilder) Select(sql ToSql) {
	d.selectSql, _ = sql.ToSql()
}
func (d *selectSqlBuilder) From(sql ToSql) {
	var pms []any
	d.from, pms = sql.ToSql()
	if pms != nil {
		d.paramers = append(d.paramers, pms...)
	}
}
func (d *selectSqlBuilder) Where(sql ToSql) {
	var pms []any
	d.where, pms = sql.ToSql()
	if pms != nil {
		d.paramers = append(d.paramers, pms...)
	}
}
func (d *selectSqlBuilder) GroupBy(sql ToSql) {
	d.groupBy, _ = sql.ToSql()
}
func (d *selectSqlBuilder) Last(sql ToSql) {
	var pms []any
	d.last, pms = sql.ToSql()
	if pms != nil {
		d.paramers = append(d.paramers, pms...)
	}
}

//	func (d *selectSqlBuilder) genDeletSql() string {
//		buf := bytes.Buffer{}
//		buf.Write([]byte("delete from "))
//		buf.Write([]byte(d.from))
//		if d.where != "" {
//			buf.Write([]byte(" where "))
//			buf.Write([]byte(d.where))
//		}
//		if d.last != "" {
//			buf.Write([]byte(" "))
//			buf.Write([]byte(d.last))
//		}
//		return buf.String()
//	}
func (d *selectSqlBuilder) ToSql() (string, any) {
	buf := bytes.Buffer{}
	buf.Write([]byte("select "))
	if d.selectSql == "" {
		d.selectSql = "*"
	}
	buf.Write([]byte(d.selectSql))
	buf.Write([]byte("from "))
	buf.Write([]byte(d.from))
	if d.where != "" {
		buf.Write([]byte(" where "))
		buf.Write([]byte(d.where))
	}
	if d.groupBy != "" {
		buf.Write([]byte(" "))
		buf.Write([]byte(d.groupBy))
	}
	if d.last != "" {
		buf.Write([]byte(" "))
		buf.Write([]byte(d.last))
	}
	return buf.String(), d.paramers
}

//func (d *selectSqlBuilder) genInsertSql() string {
//	buf := bytes.Buffer{}
//	buf.Write([]byte("insert into "))
//	buf.Write([]byte(d.from))
//	buf.Write([]byte(" "))
//	buf.Write([]byte(d.set))
//	if d.last != "" {
//		buf.Write([]byte(" "))
//		buf.Write([]byte(d.last))
//	}
//	return buf.String()
//}

type rawSql struct {
	sql string
	pms []any
}

func (r *rawSql) ToSql() (string, []any) {
	return r.sql, r.pms
}

func RawSql(sql string, pms ...any) ToSql {
	return &rawSql{
		sql: sql,
		pms: pms,
	}
}
