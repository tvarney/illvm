package types

import (
	"github.com/tvarney/illvm/types/typeid"
)

type Float32 float32

func (f Float32) ID() typeid.ID {
	return typeid.Float32
}

func (f Float32) Size() int {
	return 4
}

func (f Float32) Upcast() StackValue {
	return Float64(f)
}

type Float64 float64

func (f Float64) ID() typeid.ID {
	return typeid.Float64
}

func (f Float64) Size() int {
	return 8
}

func (f Float64) Upcast() StackValue {
	return f
}

func (f Float64) Downcast(to typeid.ID) (Value, error) {
	switch to {
	case typeid.Uint8:
		return Uint8(f), nil
	case typeid.Uint16:
		return Uint16(f), nil
	case typeid.Uint32:
		return Uint32(f), nil
	case typeid.Uint64:
		return Uint64(f), nil
	case typeid.Int8:
		return Int8(f), nil
	case typeid.Int16:
		return Int16(f), nil
	case typeid.Int32:
		return Int32(f), nil
	case typeid.Int64:
		return Int64(f), nil
	case typeid.Float32:
		return Float32(f), nil
	case typeid.Float64:
		return f, nil
	default:
		return nil, CastError{From: typeid.Float64, To: to}
	}
}
