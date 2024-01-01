package gmodel

import (
	"database/sql"
	"reflect"
)

var (
	scanTypeFloat32    = reflect.TypeOf(float32(0))
	scanTypeFloat64    = reflect.TypeOf(float64(0))
	scanTypeInt8       = reflect.TypeOf(int8(0))
	scanTypeInt16      = reflect.TypeOf(int16(0))
	scanTypeInt32      = reflect.TypeOf(int32(0))
	scanTypeInt64      = reflect.TypeOf(int64(0))
	scanTypeInt        = reflect.TypeOf(0)
	scanTypeNullFloat  = reflect.TypeOf(sql.NullFloat64{})
	scanTypeNullInt    = reflect.TypeOf(sql.NullInt64{})
	scanTypeNullTime   = reflect.TypeOf(sql.NullTime{})
	scanTypeUint8      = reflect.TypeOf(uint8(0))
	scanTypeUint16     = reflect.TypeOf(uint16(0))
	scanTypeUint32     = reflect.TypeOf(uint32(0))
	scanTypeUint64     = reflect.TypeOf(uint64(0))
	scanTypeString     = reflect.TypeOf("")
	scanTypeNullString = reflect.TypeOf(sql.NullString{})
	scanTypeBytes      = reflect.TypeOf([]byte{})
	scanTypeUnknown    = reflect.TypeOf(new(interface{}))
)
