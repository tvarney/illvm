package types

import (
	"github.com/tvarney/illvm/types/typeid"
)

// Value is a generic value type.
//
// Values must be able to return their ID and be upcast to a StackValue. They
// are used inside struct and class definitions.
type Value interface {
	Id() typeid.Id
	Size() int
	Upcast() StackValue
}

// StackValue is a Value which may be pushed onto the stack of a thread.
type StackValue interface {
	Value

	Downcast(to typeid.Id) (Value, error)
}
