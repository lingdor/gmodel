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
	if db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/db1?parseTime=true"); err != nil {
		panic(err)
	}
	//you can use:
	ctx := context.WithValue(context.Background(), gmodel.OptLogSql, 1)

	var entities []entity.TbUserEntity

	selectSql := orm.Select(entity.TbUserSchema.Age.As("a"), entity.TbUserSchema.All()).From(entity.TbUserSchema).Where(orm.Gt(entity.TbUserSchema.Age, 10)).Last(orm.Limit(1))
	if entities, err = gmodel.QueryEntitiesContext[entity.TbUserEntity](ctx, db, selectSql); err == nil {
		arr, _ := array.ValueofStructs(entities)
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
