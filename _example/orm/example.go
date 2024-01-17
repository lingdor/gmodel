package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lingdor/gmodel"
	"github.com/lingdor/gmodel/orm"
	"github.com/lingdor/magicarray/array"
	"main/entity"
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

	selectSql := orm.Select(entity.TbUserSchema.Age.As("a"), entity.TbUserSchema.Name.As("user_name")).From(entity.TbUserSchema).Where(orm.Gt(entity.TbUserSchema.Age, 10)).Last(orm.Limit(1))
	if arr, err = gmodel.QueryArrRowsContext(ctx, db, selectSql); err == nil {
		if bs, err := array.JsonMarshal(arr, array.JsonOptDefaultNamingUnderscoreToCamelCase()); err == nil {
			fmt.Println(string(bs))
		} else {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}
}
