package gsql

import (
	"testing"
)

func TestSum(t *testing.T) {

	sql := As(Sum("field1"), "f")
	if sql != "sum(field1) as f" {
		t.Error("sum assert faild")
	}
	//fmt.Println(sql)

}
