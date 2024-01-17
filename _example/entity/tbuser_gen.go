package entity

import (
	"github.com/lingdor/gmodel"
	"github.com/lingdor/gmodel/orm"
	"time"
)

//gmodel:gen:schema:tb_user:4cf0cbc7e03556b2ab20fc2e0d9b6102
type TbUserSchemaType struct {
	// Id
	Id gmodel.Field
	// Name
	Name gmodel.Field
	// Age
	Age gmodel.Field
	// Createtime
	Createtime gmodel.Field
}

var TbUserSchema *TbUserSchemaType = &TbUserSchemaType{
	Id:         gmodel.NewField("id", "int unsigned", false, true),
	Name:       gmodel.NewField("name", "varchar(50)", true, false),
	Age:        gmodel.NewField("age", "int", true, false),
	Createtime: gmodel.NewField("createtime", "timestamp", false, false),
}

func (s *TbUserSchemaType) ToSql(config gmodel.ToSqlConfig) (string, []any) {
	return config.TableFormat(s.TableName()), nil
}

func (s *TbUserSchemaType) As(name string) gmodel.ToSql {
	return orm.WrapField(s).As(name)
}

func (s *TbUserSchemaType) TableName() string {
	return "tb_user"
}

//gmodel:gen:end
//gmodel:gen:entity:tb_user:1c8640609a92a5f7cdc856d6758603ee
type TbUserEntity struct {
	// Id
	Id *string `gmodel:"id"` //
	// Name
	Name *string `gmodel:"name"` //
	// Age
	Age *int `gmodel:"age"` //
	// Createtime
	Createtime *time.Time `gmodel:"createtime"` //

}

func (entity *TbUserEntity) GetFieldsHandler(fields []string) ([]any, bool) {
	handlers := make([]any, len(fields))
	lack := false
	for i, field := range fields {
		switch field {
		case "":
		case "id":
			handlers[i] = &entity.Id
		case "name":
			handlers[i] = &entity.Name
		case "age":
			handlers[i] = &entity.Age
		case "createtime":
			handlers[i] = &entity.Createtime
		default:
			lack = true
		}
	}
	return handlers, lack
}

//gmodel:gen:end
