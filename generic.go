package conversion

import (
	"errors"
	"reflect"

	"github.com/spf13/cast"
)

var unsupportedType = errors.New("unsupported type")

func Convert[T any](f any) (t T, err error) {
	if f == nil {
		return
	}

	exceptType := reflect.TypeOf(t)
	if exceptType.Kind() == reflect.Ptr {
		exceptType = exceptType.Elem()
	}

	switch exceptType.Kind() {
	case reflect.String:
		var tmpT string
		tmpT, err = cast.ToStringE(f)
		if err != nil {
			return
		}

		rv := reflect.New(exceptType)
		rv.Elem().Set(reflect.ValueOf(tmpT))
		t = rv.Elem().Interface().(T)
		return
	default:
		err = unsupportedType
	}

	err = unsupportedType
	return
}
