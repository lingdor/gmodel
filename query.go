package gmodel

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lingdor/gmodel/common"
	"github.com/lingdor/magicarray/array"
	"github.com/lingdor/magicarray/zval"
	"log"
	"reflect"
	"strings"
)

func toSqlCall(ctx context.Context, toSql ToSql, config common.ToSqlConfig) (sqlStr string, pms []any) {
	sqlStr, pms = toSql.ToSql(config)
	if ctx.Value(OptLogSql) != nil {
		log.Printf("sql:%s, parameters:%+v\n", sqlStr, pms)
	}
	return
}

// todo 1. tosql input tosqlconfig
// 2. struct get scan pointers
func QueryMapContext(ctx context.Context, db DBHandler, toSql ToSql) (ret map[string]any, err error) {
	tosqlConfig := GetToSqlConfig(db)
	sqlStr, ps := toSqlCall(ctx, toSql, tosqlConfig)
	var rows *sql.Rows
	if rows, err = db.QueryContext(ctx, sqlStr, ps...); err == nil {
		defer func() {
			err2 := rows.Close()
			if err2 != nil && err == nil {
				err = err2
			}
		}()
		var cols []*sql.ColumnType
		if cols, err = rows.ColumnTypes(); err == nil {
			if rows.Next() {
				ret = make(map[string]any, len(cols))
				vals := make([]any, len(cols))
				var pms = make([]ScanParam, len(cols))
				for i, k := range cols {
					pms[i] = newVal(ctx, k)
					vals[i] = pms[i].GetPointer()
				}
				if err = rows.Scan(vals...); err == nil {
					for i, col := range cols {
						ret[col.Name()] = pms[i].GetVal()
					}
					//ret = append(ret, row)
				}

				//for i, k := range cols {
				//ret[k.Name()], vals[i] = newVal(ctx, k)
				//}
				//err = rows.Scan(vals...)
			}
		}
	}
	return
}

func newVal(ctx context.Context, columnType *sql.ColumnType) ScanParam {
	//todo
	//func newVal(ctx context.Context, columnType *sql.ColumnType, val *any) any {
	if opt := ctx.Value(OptQueryNoType); opt != nil {
		var val any
		return &DefaultScanParam[any]{v: &val}
	}
	switch columnType.ScanType() {
	case scanTypeInt:
		var v int
		pp := &v
		return &DefaultScanParam[*int]{v: &pp}
	case scanTypeString:
		var v string
		pp := &v
		return &DefaultScanParam[*string]{v: &pp}
	case scanTypeBytes:
		var v interface{}
		return &DefaultScanParam[any]{v: &v}
	case scanTypeFloat32:
		var v float32
		var pp = &v
		return &DefaultScanParam[*float32]{v: &pp}
	case scanTypeFloat64:
		var v float64
		var pp = &v
		return &DefaultScanParam[*float64]{v: &pp}
	case scanTypeInt8:
		var v int8
		var pp = &v
		return &DefaultScanParam[*int8]{v: &pp}
	case scanTypeInt16:
		var v int16
		var pp = &v
		return &DefaultScanParam[*int16]{v: &pp}
	case scanTypeInt32:
		var v int32
		var pp = &v
		return &DefaultScanParam[*int32]{v: &pp}
	case scanTypeInt64:
		var v int64
		var pp = &v
		return &DefaultScanParam[*int64]{v: &pp}
	case scanTypeNullFloat:
		var v sql.NullFloat64
		return &DefaultScanParam[sql.NullFloat64]{v: &v}
	case scanTypeNullInt:
		var v sql.NullInt64
		return &DefaultScanParam[sql.NullInt64]{v: &v}
	case scanTypeNullString:
		var v sql.NullString
		return &DefaultScanParam[sql.NullString]{v: &v}
	case scanTypeNullTime:
		//var v sql.NullTime
		var v string
		var pp = &v
		return &DefaultScanParam[*string]{v: &pp}
	case scanTypeUint8:
		var v uint8
		var pp = &v
		return &DefaultScanParam[*uint8]{v: &pp}
	case scanTypeUint16:
		var v uint16
		var pp = &v
		return &DefaultScanParam[*uint16]{v: &pp}
	case scanTypeUint32:
		var v uint32
		var pp = &v
		return &DefaultScanParam[*uint32]{v: &pp}
	case scanTypeUint64:
		var v uint64
		var pp = &v
		return &DefaultScanParam[*uint64]{v: &pp}
	}
	switch strings.ToLower(columnType.DatabaseTypeName()) {
	case "decimal", "numeric", "double":
		var v float64
		var pp = &v
		return &DefaultScanParam[*float64]{v: &pp}
	case "float":
		var v float32
		var pp = &v
		return &DefaultScanParam[*float32]{v: &pp}
	}
	//switch strings.ToLower(columnType.DatabaseTypeName()) {
	//case "decimal", "numeric", "double", "float":
	//	var v string
	//	return &v
	//}

	{
		var v any
		return &DefaultScanParam[any]{v: &v}
	}
}

