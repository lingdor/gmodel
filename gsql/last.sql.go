package gsql

import (
	"bytes"
	"fmt"
	"strconv"
)

// todo Is support parameters?
type sqlLimit string

func (l *sqlLimit) Offset(v int) sqlLimit {
	*l = sqlLimit(fmt.Sprintf("%s offset %d", *l, v))
	return *l
}

func (l sqlLimit) ToSql() (string, []any) {
	return string(l), nil
}

func Limit(offset1 int, offsets ...int) sqlLimit {
	if len(offsets) > 1 {
		panic(fmt.Sprintf("offset values max 2, but give %d", len(offsets)+1))
	}
	buf := bytes.Buffer{}
	buf.Write([]byte(" limit "))
	buf.Write([]byte(strconv.Itoa(offset1)))
	if len(offsets) == 1 {
		buf.Write([]byte{byte(',')})
		buf.Write([]byte(strconv.Itoa(offsets[0])))
	}
	return sqlLimit(buf.String())
}
