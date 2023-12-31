package gmodel

import (
	"context"
	"database/sql"
	"github.com/lingdor/magicarray/array"
	"github.com/lingdor/magicarray/zval"
	"log"
	"strings"
)

func toSqlCall(ctx context.Context, toSql ToSql) (sqlStr string, pms []any) {
	sqlStr, pms = toSql.ToSql()
	if ctx.Value(OptLogSql) != nil {
		log.Printf("sql:%s, parameters:%+v\n", sqlStr, pms)
	}
	return
}

func QueryMapContext(ctx context.Context, db DBHandler, toSql ToSql) (ret map[string]any, err error) {
	sqlStr, ps := toSqlCall(ctx, toSql)
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
				for i, k := range cols {
					vals[i] = newVal(ctx, k)
					ret[k.Name()] = vals[i]
				}
				err = rows.Scan(vals...)
			}
		}
	}
	return
}
func newVal(ctx context.Context, columnType *sql.ColumnType) any {
	if opt := ctx.Value(OptQueryNoType); opt != nil {
		var val any
		return &val
	}
	switch columnType.ScanType() {
	case scanTypeInt:
		var v int
		return &v
	case scanTypeString:
		var v string
		return &v
	case scanTypeBytes:
		var v interface{}
		return &v
	case scanTypeFloat32:
		var v float32
		return &v
	case scanTypeFloat64:
		var v float64
		return &v
	case scanTypeInt8:
		var v int8
		return &v
	case scanTypeInt16:
		var v int16
		return &v
	case scanTypeInt32:
		var v int32
		return &v
	case scanTypeInt64:
		var v int64
		return &v
	case scanTypeNullFloat:
		var v sql.NullFloat64
		return &v
	case scanTypeNullInt:
		var v sql.NullInt64
		return &v
	case scanTypeNullString:
		var v sql.NullString
		return &v
	case scanTypeNullTime:
		//var v sql.NullTime
		var v string
		return &v
	case scanTypeUint8:
		var v uint8
		return &v
	case scanTypeUint16:
		var v uint16
		return &v
	case scanTypeUint32:
		var v uint32
		return &v
	case scanTypeUint64:
		var v uint64
		return &v
	}
	switch strings.ToLower(columnType.DatabaseTypeName()) {
	case "decimal", "numeric", "double":
		var v float64
		return &v
	case "float":
		var v float32
		return &v
	}
	//switch strings.ToLower(columnType.DatabaseTypeName()) {
	//case "decimal", "numeric", "double", "float":
	//	var v string
	//	return &v
	//}

	{
		var v any
		return &v
	}
}

func QueryMapRowsContext(ctx context.Context, db DBHandler, toSql ToSql) (ret []map[string]any, err error) {
	sqlStr, ps := toSqlCall(ctx, toSql)
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
				for i, k := range cols {
					vals[i] = newVal(ctx, k)
					row[k.Name()] = vals[i]
				}
				if err = rows.Scan(vals...); err == nil {
					ret = append(ret, row)
				}
			}
		}
	}
	return
}
func QueryValContext(ctx context.Context, db DBHandler, toSql ToSql) (val array.ZVal, err error) {
	sqlStr, ps := toSqlCall(ctx, toSql)
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
				var v = newVal(ctx, cols[0])
				if err = rows.Scan(v); err == nil {
					val = zval.NewZVal(v)
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
		return array.ValueofMap(mp), nil
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
