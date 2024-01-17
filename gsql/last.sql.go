package gsql

import (
	"bytes"
	"fmt"
	"github.com/lingdor/gmodel/common"
	"strconv"
)

// todo Is support parameters?
type sqlLimit struct {
	limit1 int
	limit2 int
	offset int
}

func (l *sqlLimit) Offset(v int) sqlLimit {
	l.offset = v
	return *l
}

func (l sqlLimit) ToSql(config common.ToSqlConfig) (string, []any) {

	var parameters []any
	buf := bytes.Buffer{}
	buf.WriteString(" limit ?")
	parameters = append(parameters, l.limit1)
	if l.limit2 != 0 {
		buf.WriteString(", ?")
		parameters = append(parameters, l.limit2)
	} else if l.offset != 0 {
		buf.WriteString(" offset ")
		buf.WriteString(strconv.Itoa(l.offset))
	}

	return string(buf.String()), parameters
}

func Limit(limit1 int, limit2 ...int) *sqlLimit {

	if len(limit2) > 1 {
		panic(fmt.Sprintf("offset values max 2, but give %d", len(limit2)+1))
	}
	ret := &sqlLimit{
		limit1: limit1,
	}
	if len(limit2) > 0 {
		ret.limit2 = limit2[0]
	}
	return ret
}
