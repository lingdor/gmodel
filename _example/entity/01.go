package entity

// you can install gmodeltool with: go install github.com/lingdor/gmodeltool

//go:generate gmodeltool gen schema  --tables tb_% --to-files ./
//go:generate gmodeltool gen entity --tables tb_% --to-files ./ --parse-time
