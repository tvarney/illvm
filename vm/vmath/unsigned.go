package vmath

import (
	"math/bits"
)

// UnsignedByteSize returns how many bytes are needed to hold the given uint64.
func UnsignedByteSize(val uint64) int {
	if val == 0 {
		return 1
	}
	bitSize := bits.Len64(val)
	byteSize := bitSize / 8
	if bitSize%8 == 0 {
		return byteSize
	}
	return byteSize + 1
}

// U16FromBytes assembles a u16 value from 2 bytes.
func U16FromBytes(msb, lsb uint8) uint16 {
	return uint16(msb)<<8 | uint16(lsb)
}

// U24FromBytes assembles a u24 value from 3 bytes.
//
// As uint24 does not exist, this function returns a uint32.
func U24FromBytes(msb, b2, lsb uint8) uint32 {
	return uint32(msb)<<16 | uint32(b2)<<8 | uint32(lsb)
}

// U32FromBytes assembles a u32 value from 4 bytes.
func U32FromBytes(msb, b2, b3, lsb uint8) uint32 {
	return uint32(msb)<<24 | uint32(b2)<<16 | uint32(b3)<<8 | uint32(lsb)
}

// U40FromBytes assembles a u40 value from 5 bytes.
//
// As uint40 does not exist, this function returns a uint64.
func U40FromBytes(msb, b2, b3, b4, lsb uint8) uint64 {
	return uint64(msb)<<32 | uint64(b2)<<24 | uint64(b3)<<16 | uint64(b4)<<8 |
		uint64(lsb)
}

// U48FromBytes assembles a u48 value from 6 bytes.
//
// As uint48 does not exist, this function returns a uint64.
func U48FromBytes(msb, b2, b3, b4, b5, lsb uint8) uint64 {
	return uint64(msb)<<40 | uint64(b2)<<32 | uint64(b3)<<24 | uint64(b4)<<16 |
		uint64(b5)<<8 | uint64(lsb)
}

// U56FromBytes assembles a u56 value from 7 bytes.
//
// As uint56 does not exist, this function returns a uint64.
func U56FromBytes(msb, b2, b3, b4, b5, b6, lsb uint8) uint64 {
	return uint64(msb)<<48 | uint64(b2)<<40 | uint64(b3)<<32 | uint64(b4)<<24 |
		uint64(b5)<<16 | uint64(b6)<<8 | uint64(lsb)
}

// U64FromBytes assembles a u64 value from 8 bytes.
func U64FromBytes(msb, b2, b3, b4, b5, b6, b7, lsb uint8) uint64 {
	return uint64(msb)<<56 | uint64(b2)<<48 | uint64(b3)<<40 | uint64(b4)<<32 |
		uint64(b5)<<24 | uint64(b6)<<16 | uint64(b7)<<8 | uint64(lsb)
}

// U16ToBytes decomposes a u16 to an array of 2 bytes.
//
// The resulting array can be used to reassemble the u16 by passing the values
// to U16FromBytes.
func U16ToBytes(val uint16) []uint8 {
	return []uint8{
		uint8((val & 0xFF00) >> 8),
		uint8(val & 0x00FF),
	}
}

// U24ToBytes decomposes a u24 to an array of 3 bytes.
//
// As uint24 does not exist, this function takes a uint32 and ignores the most
// significant byte.
//
// The resulting array can be used to reassemble the u24 by passing the values
// to U24FromBytes.
func U24ToBytes(val uint32) []uint8 {
	return []uint8{
		uint8((val & 0xFF0000) >> 16),
		uint8((val & 0x00FF00) >> 8),
		uint8(val & 0x0000FF),
	}
}

// U32ToBytes decomposes a u32 to an array of 4 bytes.
//
// The resulting array can be used to reassemble the u32 by passing the values
// to U32FromBytes.
func U32ToBytes(val uint32) []uint8 {
	return []uint8{
		uint8((val & 0xFF000000) >> 24),
		uint8((val & 0x00FF0000) >> 16),
		uint8((val & 0x0000FF00) >> 8),
		uint8(val & 0x000000FF),
	}
}

