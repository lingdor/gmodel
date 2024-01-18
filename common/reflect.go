package common

import (
	"errors"
	"github.com/lingdor/gmodel/stringutil"
	"reflect"
	"time"
	"unsafe"
)

const TagName = "gmodel"

func FillHandlers(entity any, columns []string, handlers []any) []any {
	rval := reflect.ValueOf(entity)
	if len(columns) != len(handlers) {
		panic(errors.New("FillHandlers error, the length of  columns and handlers are not equal"))
	} else if rval.Kind() != reflect.Pointer {
		panic("assert failed for entity parameter,must is pointer type")
	}

	pointer := rval.UnsafePointer()
	rval = rval.Elem()
	rtype := rval.Type()
	for i := 0; i < rval.NumField(); i++ {
		vfield := rval.Field(i)
		field := rtype.Field(i)
		fieldPointer := unsafe.Add(pointer, field.Offset)
		if tag, ok := field.Tag.Lookup(TagName); ok {
			if pos := stringutil.Contains(columns, tag); pos != -1 && handlers[pos] == nil {
				if safePointer := ToSafePointerOfBaseType(vfield.Interface(), fieldPointer); safePointer != nil {
					handlers[pos] = safePointer
				}
			}
		} else {
			fieldName := field.Name
			for pos, col := range columns {
				if handlers[pos] == nil && stringutil.UpperFrist(col) == fieldName {
					if safePointer := ToSafePointerOfBaseType(vfield.Interface(), fieldPointer); safePointer != nil {
						handlers[pos] = safePointer
					}
				}
			}
		}
	}
	for i, v := range handlers {
		if v == nil {
			var v []byte
			handlers[i] = &v
		}
	}
	return handlers
}

func ToSafePointerOfBaseType(val any, addr unsafe.Pointer) any {

	switch val.(type) {
	case *string:
		return (**string)(addr)
	case *int:
		return (**int)(addr)
	case *uint:
		return (**uint)(addr)
	case *int64:
		return (**int64)(addr)
	case *uint64:
		return (**uint64)(addr)
	case *int8:
		return (**int8)(addr)
	case *uint8:
		return (**uint8)(addr)
	case *int16:
		return (**int16)(addr)
	case *uint16:
		return (**uint16)(addr)
	case *int32:
		return (**int32)(addr)
	case *uint32:
		return (**uint32)(addr)
	case *float32:
		return (**float32)(addr)
	case *float64:
		return (**float64)(addr)
	case *time.Time:
		return (**time.Time)(addr)
	case *bool:
		return (**bool)(addr)
	case *[]byte:
		return (**[]byte)(addr)
	case string:
		return (*string)(addr)
	case int:
		return (*int)(addr)
	case uint:
		return (*uint)(addr)
	case int64:
		return (*int64)(addr)
	case uint64:
		return (*uint64)(addr)
	case int8:
		return (*int8)(addr)
	case uint8:
		return (*uint8)(addr)
	case int16:
		return (*int16)(addr)
	case uint16:
		return (*uint16)(addr)
	case int32:
		return (*int32)(addr)
	case uint32:
		return (*uint32)(addr)
	case float32:
		return (*float32)(addr)
	case float64:
		return (*float64)(addr)
	case time.Time:
		return (*time.Time)(addr)
	case bool:
		return (*bool)(addr)
	case []byte:
		return (*[]byte)(addr)
	}
	return nil
}
