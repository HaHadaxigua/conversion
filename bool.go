package conversion

import "github.com/spf13/cast"

func Bool[T ~*bool](t T) bool {
	if t == nil {
		return false
	}
	return *t
}

func BoolPtr[T ~bool](t T) *bool {
	v := cast.ToBool(t)
	return &v
}
