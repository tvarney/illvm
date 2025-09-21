package vm

import (
	"strconv"

	"github.com/tvarney/consterr"
)

const (
	// ErrNotEnoughBytes indicates that there weren't enough bytes in the
	// bytecode to read some number of bytes.
	ErrNotEnoughBytes consterr.Error = "not enough bytes to read"

	// ErrInvalidImmediateSize indicates that an opcode requested an invalid
	// number of bytes for an immediate value.
	//
	// This is generally not possible, but included for completeness.
	ErrInvalidImmediateSize consterr.Error = "invalid immediate size"
)

// FetchNotEnoughBytesError is an error which indicates that a fetch from
// bytecode did not have enough bytes left in the bytecode to complete.
type FetchNotEnoughBytesError struct {
	Bytes int
}

func (e FetchNotEnoughBytesError) Error() string {
	return string(ErrNotEnoughBytes) + ": expected " + strconv.FormatInt(int64(e.Bytes), 10) + " bytes"
}

func (e FetchNotEnoughBytesError) Unwrap() error {
	return ErrNotEnoughBytes
}

// ImmediateFetchSizeError is an error which indciates that an opcode requested
// an invalid number of bytes for an immediate value.
type ImmediateFetchSizeError struct {
	Bytes int
}

func (e ImmediateFetchSizeError) Error() string {
	return string(ErrInvalidImmediateSize) +
		": fetch byte count must be between 1 and 8 bytes, " +
		strconv.FormatInt(int64(e.Bytes), 10) + " bytes were requested"
}

func (e ImmediateFetchSizeError) Unwrap() error {
	return ErrInvalidImmediateSize
}
