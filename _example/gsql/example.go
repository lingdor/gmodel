package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lingdor/gmodel"
	"github.com/lingdor/gmodel/gsql"
	"github.com/lingdor/magicarray/array"
)

func main() {

	var err error
	var db *sql.DB
	if db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db1"); err != nil {
		panic(err)
	}
	var arr array.MagicArray
	//you can use:
	ctx := context.WithValue(context.Background(), gmodel.OptLogSql, 1)
	if arr, err = gmodel.QueryArrRowsContext(ctx, db, gsql.Select().From("tb_user").Where(gsql.Raw("age is not null and age>?", 1)).Last(gsql.Limit(1))); err == nil {
		var bs []byte
		if bs, err = array.JsonMarshal(arr, array.JsonOptDefaultNamingUnderscoreToCamelCase(), array.JsonOptIndent4()); err == nil {
			fmt.Println(string(bs))
		}
	}
	if err != nil {
		panic(err)
	}

	where := gsql.Gt("age", 10).And(gsql.Raw("age is not null"))
	selectSql := gsql.Select(gsql.As(gsql.Sum("age"), "x")).From("tb_user")
	selectSql = selectSql.Where(where).Last(gsql.Limit(1))
	var val array.ZVal
	if val, err = gmodel.QueryValContext(ctx, db, selectSql); err == nil {
		fmt.Printf("sum age:%d\n", val.MustInt())
	}
	if err != nil {
		panic(err)
	}
}
