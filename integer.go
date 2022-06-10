package conversion

import (
	"github.com/HaHadaxigua/conversion/constraints"

	"github.com/spf13/cast"
)

// -----------------------------------------------------------------------------

func Int[T constraints.IntegerPtr](t T) int {
	return cast.ToInt(t)
}

func Intptr[T constraints.Integer](t T) *int {
	v := cast.ToInt(t)
	return &v
}

// -----------------------------------------------------------------------------

func Int8[T constraints.IntegerPtr](t T) int8 {
	return cast.ToInt8(t)
}

func Int8ptr[T constraints.Integer](t T) *int8 {
	v := cast.ToInt8(t)
	return &v
}

// -----------------------------------------------------------------------------

func Int16[T constraints.IntegerPtr](t T) int16 {
	return cast.ToInt16(t)
}

func Int16ptr[T constraints.Integer](t T) *int16 {
	v := cast.ToInt16(t)
	return &v
}

// -----------------------------------------------------------------------------

func Int32[T constraints.IntegerPtr](t T) int32 {
	return cast.ToInt32(t)
}

func Int32ptr[T constraints.Integer](t T) *int32 {
	v := cast.ToInt32(t)
	return &v
}

// -----------------------------------------------------------------------------

func Int64[T constraints.IntegerPtr](t T) int64 {
	return cast.ToInt64(t)
}

func Int64ptr[T constraints.Integer](t T) *int64 {
	v := cast.ToInt64(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint[T constraints.IntegerPtr](t T) uint {
	return cast.ToUint(t)
}

func Uintptr[T constraints.Integer](t T) *uint {
	v := cast.ToUint(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint8[T constraints.IntegerPtr](t T) uint8 {
	return cast.ToUint8(t)
}

func Uint8ptr[T constraints.Integer](t T) *uint8 {
	v := cast.ToUint8(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint16[T constraints.IntegerPtr](t T) uint16 {
	return cast.ToUint16(t)
}

func Uint16ptr[T constraints.Integer](t T) *uint16 {
	v := cast.ToUint16(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint32[T constraints.IntegerPtr](t T) uint32 {
	return cast.ToUint32(t)
}

func Uint32ptr[T constraints.Integer](t T) *uint32 {
	v := cast.ToUint32(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint64[T constraints.IntegerPtr](t T) uint64 {
	return cast.ToUint64(t)
}

func Uint64ptr[T constraints.Integer](t T) *uint64 {
	v := cast.ToUint64(t)
	return &v
}
