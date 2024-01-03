package gmodel

// QueryOptNoType If true, will be []byte forever when select fields
// Mybe if you use the MagicArray, the type translate is unnecessary
const (
	OptQueryNoType = iota
	OptLogSql
)
