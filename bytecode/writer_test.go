package bytecode_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/bytecode"
)

func TestWriter(t *testing.T) {
	t.Parallel()
	t.Run("WriteU8", testWriterWriteU8)
	t.Run("WriteU16", testWriterWriteU16)
	t.Run("WriteU24", testWriterWriteU24)
	t.Run("WriteU32", testWriterWriteU32)
	t.Run("WriteU40", testWriterWriteU40)
	t.Run("WriteU48", testWriterWriteU48)
	t.Run("WriteU56", testWriterWriteU56)
	t.Run("WriteU64", testWriterWriteU64)
	t.Run("WriteVarInt", testWriterWriteVarInt)
}

func testWriterWriteU8(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint8
		expected []uint8
	}{
		{"Zero", 0, []uint8{0}},
		{"ValueU8", 0x69, []byte{0x69}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteU8(test.value)
			require.NoError(t, err)
			require.Equal(t, 1, count, "must always return that 1 byte was written")
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}

func testWriterWriteU16(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint16
		expected []uint8
	}{
		{"Zero", 0, []uint8{0, 0}},
		{"ValueSmall", 0x0042, []uint8{0, 0x42}},
		{"ValueLarge", 0x1F20, []uint8{0x1F, 0x20}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteU16(test.value)
			require.NoError(t, err)
			require.Equal(t, 2, count, "must always return that 2 bytes were written")
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}

func testWriterWriteU24(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint32
		expected []uint8
	}{
		{"Zero", 0x00000000, []uint8{0, 0, 0}},
		{"ValueU8", 0x00000079, []uint8{0, 0, 0x79}},
		{"ValueU16", 0x0000ABCD, []uint8{0, 0xAB, 0xCD}},
		{"ValueU24", 0x00F00100, []uint8{0xF0, 0x01, 0}},
		{"ValueU32", 0x11010203, []uint8{0x01, 0x02, 0x03}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteU24(test.value)
			require.NoError(t, err)
			require.Equal(t, 3, count, "must always return that 3 bytes were written")
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}

func testWriterWriteU32(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint32
		expected []uint8
	}{
		{"Zero", 0x00000000, []uint8{0, 0, 0, 0}},
		{"ValueU8", 0x00000079, []uint8{0, 0, 0, 0x79}},
		{"ValueU16", 0x0000ABCD, []uint8{0, 0, 0xAB, 0xCD}},
		{"ValueU24", 0x00F00100, []uint8{0, 0xF0, 0x01, 0}},
		{"ValueU32", 0x3124AB09, []uint8{0x31, 0x24, 0xAB, 0x09}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteU32(test.value)
			require.NoError(t, err)
			require.Equal(t, 4, count, "must always return that 4 bytes were written")
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}

func testWriterWriteU40(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint64
		expected []uint8
	}{
		{"Zero", 0, []byte{0, 0, 0, 0, 0}},
		{"ValueU8", 0x000000000000001F, []byte{0, 0, 0, 0, 0x1F}},
		{"ValueU16", 0x0000000000000A0B, []byte{0, 0, 0, 0x0A, 0x0B}},
		{"ValueU24", 0x0000000000FD118F, []byte{0, 0, 0xFD, 0x11, 0x8F}},
		{"ValueU32", 0x00000000DEADBEEF, []byte{0, 0xDE, 0xAD, 0xBE, 0xEF}},
		{"ValueU40", 0x0000008100000000, []byte{0x81, 0, 0, 0, 0}},
		{"ValueU48", 0x0000AB0102030405, []byte{1, 2, 3, 4, 5}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteU40(test.value)
			require.NoError(t, err)
			require.Equal(t, 5, count, "must always return that 5 bytes were written")
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}

func testWriterWriteU48(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint64
		expected []uint8
	}{
		{"Zero", 0, []byte{0, 0, 0, 0, 0, 0}},
		{"ValueU8", 0x000000000000001F, []byte{0, 0, 0, 0, 0, 0x1F}},
		{"ValueU16", 0x0000000000000A0B, []byte{0, 0, 0, 0, 0x0A, 0x0B}},
		{"ValueU24", 0x0000000000FD118F, []byte{0, 0, 0, 0xFD, 0x11, 0x8F}},
		{"ValueU32", 0x00000000DEADBEEF, []byte{0, 0, 0xDE, 0xAD, 0xBE, 0xEF}},
		{"ValueU40", 0x0000008100000000, []byte{0, 0x81, 0, 0, 0, 0}},
		{"ValueU48", 0x0000ABCDEF123456, []byte{0xAB, 0xCD, 0xEF, 0x12, 0x34, 0x56}},
		{"ValueU56", 0x00123456789ABCDE, []byte{0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteU48(test.value)
			require.NoError(t, err)
			require.Equal(t, 6, count, "must always return that 6 bytes were written")
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}

func testWriterWriteU56(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint64
		expected []uint8
	}{
		{"Zero", 0, []byte{0, 0, 0, 0, 0, 0, 0}},
		{"ValueU8", 0x000000000000001F, []byte{0, 0, 0, 0, 0, 0, 0x1F}},
		{"ValueU16", 0x0000000000000A0B, []byte{0, 0, 0, 0, 0, 0x0A, 0x0B}},
		{"ValueU24", 0x0000000000FD118F, []byte{0, 0, 0, 0, 0xFD, 0x11, 0x8F}},
		{"ValueU32", 0x00000000DEADBEEF, []byte{0, 0, 0, 0xDE, 0xAD, 0xBE, 0xEF}},
		{"ValueU40", 0x0000008100000000, []byte{0, 0, 0x81, 0, 0, 0, 0}},
		{"ValueU48", 0x0000ABCDEF123456, []byte{0, 0xAB, 0xCD, 0xEF, 0x12, 0x34, 0x56}},
		{"ValueU56", 0x00123456789ABCDE, []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}},
		{"ValueU64", 0x21436587A9010203, []byte{0x43, 0x65, 0x87, 0xA9, 0x01, 0x02, 0x03}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteU56(test.value)
			require.NoError(t, err)
			require.Equal(t, 7, count, "must always return that 7 bytes were written")
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}

func testWriterWriteU64(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint64
		expected []uint8
	}{
		{"Zero", 0, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{"ValueU8", 0x000000000000001F, []byte{0, 0, 0, 0, 0, 0, 0, 0x1F}},
		{"ValueU16", 0x0000000000000A0B, []byte{0, 0, 0, 0, 0, 0, 0x0A, 0x0B}},
		{"ValueU24", 0x0000000000FD118F, []byte{0, 0, 0, 0, 0, 0xFD, 0x11, 0x8F}},
		{"ValueU32", 0x00000000DEADBEEF, []byte{0, 0, 0, 0, 0xDE, 0xAD, 0xBE, 0xEF}},
		{"ValueU40", 0x0000008100000000, []byte{0, 0, 0, 0x81, 0, 0, 0, 0}},
		{"ValueU48", 0x0000ABCDEF123456, []byte{0, 0, 0xAB, 0xCD, 0xEF, 0x12, 0x34, 0x56}},
		{"ValueU56", 0x00123456789ABCDE, []byte{0, 0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}},
		{"ValueU64", 0x21436587A9010203, []byte{0x21, 0x43, 0x65, 0x87, 0xA9, 0x01, 0x02, 0x03}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteU64(test.value)
			require.NoError(t, err)
			require.Equal(t, 8, count, "must always return that 8 bytes were written")
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}

func testWriterWriteVarInt(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		value    uint64
		expected []uint8
	}{
		{"Zero", 0, []byte{0}},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			buf := bytes.Buffer{}
			w := bytecode.NewWriter(&buf)
			count, err := w.WriteVarInt(test.value)
			require.NoError(t, err)
			require.Equal(
				t, len(test.expected), count,
				"must write %d bytes but wrote %d", len(test.expected), count,
			)
			require.Len(t, buf.Bytes(), count, "must return the count of bytes written")
			require.Equal(t, test.expected, buf.Bytes())
		})
	}
}
