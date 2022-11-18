package database

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm/schema"
	"net/netip"
	"reflect"
)

type AutoSerializer struct{}

func (AutoSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)
	if dbValue != nil {
		var bytes []byte
		switch v := dbValue.(type) {
		case []byte:
			bytes = v
		case string:
			bytes = []byte(v)
		default:
			return fmt.Errorf("failed to unmarshal JSONB value: %#v", dbValue)
		}
		switch v := fieldValue.Interface().(type) {
		case *netip.Addr:
			err = v.UnmarshalText(bytes)
		default:
			err = json.Unmarshal(bytes, fieldValue.Interface())
		}
	}
	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

func (AutoSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	switch v := fieldValue.(type) {
	case netip.Addr:
		return []byte(v.String()), nil
	default:
		return json.Marshal(fieldValue)
	}
}
