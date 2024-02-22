package orm

import "testing"

func TestName(t *testing.T) {

	dbname := ToDbName("TbName", "tt,ok")
	if dbname != "tt" {
		t.Errorf("dbname not tt, is:%s", dbname)
	}
	dbname = ToDbName("TbName", "")
	if dbname != "tb_name" {
		t.Errorf("dbname not tb_name, is:%s", dbname)
	}

}
