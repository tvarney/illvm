package vmath_test

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/vm/vmath"
)

func TestUnsignedByteSize(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, vmath.UnsignedByteSize(0))
	for idx, expected := range [64]int{
		1, 1, 1, 1, 1, 1, 1, 1,
		2, 2, 2, 2, 2, 2, 2, 2,
		3, 3, 3, 3, 3, 3, 3, 3,
		4, 4, 4, 4, 4, 4, 4, 4,
		5, 5, 5, 5, 5, 5, 5, 5,
		6, 6, 6, 6, 6, 6, 6, 6,
		7, 7, 7, 7, 7, 7, 7, 7,
		8, 8, 8, 8, 8, 8, 8, 8,
	} {
		value := uint64(1) << idx
		count := vmath.UnsignedByteSize(value)
		assert.Equal(t, expected, count, "bytes for 1 << %d (0x%08X) (bits: %d)", idx, value, bits.Len64(value))
	}
}

func TestU16FromBytes(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		b1       uint8
		b2       uint8
		expected uint16
	}{
		{0x05, 0x06, 0x0506},
		{0x10, 0xAB, 0x10AB},
	} {
		value := vmath.U16FromBytes(test.b1, test.b2)
		require.Equal(t, test.expected, value)
	}
}

func TestU24FromBytes(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		b1       uint8
		b2       uint8
		b3       uint8
		expected uint32
	}{
		{0x10, 0x24, 0x31, 0x102431},
	} {
		value := vmath.U24FromBytes(test.b1, test.b2, test.b3)
		require.Equal(t, test.expected, value)
	}
}

func TestU32FromBytes(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		b1       uint8
		b2       uint8
		b3       uint8
		b4       uint8
		expected uint32
	}{
		{0x01, 0x02, 0x03, 0x04, 0x01020304},
	} {
		value := vmath.U32FromBytes(test.b1, test.b2, test.b3, test.b4)
		require.Equal(t, test.expected, value)
	}
}

func TestU40FromBytes(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		b1       uint8
		b2       uint8
		b3       uint8
		b4       uint8
		b5       uint8
		expected uint64
	}{
		{0x01, 0x02, 0x03, 0x04, 0x05, 0x0102030405},
	} {
		value := vmath.U40FromBytes(test.b1, test.b2, test.b3, test.b4, test.b5)
		require.Equal(t, test.expected, value)
	}
}

func TestU48FromBytes(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		b1       uint8
		b2       uint8
		b3       uint8
		b4       uint8
		b5       uint8
		b6       uint8
		expected uint64
	}{
		{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x010203040506},
	} {
		value := vmath.U48FromBytes(
			test.b1, test.b2, test.b3, test.b4, test.b5, test.b6,
		)
		require.Equal(t, test.expected, value)
	}
}

func TestU56FromBytes(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		b1       uint8
		b2       uint8
		b3       uint8
		b4       uint8
		b5       uint8
		b6       uint8
		b7       uint8
		expected uint64
	}{
		{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x01020304050607},
	} {
		value := vmath.U56FromBytes(
			test.b1, test.b2, test.b3, test.b4, test.b5, test.b6, test.b7,
		)
		require.Equal(t, test.expected, value)
	}
}

func TestU64FromBytes(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		b1       uint8
		b2       uint8
		b3       uint8
		b4       uint8
		b5       uint8
		b6       uint8
		b7       uint8
		b8       uint8
		expected uint64
	}{
		{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x0102030405060708},
	} {
		value := vmath.U64FromBytes(
			test.b1, test.b2, test.b3, test.b4, test.b5, test.b6, test.b7,
			test.b8,
		)
		require.Equal(t, test.expected, value)
	}
}

func TestU16ToBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	require.Equal(t, vmath.U16ToBytes(0x0102), d(0x01, 0x02))
	require.Equal(t, vmath.U16ToBytes(0xABCD), d(0xAB, 0xCD))
	require.Equal(t, vmath.U16ToBytes(0x0000), d(0x00, 0x00))
}

func TestU24ToBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	require.Equal(t, vmath.U24ToBytes(0x000000), d(0x00, 0x00, 0x00))
	require.Equal(t, vmath.U24ToBytes(0x010203), d(0x01, 0x02, 0x03))
	require.Equal(t, vmath.U24ToBytes(0xFFFFFFFF), d(0xFF, 0xFF, 0xFF), "must ignore first byte")
}

func TestU32ToBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	require.Equal(t, vmath.U32ToBytes(0x00000000), d(0x00, 0x00, 0x00, 0x00))
	require.Equal(t, vmath.U32ToBytes(0xFEEDBEEF), d(0xFE, 0xED, 0xBE, 0xEF))
}

func TestU40ToBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	require.Equal(t, vmath.U40ToBytes(0x0000000000), d(0x00, 0x00, 0x00, 0x00, 0x00))
	require.Equal(t, vmath.U40ToBytes(0x123456789A), d(0x12, 0x34, 0x56, 0x78, 0x9A))
	require.Equal(t, vmath.U40ToBytes(0x010203FFFFFFFFFF), d(0xFF, 0xFF, 0xFF, 0xFF, 0xFF), "must ignore first 3 bytes")
}

