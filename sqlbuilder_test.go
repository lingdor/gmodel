package gmodel

import (
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
	builder := &deleteSqlBuilder{
		from:  "tb1",
		where: "1=1",
		last:  "limit 1",
	}
	sql := builder.GenSql()
	fmt.Println(sql)
}
