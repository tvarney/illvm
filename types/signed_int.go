package types

import (
	"github.com/tvarney/illvm/types/typeid"
)

// Int8 is an 8-bit signed integer.
//
// Int8 is only a Value, and as such may not be pushed onto the stack.
// Upcasting an Int8 will return an Int64.
type Int8 int8

func (i Int8) ID() typeid.ID {
	return typeid.Int8
}

func (i Int8) Size() int {
	return 1
}

func (i Int8) Upcast() StackValue {
	return Int64(i)
}

// Int16 is a 16-bit signed integer.
//
// Int16 is only a Value, and as such may not be pushed onto the stack.
// Upcasting an Int16 will return an Int64.
type Int16 int16

func (i Int16) ID() typeid.ID {
	return typeid.Int16
}

func (i Int16) Size() int {
	return 2
}

func (i Int16) Upcast() StackValue {
	return Int64(i)
}

// Int32 is a 32-bit signed integer.
//
// Int32 is only a Value, and as such may not be pushed onto the stack.
// Upcasting an Int32 will return an Int64.
type Int32 int32

func (i Int32) ID() typeid.ID {
	return typeid.Int32
}

func (i Int32) Size() int {
	return 4
}

func (i Int32) Upcast() StackValue {
	return Int64(i)
}

// Int64 is a 64-bit signed integer.
//
// Int64 implements the StackValue interface, allowing it to be pushed onto
// the stack.
type Int64 int64

func (i Int64) ID() typeid.ID {
	return typeid.Int64
}

func (i Int64) Size() int {
	return 8
}

func (i Int64) Upcast() StackValue {
	return i
}

func (i Int64) Downcast(to typeid.ID) (Value, error) {
	switch to {
	case typeid.Uint8:
		return Uint8(i), nil
	case typeid.Uint16:
		return Uint16(i), nil
	case typeid.Uint32:
		return Uint32(i), nil
	case typeid.Uint64:
		return Uint64(i), nil
	case typeid.Int8:
		return Int8(i), nil
	case typeid.Int16:
		return Int16(i), nil
	case typeid.Int32:
		return Int32(i), nil
	case typeid.Int64:
		return i, nil
	case typeid.Float32:
		return Float32(i), nil
	case typeid.Float64:
		return Float64(i), nil
	default:
		return nil, CastError{From: typeid.Int64, To: to}
	}
}
