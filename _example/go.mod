module main

go 1.18.1

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/lib/pq v1.10.9
	github.com/lingdor/gmodel v0.0.7
	github.com/lingdor/magicarray v0.0.7
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	gopkg.in/yaml.v3 v3.0.1
)
replace github.com/lingdor/gmodel v0.0.7 => ../
replace github.com/lingdor/magicarray v0.0.7 => ../../magicarray