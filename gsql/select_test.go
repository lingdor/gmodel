package gsql

import (
	"testing"
)

func TestSum(t *testing.T) {

	sql, _ := Sum("field1").As("f").ToSql()
	if sql != "sum(field1) as f" {
		t.Error("sum assert faild")
	}
	//fmt.Println(sql)

}
