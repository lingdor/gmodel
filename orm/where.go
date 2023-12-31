package orm

import (
	"bytes"
	"fmt"
	"github.com/lingdor/gmodel/common"
)

type sqlWhereBuilder struct {
	sql        string
	parameters []any
}

func (s *sqlWhereBuilder) midStatement(statment string, sql ToSql) *sqlWhereBuilder {
	sqlStr, pms := sql.ToSql()
	if pms != nil {
		s.parameters = append(s.parameters, pms...)
	}
	s.sql = fmt.Sprintf("%s %s %s", s.sql, statment, sqlStr)
	return s
}

func (s *sqlWhereBuilder) Or(sql ToSql) *sqlWhereBuilder {
	return s.midStatement("or", sql)
}

func (s *sqlWhereBuilder) And(sql ToSql) *sqlWhereBuilder {
	return s.midStatement("and", sql)
}

func (s *sqlWhereBuilder) ToSql() (string, []any) {
	return s.sql, s.parameters
}

func operatorVal(field string, operator string, val any) *sqlWhereBuilder {

	if sql, ok := val.(ToSql); ok {
		return &sqlWhereBuilder{
			sql:        fmt.Sprintf("%s %s %s", field, operator, common.OnlySql(sql)),
			parameters: []any{},
		}
	}
	return &sqlWhereBuilder{
		sql:        fmt.Sprintf("%s %s ?", field, operator),
		parameters: []any{val},
	}
}

func operator(field string, operator string) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		sql:        fmt.Sprintf("%s %s", field, operator),
		parameters: []any{},
	}
}

func Eq(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, "=", val)
}
func Gt(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, ">", val)
}
func Le(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, "<", val)
}
func LeEq(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, "<=", val)
}
func GtEq(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, ">=", val)
}
func NotEq(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, "<>", val)
}
func IsNull(field string) *sqlWhereBuilder {
	return operator(field, " is null")
}
func IsNotNull(field string) *sqlWhereBuilder {
	return operator(field, " is not null")
}
func Between(field string, val1, val2 any) *sqlWhereBuilder {
	buf := bytes.Buffer{}
	buf.WriteString(field)
	buf.WriteString(" between ")
	var parameters []any = make([]any, 0)
	if v, ok := val1.(ToSql); ok {
		buf.WriteString(common.OnlySql(v))
	} else {
		buf.WriteString("?")
		parameters = append(parameters, val1)
	}
	buf.WriteString(" and ")
	if v, ok := val2.(ToSql); ok {
		buf.WriteString(common.OnlySql(v))
	} else {
		buf.WriteString("?")
		parameters = append(parameters, val1)
	}
	return &sqlWhereBuilder{
		sql:        buf.String(),
		parameters: parameters,
	}
}
func In(field string, vals ...any) *sqlWhereBuilder {
	return in(field, " in", vals...)
}
func NotIn(field string, vals ...any) *sqlWhereBuilder {
	return in(field, " not in", vals...)
}
func in(field string, action string, vals ...any) *sqlWhereBuilder {
	if len(vals) < 1 {
		panic(fmt.Errorf("gsql.In must have parameters, but vals is empty"))
	}
	buf := bytes.Buffer{}
	buf.WriteString(field)
	buf.WriteString(action)
	buf.WriteString(" (")
	var parameters []any = make([]any, 0, len(vals))
	for i, v := range vals {
		if i != 0 {
			buf.Write([]byte{byte(',')})
		}
		if vsql, ok := v.(ToSql); ok {
			sqlStr, pms := vsql.ToSql()
			if pms != nil {
				parameters = append(parameters, pms...)
			}
			buf.WriteString(sqlStr)
		}
		buf.Write([]byte{byte('?')})
		parameters = append(parameters, v)
	}
	buf.WriteString(")")

	return &sqlWhereBuilder{
		sql:        buf.String(),
		parameters: parameters,
	}
}

func Group(sql ToSql) *sqlWhereBuilder {
	sqlStr, pms := sql.ToSql()
	if pms != nil {
		pms = []any{}
	}
	return &sqlWhereBuilder{
		sql:        fmt.Sprintf("(%s)", sqlStr),
		parameters: pms,
	}
}
