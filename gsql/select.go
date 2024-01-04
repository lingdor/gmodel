package gsql

import (
	"bytes"
	"fmt"
	"strings"
)

type selectSqlBuilder struct {
	fields  []string
	from    string
	joins   []ToSql
	where   ToSql
	orderBy string
	groupBy string
	last    []ToSql
}
type selectJoinSqlBuilder struct {
	selectBuilder *selectSqlBuilder
	table         string
	joinType      string
	on            ToSql
}

func (s *selectJoinSqlBuilder) On(sql ToSql) *selectSqlBuilder {
	s.selectBuilder.joins = append(s.selectBuilder.joins, sql)
	return s.selectBuilder
}
func (s *selectJoinSqlBuilder) ToSql() (string, []any) {
	onStr, pms2 := s.on.ToSql()
	return fmt.Sprintf("%s %s on %s", s.joinType, s.table, onStr), pms2
}

func (d *selectSqlBuilder) From(sql ...string) *selectSqlBuilder {
	d.from = strings.Join(sql, ",")
	return d
}
func (d *selectSqlBuilder) Where(sql ToSql) *selectSqlBuilder {
	d.where = sql
	return d
}
func (d *selectSqlBuilder) OrderBy(sql string) *selectSqlBuilder {
	d.orderBy = sql
	return d
}
func (d *selectSqlBuilder) GroupBy(sql string) *selectSqlBuilder {
	d.groupBy = sql
	return d
}
func (d *selectSqlBuilder) Last(sql ...ToSql) *selectSqlBuilder {
	d.last = sql
	return d
}
func (d *selectSqlBuilder) LeftJoin(table string) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "left join", table: table}
}
func (d *selectSqlBuilder) InnerJoin(table string) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "right join", table: table}
}
func (d *selectSqlBuilder) RightJoin(table string) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "inner join", table: table}
}
func (d *selectSqlBuilder) Join(table string) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "join", table: table}
}

func (d *selectSqlBuilder) ToSql() (string, []any) {

	buf := bytes.Buffer{}
	parameters := make([]any, 0, 10)
	buf.WriteString("select ")
	if d.fields == nil {
		buf.Write([]byte{byte('*')})
	} else {
		for i, field := range d.fields {
			if i > 0 {
				buf.Write([]byte{byte(',')})
			}
			buf.WriteString(field)
		}
	}
	if d.from != "" {
		buf.WriteString(" from ")
		buf.WriteString(d.from)
	}
	if d.joins != nil {
		for _, join := range d.joins {
			buf.WriteString(" ")
			sqlStr, pms := join.ToSql()
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.WriteString(sqlStr)
		}
	}

	if d.where != nil {
		buf.WriteString(" where ")
		sqlStr, pms := d.where.ToSql()
		if pms != nil {
			parameters = append(parameters, pms...)
		}
		buf.WriteString(sqlStr)
	}
	if d.orderBy != "" {
		buf.WriteString(" order by ")
		buf.WriteString(d.orderBy)
	}
	if d.groupBy != "" {
		buf.WriteString(" group by ")
		buf.WriteString(d.groupBy)
	}
	if d.last != nil {
		buf.Write([]byte{byte(' ')})
		for _, sql := range d.last {
			sqlStr, pms := sql.ToSql()
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.WriteString(sqlStr)
		}
	}
	return buf.String(), parameters
}
func Select(fields ...string) *selectSqlBuilder {
	return &selectSqlBuilder{
		fields: fields,
	}
}

func Count(field string) string {
	return fmt.Sprintf("count(\"%s\")", field)
}
func As(exp, alias string) string {
	return fmt.Sprintf("%s as \"%s\"", exp, alias)
}
func Sum(field string) string {
	return fmt.Sprintf("sum(\"%s\")", field)
}
func Avg(field string) string {
	return fmt.Sprintf("avg(\"%s\")", field)
}
