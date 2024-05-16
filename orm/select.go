package orm

import (
	"bytes"
	"fmt"
	"github.com/lingdor/gmodel/common"
)

type selectSqlBuilder struct {
	fields  []ToSql
	from    []ToSql
	joins   []ToSql
	where   ToSql
	orderBy []ToSql
	groupBy []Field
	last    []ToSql
}
type selectJoinSqlBuilder struct {
	selectBuilder *selectSqlBuilder
	table         ToSql
	joinType      string
	on            ToSql
}

func (s *selectJoinSqlBuilder) On(sql ToSql) *selectSqlBuilder {
	s.on = sql
	s.selectBuilder.joins = append(s.selectBuilder.joins, s)
	return s.selectBuilder
}
func (s *selectJoinSqlBuilder) ToSql(config common.ToSqlConfig) (string, []any) {
	tableStr, pms := s.table.ToSql(config)
	onStr, pms2 := s.on.ToSql(config)
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

func (d *selectSqlBuilder) From(sql ...ToSql) *selectSqlBuilder {

	d.from = sql
	return d
}
func (d *selectSqlBuilder) Where(sql ToSql) *selectSqlBuilder {
	d.where = sql
	return d
}
func (d *selectSqlBuilder) OrderBy(sql ...ToSql) *selectSqlBuilder {
	d.orderBy = sql
	return d
}
func (d *selectSqlBuilder) GroupBy(fields ...Field) *selectSqlBuilder {
	d.groupBy = fields
	return d
}
func (d *selectSqlBuilder) Last(sql ...ToSql) *selectSqlBuilder {
	d.last = sql
	return d
}
func (d *selectSqlBuilder) LeftJoin(sql ToSql) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "left join", table: sql, selectBuilder: d}
}
func (d *selectSqlBuilder) InnerJoin(sql ToSql) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "inner join", table: sql, selectBuilder: d}
}
func (d *selectSqlBuilder) RightJoin(sql ToSql) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "right join", table: sql, selectBuilder: d}
}
func (d *selectSqlBuilder) Join(sql ToSql) *selectJoinSqlBuilder {
	return &selectJoinSqlBuilder{joinType: "join", table: sql, selectBuilder: d}
}

func (d *selectSqlBuilder) ToSql(config common.ToSqlConfig) (string, []any) {

	buf := bytes.Buffer{}
	parameters := make([]any, 0, 10)
	buf.WriteString("select ")
	if d.fields == nil || len(d.fields) == 0 {
		buf.Write([]byte{byte('*')})
	} else {
		for i, sql := range d.fields {
			if i > 0 {
				buf.Write([]byte{byte(',')})
			}
			sqlStr, pms := sql.ToSql(config)
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.WriteString(sqlStr)
		}
	}
	if d.from != nil && len(d.from) > 0 {
		buf.WriteString(" from ")
		for i, sql := range d.from {
			if i > 0 {
				buf.Write([]byte{byte(',')})
			}
			sqlStr, pms := sql.ToSql(config)
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.WriteString(sqlStr)
		}
	}

	if d.joins != nil {
		for _, join := range d.joins {
			buf.WriteString(" ")
			sqlStr, pms := join.ToSql(config)
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.WriteString(sqlStr)
		}
	}

	if d.where != nil {
		buf.WriteString(" where ")
		sqlStr, pms := d.where.ToSql(config)
		if pms != nil {
			parameters = append(parameters, pms...)
		}
		buf.WriteString(sqlStr)
	}
	if d.orderBy != nil && len(d.orderBy) > 0 {
		buf.WriteString(" order by ")
		for i, item := range d.orderBy {
			if i > 0 {
				buf.WriteString(",")
			}
			sql, _ := item.ToSql(config)
			buf.WriteString(sql)
		}
	}
	if d.groupBy != nil && len(d.groupBy) > 0 {
		buf.WriteString(" group by ")

		for i, field := range d.groupBy {

			if i > 0 {
				buf.Write([]byte{byte(',')})
			}
			sqlStr, pms := field.ToSql(config)
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.WriteString(sqlStr)
		}
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
func Select(fields ...ToSql) *selectSqlBuilder {
	return &selectSqlBuilder{
		fields: fields,
	}
}

type funcBuilder struct {
	pms      ToSql
	funcName string
}

func (f *funcBuilder) ToSql(config common.ToSqlConfig) (string, []any) {
	sql, pms := f.pms.ToSql(config)
	return fmt.Sprintf("%s(%s)", f.funcName, sql), pms
}

func Sum(field ToSql) *fieldWrapper {

	return WrapField(&funcBuilder{
		pms:      field,
		funcName: "sum",
	})
}
func Avg(field ToSql) *fieldWrapper {

	return WrapField(&funcBuilder{
		pms:      field,
		funcName: "avg",
	})
}
func Count(field ToSql) *fieldWrapper {

	return WrapField(&funcBuilder{
		pms:      field,
		funcName: "count",
	})
}
func CountAll() *fieldWrapper {

	return WrapField(Sql("count(*)"))
}
func Max(field ToSql) *fieldWrapper {

	return WrapField(&funcBuilder{
		pms:      field,
		funcName: "max",
	})
}
func Min(field ToSql) *fieldWrapper {

	return WrapField(&funcBuilder{
		pms:      field,
		funcName: "min",
	})
}

type sortSql struct {
	sortType string
	fields   []Field
}

func Asc(fields ...Field) ToSql {
	return &sortSql{
		sortType: "asc",
		fields:   fields,
	}
	//sortOutput("asc", fields...)
}
func Desc(fields ...Field) ToSql {
	return &sortSql{
		sortType: "desc",
		fields:   fields,
	}
}
func (s *sortSql) ToSql(config common.ToSqlConfig) (string, []any) {

	buf := bytes.Buffer{}
	for i, field := range s.fields {
		if i > 0 {
			buf.WriteString(",")
		}
		fieldSql, _ := field.ToSql(config)
		buf.WriteString(fieldSql)
	}
	buf.WriteString(" ")
	buf.WriteString(s.sortType)
	return buf.String(), nil
}
