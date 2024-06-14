package orm

import (
	"bytes"
	"fmt"
	"github.com/lingdor/gmodel/common"
	"reflect"
)

type sqlWhereBuilder struct {
	left     ToSql
	statment string
	right    ToSql
}

func (s *sqlWhereBuilder) ToSql(config common.ToSqlConfig) (sqlStr string, pms []any) {
	if s.left == nil {
		return s.right.ToSql(config)
	}
	rightStr, rightPms := s.right.ToSql(config)
	leftStr, leftPms := s.left.ToSql(config)
	if leftPms != nil {
		pms = append(pms, leftPms...)
	}
	if rightPms != nil {
		pms = append(pms, rightPms...)
	}
	sqlStr = fmt.Sprintf("%s %s %s", leftStr, s.statment, rightStr)
	return
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

func Eq(field ToSql, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "=",
		right:    anyToSql(val),
	}
}
func Gt(field ToSql, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: ">",
		right:    anyToSql(val),
	}
}
func Lt(field ToSql, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "<",
		right:    anyToSql(val),
	}
}
func Le(field ToSql, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "<=",
		right:    anyToSql(val),
	}
}
func Ge(field ToSql, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: ">=",
		right:    anyToSql(val),
	}
}
func NotEq(field ToSql, val any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "<>",
		right:    anyToSql(val),
	}
}
func Like(field ToSql, val string) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "like",
		right:    anyToSql(val),
	}
}
func IsNull(field ToSql) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "",
		right:    Sql("is null"),
	}
}
func IsNotNull(field ToSql) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		left:     field,
		statment: "",
		right:    Sql("is not null"),
	}
}

func Between(field ToSql, val1, val2 any) *sqlWhereBuilder {

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

type valuesSqlBuilder[T any] []T

func (v valuesSqlBuilder[T]) ToSql(config common.ToSqlConfig) (sql string, pms []any) {
	builder := &bytes.Buffer{}
	pms = make([]any, 0, len(v))
	builder.Write([]byte("("))
	for i, item := range v {
		if i > 0 {
			builder.WriteString(",")
		}
		vo := any(item)
		if sql, ok := vo.(ToSql); !ok {
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

func In[T any](field ToSql, vals ...T) *sqlWhereBuilder {

	ret := &sqlWhereBuilder{
		left:     field,
		statment: "in",
		right:    valuesSqlBuilder[T](vals),
	}
	if len(vals) == 1 {
		rval := reflect.ValueOf(vals[0])
		if rval.Kind() == reflect.Slice {
			var objVals = make([]any, rval.Len())
			for i := 0; i < rval.Len(); i++ {
				objVals[i] = rval.Index(i).Interface()
			}
			ret.right = valuesSqlBuilder[any](objVals)
		}
	}
	return ret
}
func NotIn[T any](field ToSql, vals ...T) *sqlWhereBuilder {

	ret := &sqlWhereBuilder{
		left:     field,
		statment: "not in",
		right:    valuesSqlBuilder[T](vals),
	}
	var objVal any = vals
	if len(vals) == 1 {
		switch objVal.(type) {
		case []any:
			rval := reflect.ValueOf(vals[0])
			if rval.Kind() == reflect.Slice {
				var objVals = make([]any, rval.Len())
				for i := 0; i < rval.Len(); i++ {
					objVals[i] = rval.Index(i).Interface()
				}
				ret.right = valuesSqlBuilder[any](objVals)
			}
		}
	}
	return ret
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
