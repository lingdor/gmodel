package entity

//go:generate gmodeltool gen schema -conn db0 --tables "tb_%" --tofiles
//go:generate gmodeltool gen entity -conn db0 --tables "tb_%" --tofiles
