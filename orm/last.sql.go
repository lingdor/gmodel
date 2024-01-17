package orm

import (
	"github.com/lingdor/gmodel/common"
)

// todo Is support parameters?
type sqlLimit struct {
	limit1   int
	limit2   int
	isOffset bool
	isLimit2 bool
}

func (l *sqlLimit) Offset(v int) *sqlLimit {
	l.limit2 = v
	l.isOffset = true
	return l
}

func (l sqlLimit) ToSql(config common.ToSqlConfig) (string, []any) {

	if l.isOffset {
		return "limit ? offset ?", []any{l.limit1, l.limit2}
	} else if l.isLimit2 {
		return "limit ?,?", []any{l.limit1, l.limit2}
	}
	return "limit ?", []any{l.limit1}
}

func Limit(limit1 int, limit2 ...int) *sqlLimit {

	val := &sqlLimit{
		limit1:   limit1,
		isOffset: false,
	}
	if len(limit2) > 0 {
		val.limit2 = limit2[0]
		val.isLimit2 = true
	}
	return val

}
