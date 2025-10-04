package vm

import "github.com/tvarney/illvm/vm/vmath"

// FetchUnsigned reads count bytes from the bytecode and returns it as a uint64.
//
// If there are not enough bytes in the bytecode to read count bytes this
// function will return a FetchNotEnoughBytesError.
//
// The value of count must be from 1 to 8 inclusive. Values less than 1 will
// result in this function returning ErrFetchUnderflow. Values greater than 8
// will result in this function returning ErrFetchOverflow.
func (t *Thread) FetchUnsigned(count int) (uint64, error) {
	switch count {
	case 1:
		v, err := t.FetchU8()
		return uint64(v), err
	case 2:
		v, err := t.FetchU16()
		return uint64(v), err
	case 3:
		v, err := t.FetchU24()
		return uint64(v), err
	case 4:
		v, err := t.FetchU32()
		return uint64(v), err
	case 5:
		v, err := t.FetchU40()
		return v, err
	case 6:
		v, err := t.FetchU48()
		return v, err
	case 7:
		v, err := t.FetchU56()
		return v, err
	case 8:
		return t.FetchU64()
	default:
		return 0, ImmediateFetchSizeError{Bytes: count}
	}
}

// FetchU8 reads a single byte from the bytecode and returns it as a uint8.
//
// If there aren't enough bytes in the bytecode to read 1 byte this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchU8() (uint8, error) {
	if t.PC >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 1}
	}

	val := t.Data[t.PC]
	t.PC++

	return val, nil
}

// FetchU16 reads 2 bytes from the bytecode and returns it as a uint16.
//
// If there aren't enough bytes in the bytecode to read 2 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchU16() (uint16, error) {
	if t.PC+1 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 2}
	}

	val := vmath.U16FromBytes(t.Data[t.PC], t.Data[t.PC+1])
	t.PC += 2

	return val, nil
}

// FetchU24 reads 3 bytes from the bytecode and returns it as a uint32.
//
// If there aren't enough bytes in the bytecode to read 3 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchU24() (uint32, error) {
	if t.PC+2 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 3}
	}

	val := vmath.U24FromBytes(t.Data[t.PC], t.Data[t.PC+1], t.Data[t.PC+2])
	t.PC += 3

	return val, nil
}

// FetchU32 reads 4 bytes from the bytecode and returns it as a uint32.
//
// If there aren't enough bytes in the bytecode to read 4 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchU32() (uint32, error) {
	if t.PC+3 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 4}
	}

	val := vmath.U32FromBytes(
		t.Data[t.PC], t.Data[t.PC+1], t.Data[t.PC+2], t.Data[t.PC+3],
	)
	t.PC += 4

	return val, nil
}

// FetchU40 reads 5 bytes from the bytecode and returns it as a uint64.
//
// If there aren't enough bytes in the bytecode to read 5 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchU40() (uint64, error) {
	if t.PC+4 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 5}
	}

	val := vmath.U40FromBytes(
		t.Data[t.PC], t.Data[t.PC+1], t.Data[t.PC+2], t.Data[t.PC+3],
		t.Data[t.PC+4],
	)
	t.PC += 5

	return val, nil
}

// FetchU48 reads 6 bytes from the bytecode and returns it as a uint64.
//
// If there aren't enough bytes in the bytecode to read 6 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchU48() (uint64, error) {
	if t.PC+5 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 6}
	}

	val := vmath.U48FromBytes(
		t.Data[t.PC], t.Data[t.PC+1], t.Data[t.PC+2], t.Data[t.PC+3],
		t.Data[t.PC+4], t.Data[t.PC+5],
	)
	t.PC += 6

	return val, nil
}

// FetchU56 reads 7 bytes from the bytecode and returns it as a uint64.
//
// If there aren't enough bytes in the bytecode to read 7 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchU56() (uint64, error) {
	if t.PC+6 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 7}
	}

	val := vmath.U56FromBytes(
		t.Data[t.PC], t.Data[t.PC+1], t.Data[t.PC+2], t.Data[t.PC+3],
		t.Data[t.PC+4], t.Data[t.PC+5], t.Data[t.PC+6],
	)
	t.PC += 7

	return val, nil
}

// FetchU64 reads 8 bytes from the bytecode and returns it as a uint64.
//
// If there aren't enough bytes in the bytecode to read 8 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchU64() (uint64, error) {
	if t.PC+7 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 8}
	}

	val := vmath.U64FromBytes(
		t.Data[t.PC], t.Data[t.PC+1], t.Data[t.PC+2], t.Data[t.PC+3],
		t.Data[t.PC+4], t.Data[t.PC+5], t.Data[t.PC+6], t.Data[t.PC+7],
	)
	t.PC += 8

	return val, nil
}
