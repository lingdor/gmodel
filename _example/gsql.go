package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lingdor/gmodel"
	"github.com/lingdor/gmodel/gsql"
	"github.com/lingdor/magicarray/array"
)

func simpleCommand() {

	var err error
	var db *sql.DB
	if db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db1"); err != nil {
		panic(err)
	}
	var arr array.MagicArray
	//you can use:
	ctx := context.WithValue(context.Background(), gmodel.OptLogSql, 1)
	if arr, err = gmodel.QueryArrRowsContext(ctx, db, gsql.Select().From("user1").Where(gsql.Raw("age is not null and age>?", 1)).Last(gsql.Limit(1))); err == nil {
		var bs []byte
		if bs, err = array.JsonMarshal(arr, array.JsonOptDefaultNamingUnderscoreToHump(), array.JsonOptIndent4()); err == nil {
			fmt.Println(string(bs))
		}
	}
	if err != nil {
		panic(err)
	}

	where := gsql.Gt("age", 1).And(gsql.Raw("age is not null"))
	selectSql := gsql.Select(gsql.As(gsql.Sum("age"), "x")).From("user1")
	selectSql = selectSql.Where(where).Last(gsql.Limit(1))
	var val array.ZVal
	if val, err = gmodel.QueryValContext(ctx, db, selectSql); err == nil {
		fmt.Printf("sum age:%d\n", val.MustInt())
	}
	if err != nil {
		panic(err)
	}
}