// U40ToBytes decomposes a u40 to an array of 5 bytes.
//
// As uint40 does not exist, this function takes a uint64 and ignores the most
// significant 3 bytes.
//
// The resulting array can be used to reassemble the u40 by passing the values
// to U40FromBytes.
func U40ToBytes(val uint64) []uint8 {
	return []uint8{
		uint8((val & 0xFF00000000) >> 32),
		uint8((val & 0x00FF000000) >> 24),
		uint8((val & 0x0000FF0000) >> 16),
		uint8((val & 0x000000FF00) >> 8),
		uint8(val & 0x00000000FF),
	}
}

// U48ToBytes decomposes a u48 to an array of 6 bytes.
//
// As uint48 does not exist, this function takes a uint64 and ignores the most
// significant 2 bytes.
//
// The resulting array can be used to reassemble the u48 by passing the values
// to U48FromBytes.
func U48ToBytes(val uint64) []uint8 {
	return []uint8{
		uint8((val & 0xFF0000000000) >> 40),
		uint8((val & 0x00FF00000000) >> 32),
		uint8((val & 0x0000FF000000) >> 24),
		uint8((val & 0x000000FF0000) >> 16),
		uint8((val & 0x00000000FF00) >> 8),
		uint8(val & 0x0000000000FF),
	}
}

// U56ToBytes decomposes a u56 to an array of 7 bytes.
//
// As uint56 does not exist, this function takes a uint64 and ignores the most
// significant byte.
//
// The resulting array can be used to reassemble the u56 by passing the values
// to U56FromBytes.
func U56ToBytes(val uint64) []uint8 {
	return []uint8{
		uint8((val & 0xFF000000000000) >> 48),
		uint8((val & 0x00FF0000000000) >> 40),
		uint8((val & 0x0000FF00000000) >> 32),
		uint8((val & 0x000000FF000000) >> 24),
		uint8((val & 0x00000000FF0000) >> 16),
		uint8((val & 0x0000000000FF00) >> 8),
		uint8(val & 0x000000000000FF),
	}
}

// U64ToBytes decomposes a u64 to an array of 8 bytes.
//
// The resulting array can be used to reassemble the u64 by passing the values
// to U64FromBytes.
func U64ToBytes(val uint64) []uint8 {
	return []uint8{
		uint8((val & 0xFF00000000000000) >> 56),
		uint8((val & 0x00FF000000000000) >> 48),
		uint8((val & 0x0000FF0000000000) >> 40),
		uint8((val & 0x000000FF00000000) >> 32),
		uint8((val & 0x00000000FF000000) >> 24),
		uint8((val & 0x0000000000FF0000) >> 16),
		uint8((val & 0x000000000000FF00) >> 8),
		uint8(val & 0x00000000000000FF),
	}
}

// UnsignedToBytes returns the smallest slice of uint8 which holds the given
// value.
func UnsignedToBytes(val uint64) []uint8 {
	switch UnsignedByteSize(val) {
	case 1:
		return []uint8{uint8(val)}
	case 2:
		return U16ToBytes(uint16(val))
	case 3:
		return U24ToBytes(uint32(val))
	case 4:
		return U32ToBytes(uint32(val))
	case 5:
		return U40ToBytes(val)
	case 6:
		return U48ToBytes(val)
	case 7:
		return U56ToBytes(val)
	default:
		return U64ToBytes(val)
	}
}

// UnsignedFromBytes builds an unsigned integer from the given slice of bytes.
//
// If the length of the slice is 0, 0 is returned. If the length of the slice is
// greater than 8 then only the first 8 bytes are returned.
func UnsignedFromBytes(data []uint8) uint64 {
	switch len(data) {
	case 0:
		return 0
	case 1:
		return uint64(data[0])
	case 2:
		return uint64(U16FromBytes(data[0], data[1]))
	case 3:
		return uint64(U24FromBytes(data[0], data[1], data[2]))
	case 4:
		return uint64(U32FromBytes(data[0], data[1], data[2], data[3]))
	case 5:
		return uint64(U40FromBytes(
			data[0], data[1], data[2], data[3], data[4],
		))
	case 6:
		return uint64(U48FromBytes(
			data[0], data[1], data[2], data[3], data[4], data[5],
		))
	case 7:
		return uint64(U56FromBytes(
			data[0], data[1], data[2], data[3], data[4], data[5], data[6],
		))
	default:
		return uint64(U64FromBytes(
			data[0], data[1], data[2], data[3], data[4], data[5], data[6],
			data[7],
		))
	}
}
