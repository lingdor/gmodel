package orm

import (
	"github.com/lingdor/gmodel/common"
	"strings"
	"testing"
)

func TestIn(t *testing.T) {

	var vals = strings.Split("a,b,c", ",")
	str, _ := NotIn(Raw("f1"), vals).ToSql(&common.DefaultConfigLoader{})
	if str != "f1 not in (?,?,?)" {
		t.Error("not in orm faild")
	}
}
