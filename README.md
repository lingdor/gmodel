# GModel
GModel are parts of MVC, meaning for model and running for Go. the go model learning from php framework models,
and Golang Features to operate databases for sql easily.

# HelloWorld!
import package
```shell
go get github.com/lingdor/gmodel
```
gsql package demos:
```go
where := gsql.group(gsql.Equal(TbUserFields.Id, 123).And(gsql.Gt(TbUserFields.Id, 5))).Or(gsql.Like(TbUserFields.Name, "%tom%"))
gmodel.QueryArr(tx, gsql.Select().From("user").Where(where).Last(gsql.Limit(1, 2)))
gmodel.QueryMap(tx, gsql.Select("id",""user_name","age").From("user").Where(gsql.Eq("id",1)).OrderBy(gsql.Asc("id")))
gmodel.QueryArrRows(tx, gsql.Select().From(TbUserSchema).Limit(10))

//Join SQLs:
gmodel.QueryMap(tx,conn,gsql.Select("u.id","u.name").from("user1").
        LeftJoin(gsql.As("school","s")).On(gsql.Eq("s.id",gsql.Sql("u.school_id"))).Where(gsql.Eq("u.age",12))

// update Sqls:
gmodel.Execute(ctx,conn,gsql.Update("user").SetMap(map[string]any{
    "user_name": 'user_updated',
    "age": 22,
}.Where(gmodel.Eq("id",1)).Last(gsql.Limit(1));

gmodel.Execute(ctx,conn,gsql.Update("user").Set("user_name","xx").Set("age",80).Where(gsql.Eq("id",1)).Last(gsql.Limit(1));

//insert gsql
gmodel.Execute(ctx,conn,gsql.Insert("user").MapValues(maps map[string]string{
                                                  "user_name":"user1",
                                                  "age":11,
                                                  }).
                                                  MapValues(maps map[string]string{
                                                    "user_name":"user2",
                                                    "age":12,
                                                    }))
}
gmodel.Execute(ctx,conn,gsql.Insert("user","user_name","age").Values("uname1",12).Values("user2",19);
ivals := map[string]any{
"user_name":   "lily",
"age": 12,
}
gmodel.Execute(ctx, conn, gsql.Insert("user").valuesMap(vals))

```

orm package demos:

```go
 //insert orm
gmodel.Execute(ctx,conn,orm.Insert(UserSchema).ValuesMap(map[Field]any{
    UserSchema.UserName:"user1",
    UserSchema.Age:11,
})
gmodel.Execute(ctx,conn,orm.Insert(Tbuser,UserSchema.UserName,UserSchema.Age).Values("uname1",12).Values("user2",19);
//insert into select:
gmodel.Execute(ctx,conn,orm.Insert(UserSchema).Select().From(UserSchema).Where(orm.Eq(UserSchema.Id,1))
// insert set:
gmodel.Execute(ctx,conn,gsql.Insert(UserSchema).Set(UserSchema.UserName,"testuser").Set(UserSchema.Age,16));

```

