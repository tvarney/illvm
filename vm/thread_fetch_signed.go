package vm

const (
	Int24NegativeBit = 0x00800000
	Int24SignExtend  = 0xFF000000
	Int40NegativeBit = 0x0000008000000000
	Int40SignExtend  = 0xFFFFFF0000000000
	Int48NegativeBit = 0x0000800000000000
	Int48SignExtend  = 0xFFFF000000000000
	Int56NegativeBit = 0x0080000000000000
	Int56SignExtend  = 0xFF00000000000000
)

// FetchI8 reads a single byte from the bytecode and returns it as an int8.
//
// If there aren't enough bytes in the bytecode to read 1 byte this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchI8() (int8, error) {
	if t.PC >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 1}
	}

	val := t.Data[t.PC]
	t.PC++

	return int8(val), nil
}

// FetchI16 reads 2 bytes from the bytecode and returns it as an int16.
//
// If there aren't enough bytes in the bytecode to read 2 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchI16() (int16, error) {
	if t.PC+1 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 2}
	}

	val := ((uint16(t.Data[t.PC])) << 8) + (uint16(t.Data[t.PC+1]))
	t.PC += 2

	return int16(val), nil
}

// FetchI24 reads 3 bytes from the bytecode and returns it as an int32.
//
// If there aren't enough bytes in the bytecode to read 3 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchI24() (int32, error) {
	if t.PC+2 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 3}
	}

	val := ((uint32(t.Data[t.PC])) << 16) + ((uint32(t.Data[t.PC+1])) << 8) +
		(uint32(t.Data[t.PC+2]))
	t.PC += 3

	// Sign extend
	if val&Int24NegativeBit != 0 {
		val |= Int24SignExtend
	}

	return int32(val), nil
}

// FetchI32 reads 4 bytes from the bytecode and returns it as an int32.
//
// If there aren't enough bytes in the bytecode to read 4 bytes this function
// will return a FetchNotEnoughBytesError.
func (t *Thread) FetchI32() (int32, error) {
	if t.PC+3 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 4}
	}

	val := ((uint32(t.Data[t.PC])) << 24) + ((uint32(t.Data[t.PC+1])) << 16) +
		((uint32(t.Data[t.PC+2])) << 8) + uint32(t.Data[t.PC+3])
	t.PC += 4

	return int32(val), nil
}

// FetchI40 reads 5 bytes from the bytecode and returns it as an int64.
//
// If there aren't enough bytes in the bytecode to read 5 bytes this function
// will return a FetchNotENoughBytesError.
func (t *Thread) FetchI40() (int64, error) {
	if t.PC+4 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 5}
	}

	val := ((uint64(t.Data[t.PC])) << 32) + ((uint64(t.Data[t.PC+1])) << 24) +
		((uint64(t.Data[t.PC+2])) << 16) + ((uint64(t.Data[t.PC+3])) << 8) +
		uint64(t.Data[t.PC+4])
	t.PC += 5

	// Sign extend
	if val&Int40NegativeBit != 0 {
		val |= Int40SignExtend
	}

	return int64(val), nil
}

// FetchI48 reads 6 bytes from the bytecode and returns it as an int64.
//
// If there aren't enough bytes in the bytecode to read 6 bytes this function
// will return a FetchNotENoughBytesError.
func (t *Thread) FetchI48() (int64, error) {
	if t.PC+5 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 6}
	}

	val := (uint64(t.Data[t.PC]) << 40) + (uint64(t.Data[t.PC+1]) << 32) +
		(uint64(t.Data[t.PC+2]) << 24) + (uint64(t.Data[t.PC+3]) << 16) +
		(uint64(t.Data[t.PC+4]) << 8) + uint64(t.Data[t.PC+5])
	t.PC += 6

	// Sign extend
	if val&Int48NegativeBit != 0 {
		val |= Int48SignExtend
	}

	return int64(val), nil
}

// FetchI56 reads 7 bytes from the bytecode and returns it as an int64.
//
// If there aren't enough bytes in the bytecode to read 7 bytes this function
// will return a FetchNotENoughBytesError.
func (t *Thread) FetchI56() (int64, error) {
	if t.PC+6 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 7}
	}

	val := (uint64(t.Data[t.PC]) << 48) + (uint64(t.Data[t.PC+1]) << 40) +
		(uint64(t.Data[t.PC+2]) << 32) + (uint64(t.Data[t.PC+3]) << 24) +
		(uint64(t.Data[t.PC+4]) << 16) + (uint64(t.Data[t.PC+5]) << 8) +
		uint64(t.Data[t.PC+6])
	t.PC += 7

	// Sign extend
	if val&Int56NegativeBit != 0 {
		val |= Int56SignExtend
	}

	return int64(val), nil
}

// FetchI64 reads 8 bytes from the bytecode and returns it as an int64.
//
// If there aren't enough bytes in the bytecode to read 8 bytes this function
// will return a FetchNotENoughBytesError.
func (t *Thread) FetchI64() (int64, error) {
	if t.PC+7 >= len(t.Data) || t.PC < 0 {
		return 0, FetchNotEnoughBytesError{Bytes: 8}
	}

	val := (uint64(t.Data[t.PC]) << 56) + (uint64(t.Data[t.PC+1]) << 48) +
		(uint64(t.Data[t.PC+2]) << 40) + (uint64(t.Data[t.PC+3]) << 32) +
		(uint64(t.Data[t.PC+4]) << 24) + (uint64(t.Data[t.PC+5]) << 16) +
		(uint64(t.Data[t.PC+6]) << 8) + uint64(t.Data[t.PC+7])
	t.PC += 8

	return int64(val), nil
}
