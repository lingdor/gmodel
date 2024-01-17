package gmodel

import (
	"database/sql/driver"
	"github.com/lingdor/gmodel/common"
	"reflect"
	"strings"
)

func init() {
	RegisterToSqlConfigLoader(func(db DBHandler) common.ToSqlConfig {
		tp := reflect.TypeOf(db.Driver())
		if tp.Elem() != nil {
			tp = tp.Elem()
		}
		pkgPath := tp.PkgPath()

		if strings.Contains(pkgPath, "mysql") {
			return &mysqlConfigLoader{}
		} else if strings.Contains(pkgPath, "/pg") {
			return &pgsqlConfigLoader{}
		} else if strings.Contains(pkgPath, "mssql") {
			return &mssqlConfigLoader{}
		}
		return nil
	})
}
func GetToSqlConfig(db DBHandler) common.ToSqlConfig {
	if conf, ok := tosqlConfigCache[db.Driver()]; ok {
		return conf
	}
	for _, loader := range tosqlConfigLoader {
		if conf := loader(db); conf != nil {
			tosqlConfigCache[db.Driver()] = conf
			return conf
		}
	}
	return &common.DefaultConfigLoader{}
}

var tosqlConfigCache = make(map[driver.Driver]common.ToSqlConfig, 1)

var tosqlConfigLoader = make([]func(db DBHandler) common.ToSqlConfig, 0, 1)

func RegisterToSqlConfigLoader(loader func(db DBHandler) common.ToSqlConfig) {
	tosqlConfigLoader = append(tosqlConfigLoader, loader)
}

type mysqlConfigLoader struct{}

func (m mysqlConfigLoader) FieldFormat(name string) string {
	return "`" + name + "`"
}

func (m mysqlConfigLoader) TableFormat(name string) string {
	return "`" + name + "`"
}

type pgsqlConfigLoader struct{}

func (m pgsqlConfigLoader) FieldFormat(name string) string {
	return "\"" + name + "\""
}

func (m pgsqlConfigLoader) TableFormat(name string) string {
	return "\"" + name + "\""
}

type mssqlConfigLoader struct{}

func (m mssqlConfigLoader) FieldFormat(name string) string {
	return "[" + name + "]"
}

func (m mssqlConfigLoader) TableFormat(name string) string {
	return "[" + name + "]"
}
