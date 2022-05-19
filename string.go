package conversion

import "github.com/spf13/cast"

func String[T ~*string](t T) string {
	if t == nil {
		return ""
	}
	return cast.ToString(t)
}

func StringOrDefault[T ~*string](t T, defaultValue string) string {
	if t == nil {
		return defaultValue
	}
	return cast.ToString(t)
}

func StringPtr[T ~string](t T) *string {
	if t == "" {
		return nil
	}
	v := cast.ToString(t)
	return &v
}
