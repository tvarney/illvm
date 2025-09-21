package vm

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

	val := ((uint16(t.Data[t.PC])) << 8) + (uint16(t.Data[t.PC+1]))
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

	val := ((uint32(t.Data[t.PC])) << 16) + ((uint32(t.Data[t.PC+1])) << 8) +
		(uint32(t.Data[t.PC+2]))

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

	val := ((uint32(t.Data[t.PC])) << 24) + ((uint32(t.Data[t.PC+1])) << 16) +
		((uint32(t.Data[t.PC+2])) << 8) + (uint32(t.Data[t.PC+3]))

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

	val := ((uint64(t.Data[t.PC])) << 32) + ((uint64(t.Data[t.PC+1])) << 24) +
		((uint64(t.Data[t.PC+2])) << 16) + ((uint64(t.Data[t.PC+3])) << 8) +
		(uint64(t.Data[t.PC+4]))

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

	val := ((uint64(t.Data[t.PC])) << 40) + ((uint64(t.Data[t.PC+1])) << 32) +
		((uint64(t.Data[t.PC+2])) << 24) + ((uint64(t.Data[t.PC+3])) << 16) +
		((uint64(t.Data[t.PC+4])) << 8) + (uint64(t.Data[t.PC+5]))

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

	val := ((uint64(t.Data[t.PC])) << 48) + ((uint64(t.Data[t.PC+1])) << 40) +
		((uint64(t.Data[t.PC+2])) << 32) + ((uint64(t.Data[t.PC+3])) << 24) +
		((uint64(t.Data[t.PC+4])) << 16) + ((uint64(t.Data[t.PC+5])) << 8) +
		(uint64(t.Data[t.PC+6]))

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

	val := ((uint64(t.Data[t.PC])) << 56) + ((uint64(t.Data[t.PC+1])) << 48) +
		((uint64(t.Data[t.PC+2])) << 40) + ((uint64(t.Data[t.PC+3])) << 32) +
		((uint64(t.Data[t.PC+4])) << 24) + ((uint64(t.Data[t.PC+5])) << 16) +
		((uint64(t.Data[t.PC+6])) << 8) + (uint64(t.Data[t.PC+7]))

	t.PC += 8

	return val, nil
}
