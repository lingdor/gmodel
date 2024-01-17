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
	ctx := context.WithValue(context.Background(), gmodel.OptQueryNoType, 1)

	if arr, err = gmodel.QueryArrRowsContext(ctx, db, gsql.Raw("select * from tb_user where id=?", 1)); err == nil {
		var bs []byte
		if bs, err = array.JsonMarshal(arr, array.JsonOptDefaultNamingUnderscoreToCamelCase()); err == nil {
			fmt.Println(string(bs))
		}
	}
	if err != nil {
		panic(err)
	}

}
