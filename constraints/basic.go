package constraints

import "golang.org/x/exp/constraints"

type Float interface {
	~float32 | ~float64
}

type FloatPtr interface {
	~*float32 | ~*float64
}

type SignedPtr interface {
	~*int | ~*int8 | ~*int16 | ~*int32 | ~*int64
}

type UnsignedPtr interface {
	~*uint | ~*uint8 | ~*uint16 | ~*uint32 | ~*uint64 | ~*uintptr
}

type Integer interface {
	constraints.Integer
}

type IntegerPtr interface {
	SignedPtr | UnsignedPtr
}
