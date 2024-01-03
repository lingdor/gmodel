package gen

//go:generate go get -u github.com/lingdor/gmodeltool
//go:generate gmodeltool gen schema -dsn 'mysql://127.0.0.1:3306@root:123456/test' -table tb_user -name TbUser
//go:gmodel:start
//somethings
//go:gmodel:end
