package types

import (
	"github.com/tvarney/illvm/types/typeid"
)

// Uint8 is an 8-bit unsigned integer.
//
// Uint8 is only a Value, and as such may not be pushed onto the stack.
// Upcasting a Uint8 will return a Uint64.
type Uint8 uint8

func (u Uint8) ID() typeid.ID {
	return typeid.Uint8
}

func (u Uint8) Size() int {
	return 1
}

func (u Uint8) Upcast() StackValue {
	return (Uint64)(u)
}

// Uint16 is a 16-bit unsigned integer.
//
// Uint16 is only a Value, and as such may not be pushed onto the stack.
// Upcasting a Uint16 will return a Uint64.
type Uint16 uint16

func (u Uint16) ID() typeid.ID {
	return typeid.Uint16
}

func (u Uint16) Size() int {
	return 2
}

func (u Uint16) Upcast() StackValue {
	return (Uint64)(u)
}

// Uint32 is a 32-bit unsigned integer.
//
// Uint32 is only a Value, and as such may not be pushed onto the stack.
// Upcasting a Uint32 will return a Uint64.
type Uint32 uint32

func (u Uint32) ID() typeid.ID {
	return typeid.Uint32
}

func (u Uint32) Size() int {
	return 4
}

func (u Uint32) Upcast() StackValue {
	return (Uint64)(u)
}

// Uint64 is a 64-bit unsigned integer.
//
// Uint64 implements the StackValue interface, allowing it to be pushed onto
// the stack.
type Uint64 uint64

func (u Uint64) ID() typeid.ID {
	return typeid.Uint64
}

func (u Uint64) Size() int {
	return 8
}

func (u Uint64) Upcast() StackValue {
	return u
}

func (u Uint64) Downcast(to typeid.ID) (Value, error) {
	switch to {
	case typeid.Uint8:
		return (Uint8)(u), nil
	case typeid.Uint16:
		return (Uint16)(u), nil
	case typeid.Uint32:
		return (Uint32)(u), nil
	case typeid.Uint64:
		return u, nil
	case typeid.Int8:
		return (Int8)(u), nil
	case typeid.Int16:
		return (Int16)(u), nil
	case typeid.Int32:
		return (Int32)(u), nil
	case typeid.Int64:
		return (Int64)(u), nil
	case typeid.Float32:
		return (Float32)(u), nil
	case typeid.Float64:
		return (Float64)(u), nil

	}

	return nil, CastError{From: typeid.Uint64, To: to}
}
