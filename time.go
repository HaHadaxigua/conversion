package conversion

import "time"

func Time[T ~*time.Time](t T) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
