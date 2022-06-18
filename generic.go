package conversion

import (
	"errors"
	"fmt"
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
	case reflect.Interface:
		return any(f), nil
	case reflect.String:
		var tmpT string
		tmpT, err = cast.ToStringE(f)
		if err != nil {
			return
		}

		return buildT[T](tmpT, exceptType)
	case reflect.Int:
		var tmpT int
		tmpT, err = cast.ToIntE(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Int8:
		var tmpT int8
		tmpT, err = cast.ToInt8E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Int16:
		var tmpT int16
		tmpT, err = cast.ToInt16E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Int32:
		var tmpT int32
		tmpT, err = cast.ToInt32E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Int64:
		var tmpT int64
		tmpT, err = cast.ToInt64E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Uint:
		var tmpT uint
		tmpT, err = cast.ToUintE(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Uint8:
		var tmpT uint8
		tmpT, err = cast.ToUint8E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Uint16:
		var tmpT uint16
		tmpT, err = cast.ToUint16E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Uint32:
		var tmpT uint32
		tmpT, err = cast.ToUint32E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Uint64:
		var tmpT uint64
		tmpT, err = cast.ToUint64E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Float32:
		var tmpT float32
		tmpT, err = cast.ToFloat32E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Float64:
		var tmpT float64
		tmpT, err = cast.ToFloat64E(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Bool:
		var tmpT bool
		tmpT, err = cast.ToBoolE(f)
		if err != nil {
			return
		}
		return buildT[T](tmpT, exceptType)
	case reflect.Slice:
		switch any(t).(type) {
		case []int:
			var tmpT []int
			tmpT, err = toSliceE[int](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []int8:
			var tmpT []int8
			tmpT, err = toSliceE[int8](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []int16:
			var tmpT []int16
			tmpT, err = toSliceE[int16](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []int32:
			var tmpT []int32
			tmpT, err = toSliceE[int32](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []int64:
			var tmpT []int64
			tmpT, err = toSliceE[int64](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []uint:
			var tmpT []uint
			tmpT, err = toSliceE[uint](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []uint8:
			var tmpT []uint8
			tmpT, err = toSliceE[uint8](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []uint16:
			var tmpT []uint16
			tmpT, err = toSliceE[uint16](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []uint32:
			var tmpT []uint32
			tmpT, err = toSliceE[uint32](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []uint64:
			var tmpT []uint64
			tmpT, err = toSliceE[uint64](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []bool:
			var tmpT []bool
			tmpT, err = toSliceE[bool](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		case []string:
			var tmpT []string
			tmpT, err = toSliceE[string](f)
			if err != nil {
				return
			}
			return buildT[T](tmpT, exceptType)
		}
	case reflect.Map:

	}

	err = unsupportedType
	return
}

func buildT[T any](tmpT any, exceptType reflect.Type) (t T, err error) {
	rv := reflect.New(exceptType)
	rv.Elem().Set(reflect.ValueOf(tmpT))
	t = rv.Elem().Interface().(T)
	return
}

// toSliceE casts an interface to a []any type.
func toSliceE[T any](i interface{}) ([]T, error) {
	if i == nil {
		return []T{}, fmt.Errorf("unable to cast %#v of type %T to []T", i, i)
	}

	switch v := i.(type) {
	case []T:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]T, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := Convert[T](s.Index(j).Interface())
			if err != nil {
				return []T{}, fmt.Errorf("unable to cast %#v of type %T to []T", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		convert, err := Convert[T](i)
		if err != nil {
			return nil, err
		}
		return []T{convert}, nil
	}
}

// toMapE casts an interface to any map type
func toMapE[K comparable, V any](i interface{}) (map[K]V, error) {
	var m = map[K]V{}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			rk, err := Convert[K](k)
			if err != nil {
				return nil, err
			}
			rv, err := Convert[V](val)
			if err != nil {
				return nil, err
			}
			m[rk] = rv
		}
		return m, nil
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	}
}
