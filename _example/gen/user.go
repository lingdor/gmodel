package gen

//go:generate go get -u github.com/lingdor/gmodeltool
//go:generate gmodeltool gen schema -dsn 'mysql://127.0.0.1:3306@root:123456/test' -tables tb_user
//go:generate gmodeltool gen entity --gorm -dsn 'mysql://127.0.0.1:3306@root:123456/test' -tables tb_user
//go:gmodel:start
//somethings
//go:gmodel:end
