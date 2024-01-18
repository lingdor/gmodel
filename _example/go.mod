module main

go 1.18.1

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/lingdor/gmodel v0.0.7
	github.com/lingdor/magicarray v0.0.7
)

replace github.com/lingdor/gmodel v0.0.7 => ../

replace github.com/lingdor/magicarray v0.0.7 => ../../magicarray