func QueryMapRowsContext(ctx context.Context, db DBHandler, toSql ToSql) (ret []map[string]any, err error) {
	tosqlConfig := GetToSqlConfig(db)
	sqlStr, ps := toSqlCall(ctx, toSql, tosqlConfig)
	ret = make([]map[string]any, 0, 10)
	var rows *sql.Rows
	if rows, err = db.QueryContext(ctx, sqlStr, ps...); err == nil {
		defer func() {
			err2 := rows.Close()
			if err2 != nil && err == nil {
				err = err2
			}
		}()
		var cols []*sql.ColumnType
		if cols, err = rows.ColumnTypes(); err == nil {
			for rows.Next() {
				row := make(map[string]any, len(cols))
				vals := make([]any, len(cols))
				var pms = make([]ScanParam, len(cols))
				for i, k := range cols {
					pms[i] = newVal(ctx, k)
					vals[i] = pms[i].GetPointer()
					//row[k.Name()],
				}
				if err = rows.Scan(vals...); err == nil {
					for i, col := range cols {
						row[col.Name()] = pms[i].GetVal()
					}
					ret = append(ret, row)
				}
			}
		}
	}
	return
}
func QueryValContext(ctx context.Context, db DBHandler, toSql ToSql) (val array.ZVal, err error) {
	tosqlConfig := GetToSqlConfig(db)
	sqlStr, ps := toSqlCall(ctx, toSql, tosqlConfig)
	var rows *sql.Rows
	if rows, err = db.QueryContext(ctx, sqlStr, ps...); err == nil {
		defer func() {
			err2 := rows.Close()
			if err2 != nil && err == nil {
				err = err2
			}
		}()
		if rows.Next() {
			var cols []*sql.ColumnType
			if cols, err = rows.ColumnTypes(); err == nil {
				if len(cols) < 1 {
					val = zval.NewZValNil()
					return
				}
				var scanval = newVal(ctx, cols[0])
				if err = rows.Scan(scanval.GetPointer()); err == nil {
					val = zval.NewZVal(scanval.GetVal())
					return
				}
			} else {
				return
			}
		}
	}
	return
}

func QueryArrContext(ctx context.Context, db DBHandler, toSql ToSql) (arr array.MagicArray, err error) {
	var mp map[string]any
	if mp, err = QueryMapContext(ctx, db, toSql); err == nil {
		return array.Valueof(mp)
	}
	return
}
func QueryArrRowsContext(ctx context.Context, db DBHandler, toSql ToSql) (arr array.MagicArray, err error) {
	var mps []map[string]any
	if mps, err = QueryMapRowsContext(ctx, db, toSql); err == nil {
		return array.ValueOfSlice(mps), nil
	}
	return
}

func QueryEntitiesContext[T any](ctx context.Context, db DBHandler, toSql ToSql) (entities []T, err error) {
	var entity T
	typ := reflect.TypeOf(entity)
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("QueryEntitiesContext T type must is a struct kind")
	}

	entities = make([]T, 0, 10)
	tosqlConfig := GetToSqlConfig(db)
	sqlStr, ps := toSqlCall(ctx, toSql, tosqlConfig)
	var rows *sql.Rows
	if rows, err = db.QueryContext(ctx, sqlStr, ps...); err == nil {
		defer func() {
			err2 := rows.Close()
			if err2 != nil && err == nil {
				err = err2
			}
		}()
		var cols []string
		if cols, err = rows.Columns(); err == nil {
			for rows.Next() {
				var entity T
				var obj = any(&entity)
				//= reflect.New(typ).Interface()
				points := make([]any, len(cols))
				var lack = true
				if handler, ok := obj.(EntityFieldHandler); ok {
					points, lack = handler.GetFieldsHandler(cols)
				}
				if lack {
					points = common.FillHandlers(&entity, cols, points)
				}
				if err = rows.Scan(points...); err == nil {
					entities = append(entities, entity)
				}

			}
		}
	}
	return
}
func QueryMap(db DBHandler, toSql ToSql) (ret map[string]any, err error) {
	return QueryMapContext(context.Background(), db, toSql)
}

func QueryMapRows(db DBHandler, toSql ToSql) (ret []map[string]any, err error) {
	return QueryMapRowsContext(context.Background(), db, toSql)
}
func QueryVal(db DBHandler, toSql ToSql) (val array.ZVal, err error) {
	return QueryValContext(context.Background(), db, toSql)
}

func QueryArrRows(db DBHandler, toSql ToSql) (arr array.MagicArray, err error) {
	return QueryArrRowsContext(context.Background(), db, toSql)
}

func QueryArr(db DBHandler, toSql ToSql) (arr array.MagicArray, err error) {
	return QueryArrContext(context.Background(), db, toSql)
}
