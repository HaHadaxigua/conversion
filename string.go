package conversion

import "github.com/spf13/cast"

func String[T ~*string](t T) string {
	return cast.ToString(t)
}

func StringPtr[T ~string](t T) *string {
	v := cast.ToString(t)
	return &v
}
