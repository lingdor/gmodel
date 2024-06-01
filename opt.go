package gmodel

// QueryOptNoType If true, will be []byte forever when select fields
// Mybe if you use the MagicArray, the type translate is unnecessary

type optQueryNoTypeT struct{}

var OptQueryNoType optQueryNoTypeT

type optLogSqlT struct{}

var OptLogSql optLogSqlT
