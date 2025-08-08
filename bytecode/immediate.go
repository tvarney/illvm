package bytecode

import (
	"fmt"

	"github.com/tvarney/consterr"
)

const (
	// ErrNotEnoughBytes indicates that there were not enough bytes left in a
	// bytecode chunk to read an immediate value.
	ErrNotEnoughBytes consterr.Error = "not enough bytes to read"

	// ErrFetchOverflow indicates that the VM was instructed to read more than
	// 8 bytes.
	ErrFetchOverflow consterr.Error = "fetch overflow"

	// ErrFetchUnderflow indicates that the VM was instructed to read less than
	// 0 bytes.
	ErrFetchUnderflow consterr.Error = "fetch underflow"
)

// FetchU8 reads 1 byte from data starting at pc and returns it as a uint8.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If there aren't enough bytes in the
// data to read 1 byte this function will return ErrNotEnoughBytes.
func FetchU8(data []uint8, pc int) (int, uint8, error) {
	if pc >= len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected 1 bytes", ErrNotEnoughBytes)
	}
	return pc + 1, data[pc], nil
}

// FetchU16 reads 2 bytes from data starting at pc and returns it as a uint16.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If there aren't enough bytes in the
// data to read 2 bytes this function will return ErrNotEnoughBytes.
func FetchU16(data []uint8, pc int) (int, uint16, error) {
	if pc+1 >= len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected 2 bytes", ErrNotEnoughBytes)
	}
	val := (((uint16)(data[pc])) << 8) + ((uint16)(data[pc+1]))
	return pc + 2, val, nil
}

// FetchU24 reads 3 bytes from data starting at pc and returns it as a uint32.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If there aren't enough bytes in the
// data to read 3 bytes this function will return ErrNotEnoughBytes.
func FetchU24(data []uint8, pc int) (int, uint32, error) {
	if pc+2 >= len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected 3 bytes", ErrNotEnoughBytes)
	}
	val := (((uint32)(data[pc])) << 16) + (((uint32)(data[pc+1])) << 8) + ((uint32)(data[pc+2]))
	return pc + 3, val, nil
}

// FetchU32 reads 4 bytes from data starting at pc and returns it as a uint32.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If there aren't enough bytes in the
// data to read 4 bytes this function will return ErrNotEnoughBytes.
func FetchU32(data []uint8, pc int) (int, uint32, error) {
	if pc+3 >= len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected 4 bytes", ErrNotEnoughBytes)
	}
	val := (((uint32)(data[pc])) << 24) + (((uint32)(data[pc+1])) << 16) +
		(((uint32)(data[pc+2])) << 8) + ((uint32)(data[pc+3]))
	return pc + 4, val, nil
}

// FetchU40 reads 5 bytes from data starting at pc and returns it as a uint64.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If there aren't enough bytes in the
// data to read 5 bytes this function will return ErrNotEnoughBytes.
func FetchU40(data []uint8, pc int) (int, uint64, error) {
	if pc+4 >= len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected 5 bytes", ErrNotEnoughBytes)
	}
	val := (((uint64)(data[pc])) << 32) + (((uint64)(data[pc+1])) << 24) +
		(((uint64)(data[pc+2])) << 16) + (((uint64)(data[pc+3])) << 8) +
		((uint64)(data[pc+4]))
	return pc + 5, val, nil
}

// FetchU48 reads 6 bytes from data starting at pc and returns it as a uint64.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If there aren't enough bytes in the
// data to read 6 bytes this function will return ErrNotEnoughBytes.
func FetchU48(data []uint8, pc int) (int, uint64, error) {
	if pc+5 >= len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected 6 bytes", ErrNotEnoughBytes)
	}
	val := (((uint64)(data[pc])) << 40) + (((uint64)(data[pc+1])) << 32) +
		(((uint64)(data[pc+2])) << 24) + (((uint64)(data[pc+3])) << 16) +
		(((uint64)(data[pc+4])) << 8) + ((uint64)(data[pc+5]))
	return pc + 6, val, nil
}

// FetchU56 reads 7 bytes from data starting at pc and returns it as a uint64.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If there aren't enough bytes in the
// data to read 7 bytes this function will return ErrNotEnoughBytes.
func FetchU56(data []uint8, pc int) (int, uint64, error) {
	if pc+6 >= len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected 7 bytes", ErrNotEnoughBytes)
	}
	val := (((uint64)(data[pc])) << 48) + (((uint64)(data[pc+1])) << 40) +
		(((uint64)(data[pc+2])) << 32) + (((uint64)(data[pc+3])) << 24) +
		(((uint64)(data[pc+4])) << 16) + (((uint64)(data[pc+5])) << 8) +
		((uint64)(data[pc+6]))
	return pc + 7, val, nil
}

// FetchU64 reads 8 bytes from data starting at pc and returns it as a uint64.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If there aren't enough bytes in the
// data to read 8 bytes this function will return ErrNotEnoughBytes.
func FetchU64(data []uint8, pc int) (int, uint64, error) {
	if pc+7 >= len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected 8 bytes", ErrNotEnoughBytes)
	}
	val := (((uint64)(data[pc])) << 56) + (((uint64)(data[pc+1])) << 48) +
		(((uint64)(data[pc+2])) << 40) + (((uint64)(data[pc+3])) << 32) +
		(((uint64)(data[pc+4])) << 24) + (((uint64)(data[pc+5])) << 16) +
		(((uint64)(data[pc+6])) << 8) + ((uint64)(data[pc+7]))
	return pc + 8, val, nil
}

// FetchN reads N bytes from data starting at pc an returns it as a uint64.
//
// This function will return the new PC value after the read as the first value,
// with the actual value read as the second. If the count is negative this
// function will return ErrFetchUnderflow. If the count is greater than 8 this
// function will return ErrFetchOverflow. If the count is between 0 and 8 but
// there aren't enough bytes in the data to read that many bytes this function
// will return ErrNotEnoughBytes.
func FetchN(data []uint8, pc, count int) (int, uint64, error) {
	if count < 0 {
		return pc, 0, ErrFetchUnderflow
	}
	if count > 8 {
		return pc, 0, ErrFetchOverflow
	}
	if pc+count > len(data) || pc < 0 {
		return pc, 0, fmt.Errorf("%w: expected %d bytes", ErrNotEnoughBytes, count)
	}

	var accumulator uint64
	for i := range count {
		accumulator = (accumulator << 8) + ((uint64)(data[pc+i]))
	}
	return pc + count, accumulator, nil
}
