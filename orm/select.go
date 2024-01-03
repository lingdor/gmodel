package orm

import (
	"bytes"
	"fmt"
	"strings"
)

type selectSqlBuilder struct {
	fields  []ToSql
	from    string
	joins   []ToSql
	where   ToSql
	orderBy string
	groupBy string
	last    []ToSql
}
type selectJoinSqlBuilder struct {
	selectBuilder *selectSqlBuilder
	table         ToSql
	joinType      string
	on            ToSql
}

func (s *selectJoinSqlBuilder) On(sql ToSql) *selectSqlBuilder {
	s.selectBuilder.joins = append(s.selectBuilder.joins, sql)
	return s.selectBuilder
}
func (s *selectJoinSqlBuilder) ToSql() (string, []any) {
	tableStr, pms := s.table.ToSql()
	onStr, pms2 := s.on.ToSql()
	if pms != nil && pms2 != nil {
		if pms == nil {
			pms = make([]any, 0)
		}
		if pms2 == nil {
			pms2 = make([]any, 0)
		}
		pms = append(pms, pms2...)
	}

	return fmt.Sprintf("%s %s on %s", s.joinType, tableStr, onStr), pms
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
func (d *selectSqlBuilder) LeftJoin(sql ToSql) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "left join", table: sql}
}
func (d *selectSqlBuilder) InnerJoin(sql ToSql) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "right join", table: sql}
}
func (d *selectSqlBuilder) RightJoin(sql ToSql) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "inner join", table: sql}
}
func (d *selectSqlBuilder) Join(sql ToSql) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "join", table: sql}
}

func (d *selectSqlBuilder) ToSql() (string, []any) {

	if d.from == "" {
		panic(fmt.Errorf("select sql generate faild, no found parts of:'from'"))
	}
	buf := bytes.Buffer{}
	parameters := make([]any, 0, 10)
	buf.Write([]byte("select "))
	if d.fields == nil {
		buf.Write([]byte{byte('*')})
	} else {
		first := true
		for _, sql := range d.fields {
			if !first {
				buf.Write([]byte{byte(',')})
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
	}
	buf.Write([]byte(" from "))
	buf.Write([]byte(d.from))

	if d.joins != nil {
		for _, join := range d.joins {
			buf.Write([]byte(" "))
			sqlStr, pms := join.ToSql()
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.Write([]byte(sqlStr))
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
	if d.orderBy != "" {
		buf.Write([]byte(" order by "))
		buf.Write([]byte(d.orderBy))
	}
	if d.groupBy != "" {
		buf.Write([]byte(" group by "))
		buf.Write([]byte(d.groupBy))
	}
	if d.last != nil {
		buf.Write([]byte{byte(' ')})
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
func Select(fields ...ToSql) *selectSqlBuilder {
	return &selectSqlBuilder{
		fields: fields,
	}
}

type sumSql string

func (s sumSql) ToSql() (string, []any) {
	return string(s), nil
}
func (s sumSql) As(alias string) sumSql {
	s = sumSql(fmt.Sprintf("sum(%s) as %s", s, alias))
	return s
}

func Sum(field string) sumSql {
	return sumSql(field)
}

type avgSql string

func (s avgSql) ToSql() (string, []any) {
	return string(s), nil
}
func (s avgSql) As(alias string) avgSql {
	s = avgSql(fmt.Sprintf("%s as %s", s, alias))
	return s
}
func Avg(field string) avgSql {
	return avgSql(fmt.Sprintf("avg(\"%s\")", field))
}

type countSql string

func (s countSql) ToSql() (string, []any) {
	return string(s), nil
}
func (s countSql) As(alias string) countSql {
	s = countSql(fmt.Sprintf("%s as %s", s, alias))
	return s
}
func Count(field string) countSql {
	if field == "*" || field == "1" {
		return countSql(fmt.Sprintf("count(%s)", field))
	}
	return countSql(fmt.Sprintf("count(\"%s\")", field))
}
