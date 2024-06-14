package gsql

import (
	"bytes"
	"fmt"
	"github.com/lingdor/gmodel/common"
)

type sqlWhereBuilder struct {
	left     any
	statment string
	right    ToSql
}

func (s *sqlWhereBuilder) ToSql(config common.ToSqlConfig) (string, []any) {
	if s.left == nil {
		return s.right.ToSql(config)
	}
	sqlStr, pms := s.right.ToSql(config)
	if tosql, ok := s.left.(ToSql); ok {
		leftsql, leftpms := tosql.ToSql(config)
		if leftpms != nil {
			pms = append(pms, leftpms...)
		}
		return fmt.Sprintf("%s %s %s", leftsql, s.statment, sqlStr), pms
	} else if fieldstr, ok := s.left.(string); ok {
		return fmt.Sprintf("%s %s %s", fieldstr, s.statment, sqlStr), pms
	} else {
		return fmt.Sprintf("%s %s %s", fmt.Sprint(s.left), s.statment, sqlStr), pms
	}
}

func (s *sqlWhereBuilder) Or(sql ToSql) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     s,
		statment: "or",
		right:    sql,
	}
}

func (s *sqlWhereBuilder) And(sql ToSql) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     s,
		statment: "and",
		right:    sql,
	}
}

func anyToSql(val any) ToSql {

	if v, ok := val.(ToSql); ok {
		return v
	}
	return Raw("?", val)
}

func Eq(field string, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "=",
		right:    anyToSql(val),
	}
}
func Gt(field string, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: ">",
		right:    anyToSql(val),
	}
}
func Lt(field string, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "<",
		right:    anyToSql(val),
	}
}
func Le(field string, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "<=",
		right:    anyToSql(val),
	}
}
func Ge(field string, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: ">=",
		right:    anyToSql(val),
	}
}
func NotEq(field string, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "<>",
		right:    anyToSql(val),
	}
}
func IsNull(field string) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "",
		right:    Sql("is null"),
	}
}
func IsNotNull(field string) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "",
		right:    Sql("is not null"),
	}
}
func Between(field string, val1, val2 any) *sqlWhereBuilder {

	return &sqlWhereBuilder{
		left:     field,
		statment: "between",
		right: &sqlWhereBuilder{
			left:     anyToSql(val1),
			statment: "and",
			right:    anyToSql(val2),
		},
	}
}

type valuesSqlBuilder []any

func (v valuesSqlBuilder) ToSql(config common.ToSqlConfig) (sql string, pms []any) {
	builder := &bytes.Buffer{}
	builder.Write([]byte("("))
	pms = make([]any, 0, len(v))
	for i, item := range v {
		if i > 0 {
			builder.WriteString(",")
		}
		if sql, ok := item.(ToSql); !ok {
			builder.WriteString("?")
			pms = append(pms, item)
		} else {
			sqlStr, sqlpms := sql.ToSql(config)
			if sqlpms != nil {
				pms = append(pms, sqlpms...)
			}
			builder.WriteString(sqlStr)
		}
	}
	builder.Write([]byte(")"))
	sql = builder.String()
	return
}

func In(field string, vals ...any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "in",
		right:    valuesSqlBuilder(vals),
	}
}
func NotIn(field string, vals ...any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "not in",
		right:    valuesSqlBuilder(vals),
	}
}

type sqlBuilderGroup struct {
	sql ToSql
}

func (s sqlBuilderGroup) ToSql(config common.ToSqlConfig) (string, []any) {
	sql, pms := s.sql.ToSql(config)
	return fmt.Sprintf("(%s)", sql), pms
}

func Group(sql ToSql) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:  nil,
		right: sql,
	}
}
