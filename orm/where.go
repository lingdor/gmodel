package orm

import (
	"bytes"
	"fmt"
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
func Less(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, "<", val)
}
func LessEqual(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, "<=", val)
}
func GtEqual(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, ">=", val)
}
func NotEqual(field string, val any) *sqlWhereBuilder {
	return operatorVal(field, "<>", val)
}
func IsNull(field string) *sqlWhereBuilder {
	return operator(field, " is null")
}
func IsNotNull(field string) *sqlWhereBuilder {
	return operator(field, " is not null")
}
func Between(field string, val1, val2 any) *sqlWhereBuilder {
	return &sqlWhereBuilder{
		sql:        fmt.Sprintf("%s between ? and ?", field),
		parameters: []any{val1, val2},
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
	buf.Write([]byte(field))
	buf.Write([]byte(action))
	buf.Write([]byte(" ("))
	for i, _ := range vals {
		if i != 0 {
			buf.Write([]byte{byte(',')})
		}
		buf.Write([]byte{byte('?')})
	}
	buf.Write([]byte(")"))

	return &sqlWhereBuilder{
		sql:        buf.String(),
		parameters: vals,
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
