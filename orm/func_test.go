package orm

import (
	"fmt"
	"github.com/lingdor/gmodel/common"
	"testing"
)

func TestWrap(t *testing.T) {

	wrap := NewNameField("field1").As("f")

	sql, _ := wrap.ToSql(common.DefaultConfigLoader{})
	expect := "field1 as f"
	if sql != expect {
		fmt.Printf("expect:%s\n value:%s\n", expect, sql)
	}
	fmt.Println(sql)

}