func TestU48ToBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	require.Equal(t, vmath.U48ToBytes(0x000000000000), d(0x00, 0x00, 0x00, 0x00, 0x00, 0x00))
	require.Equal(t, vmath.U48ToBytes(0x123456789ABC), d(0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC))
	require.Equal(t, vmath.U48ToBytes(0x0102FFFFFFFFFFFF), d(0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF), "must ignore first 2 bytes")
}

func TestU56ToBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	require.Equal(t, vmath.U56ToBytes(0), d(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00))
	require.Equal(t, vmath.U56ToBytes(0x123456789ABCDE), d(0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE))
	require.Equal(t, vmath.U56ToBytes(0xFF01020304050607), d(0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07), "must ignore first byte")
}

func TestU64ToBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	require.Equal(t, vmath.U64ToBytes(0), d(0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00))
	require.Equal(t, vmath.U64ToBytes(0x123456789ABCDEF0), d(0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0))
}

func TestUnsignedFromBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	require.Equal(t, uint64(0), vmath.UnsignedFromBytes(d()), "Len0")
	require.Equal(t, uint64(0x01), vmath.UnsignedFromBytes(d(0x01)), "Len1")
	require.Equal(t, uint64(0xACDC), vmath.UnsignedFromBytes(d(0xAC, 0xDC)), "Len2")
	require.Equal(t, uint64(0x012304), vmath.UnsignedFromBytes(d(0x01, 0x23, 0x04)), "Len3")
	require.Equal(t, uint64(0xFEEDBEEF), vmath.UnsignedFromBytes(d(0xFE, 0xED, 0xBE, 0xEF)), "Len4")
	require.Equal(t, uint64(0x31337BEEF0), vmath.UnsignedFromBytes(d(0x31, 0x33, 0x7B, 0xEE, 0xF0)), "Len5")
	require.Equal(t, uint64(0x010203040506), vmath.UnsignedFromBytes(d(1, 2, 3, 4, 5, 6)), "Len7")
	require.Equal(t, uint64(0xABCDEF12345678), vmath.UnsignedFromBytes(
		d(0xAB, 0xCD, 0xEF, 0x12, 0x34, 0x56, 0x78),
	), "Len8")
	require.Equal(t, uint64(0x123ABC456DEF7890), vmath.UnsignedFromBytes(
		d(0x12, 0x3A, 0xBC, 0x45, 0x6D, 0xEF, 0x78, 0x90, 0xFF),
	), "Len9")
}

func TestUnsignedToBytes(t *testing.T) {
	t.Parallel()

	d := func(data ...uint8) []uint8 { return data }

	// 1 byte
	require.Equal(t, d(0x00), vmath.UnsignedToBytes(0))
	require.Equal(t, d(0x80), vmath.UnsignedToBytes(0x80))
	require.Equal(t, d(0xFF), vmath.UnsignedToBytes(0xFF))
	// 2 bytes
	require.Equal(t, d(0x01, 0x00), vmath.UnsignedToBytes(0x100))
	require.Equal(t, d(0xFF, 0xFF), vmath.UnsignedToBytes(0xFFFF))
	// 3 bytes
	require.Equal(t, d(0x01, 0x00, 0x00), vmath.UnsignedToBytes(0x10000))
	require.Equal(t, d(0xFF, 0xFF, 0xFF), vmath.UnsignedToBytes(0xFFFFFF))
	// 4 bytes
	require.Equal(t, d(0x01, 0x00, 0x00, 0x00), vmath.UnsignedToBytes(0x1000000))
	require.Equal(t, d(0xFF, 0xFF, 0xFF, 0xFF), vmath.UnsignedToBytes(0xFFFFFFFF))
	// 5 bytes
	require.Equal(t, d(0x01, 0x00, 0x00, 0x00, 0x00), vmath.UnsignedToBytes(0x100000000))
	require.Equal(t, d(0xFF, 0xFF, 0xFF, 0xFF, 0xFF), vmath.UnsignedToBytes(0xFFFFFFFFFF))
	// 6 bytes
	require.Equal(t, d(0x01, 0x00, 0x00, 0x00, 0x00, 0x00), vmath.UnsignedToBytes(0x10000000000))
	require.Equal(t, d(0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF), vmath.UnsignedToBytes(0xFFFFFFFFFFFF))
	// 7 bytes
	require.Equal(t, d(0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00), vmath.UnsignedToBytes(0x1000000000000))
	require.Equal(t, d(0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF), vmath.UnsignedToBytes(0xFFFFFFFFFFFFFF))
	// 8 bytes
	require.Equal(t, d(0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00), vmath.UnsignedToBytes(0x100000000000000))
	require.Equal(t, d(0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF), vmath.UnsignedToBytes(0xFFFFFFFFFFFFFFFF))
}
