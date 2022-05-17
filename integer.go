package conversion

import (
	"github.com/spf13/cast"
	"golang.org/x/exp/constraints"
)

type SignedPtr interface {
	~*int | ~*int8 | ~*int16 | ~*int32 | ~*int64
}

type UnsignedPtr interface {
	~*uint | ~*uint8 | ~*uint16 | ~*uint32 | ~*uint64 | ~*uintptr
}

type IntegerPtr interface {
	SignedPtr | UnsignedPtr
}

type FloatPtr interface {
	~*float32 | ~*float64
}

type intConstraints interface {
	constraints.Integer | IntegerPtr
}

// -----------------------------------------------------------------------------

func Int[T intConstraints](t T) int {
	return cast.ToInt(t)
}

func Intptr[T constraints.Integer](t T) *int {
	v := cast.ToInt(t)
	return &v
}

// -----------------------------------------------------------------------------

func Int8[T intConstraints](t T) int8 {
	return cast.ToInt8(t)
}

func Intptr8[T constraints.Integer](t T) *int8 {
	v := cast.ToInt8(t)
	return &v
}

// -----------------------------------------------------------------------------

func Int16[T intConstraints](t T) int16 {
	return cast.ToInt16(t)
}

func Intptr16[T constraints.Integer](t T) *int16 {
	v := cast.ToInt16(t)
	return &v
}

// -----------------------------------------------------------------------------

func Int32[T intConstraints](t T) int32 {
	return cast.ToInt32(t)
}

func Intptr32[T constraints.Integer](t T) *int32 {
	v := cast.ToInt32(t)
	return &v
}

// -----------------------------------------------------------------------------

func Int64[T intConstraints](t T) int64 {
	return cast.ToInt64(t)
}

func Intptr64[T constraints.Integer](t T) *int64 {
	v := cast.ToInt64(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint[T intConstraints](t T) uint {
	return cast.ToUint(t)
}

func Uintptr[T constraints.Integer](t T) *uint {
	v := cast.ToUint(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint8[T intConstraints](t T) uint8 {
	return cast.ToUint8(t)
}

func Uintptr8[T constraints.Integer](t T) *uint8 {
	v := cast.ToUint8(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint16[T intConstraints](t T) uint16 {
	return cast.ToUint16(t)
}

func Uintptr16[T constraints.Integer](t T) *uint16 {
	v := cast.ToUint16(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint32[T intConstraints](t T) uint32 {
	return cast.ToUint32(t)
}

func Uintptr32[T constraints.Integer](t T) *uint32 {
	v := cast.ToUint32(t)
	return &v
}

// -----------------------------------------------------------------------------

func Uint64[T intConstraints](t T) uint64 {
	return cast.ToUint64(t)
}

func Uintptr64[T constraints.Integer](t T) *uint64 {
	v := cast.ToUint64(t)
	return &v
}
