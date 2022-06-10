package conversion

import (
	"github.com/HaHadaxigua/conversion/constraints"

	"github.com/spf13/cast"
)

// -----------------------------------------------------------------------------

func Float32[T constraints.FloatPtr | constraints.Integer | constraints.IntegerPtr](t T) float32 {
	return cast.ToFloat32(t)
}

func Float32ptr[T constraints.Float | constraints.Integer | constraints.IntegerPtr](t T) *float32 {
	v := cast.ToFloat32(t)
	return &v
}

// -----------------------------------------------------------------------------

func Float64[T constraints.FloatPtr | constraints.Integer | constraints.IntegerPtr](t T) float64 {
	return cast.ToFloat64(t)
}

func Float64ptr[T constraints.Float | constraints.Integer | constraints.IntegerPtr](t T) *float64 {
	v := cast.ToFloat64(t)
	return &v
}
