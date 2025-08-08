package bytecode

import (
	"io"
)

// Writer writes values as bytes to an underlying [io.Writer].
type Writer struct {
	writer io.Writer
}

// NewWriter returns a Writer wrapper around the given [io.Writer].
func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

// WriteVarInt writes a variable number of bytes based on the value of the
// given uint64.
//
// This will write the smallest number of bytes necessary to fully encode the
// value given.
func (w *Writer) WriteVarInt(u uint64) (int, error) {
	switch {
	case u <= 0xFF:
		return w.WriteU8(uint8(u))
	case u <= 0xFFFF:
		return w.WriteU16(uint16(u))
	case u <= 0xFFFFFF:
		return w.WriteU24(uint32(u))
	case u <= 0xFFFFFFFF:
		return w.WriteU32(uint32(u))
	case u <= 0xFFFFFFFFFF:
		return w.WriteU40(u)
	case u <= 0xFFFFFFFFFFFF:
		return w.WriteU48(u)
	case u <= 0xFFFFFFFFFFFFFF:
		return w.WriteU56(u)
	}
	return w.WriteU64(u)
}

// WriteU8 writes a byte to the underlying writer.
func (w *Writer) WriteU8(u uint8) (int, error) {
	return w.writer.Write([]byte{u})
}

// WriteU16 writes the given uint16 as two bytes to the underlying writer.
func (w *Writer) WriteU16(u uint16) (int, error) {
	return w.writer.Write([]byte{
		(uint8)((u & 0xFF00) >> 8),
		(uint8)((u & 0x00FF) >> 0),
	})
}

// WriteU24 writes the given uint32 as three bytes to the underlying writer.
func (w *Writer) WriteU24(u uint32) (int, error) {
	return w.writer.Write([]byte{
		(uint8)((u & 0xFF0000) >> 16),
		(uint8)((u & 0x00FF00) >> 8),
		(uint8)((u & 0x0000FF) >> 0),
	})
}

// WriteU32 writes the given uint32 as four bytes to the underlying writer.
func (w *Writer) WriteU32(u uint32) (int, error) {
	return w.writer.Write([]byte{
		(uint8)((u & 0xFF000000) >> 24),
		(uint8)((u & 0x00FF0000) >> 16),
		(uint8)((u & 0x0000FF00) >> 8),
		(uint8)((u & 0x000000FF) >> 0),
	})
}

// WriteU40 writes the given uint64 as five bytes to the underlying writer.
func (w *Writer) WriteU40(u uint64) (int, error) {
	return w.writer.Write([]byte{
		(uint8)((u & 0xFF00000000) >> 32),
		(uint8)((u & 0x00FF000000) >> 24),
		(uint8)((u & 0x0000FF0000) >> 16),
		(uint8)((u & 0x000000FF00) >> 8),
		(uint8)((u & 0x00000000FF) >> 0),
	})
}

// WriteU48 writes the given uint64 as six bytes to the underlying writer.
func (w *Writer) WriteU48(u uint64) (int, error) {
	return w.writer.Write([]byte{
		(uint8)((u & 0xFF0000000000) >> 40),
		(uint8)((u & 0x00FF00000000) >> 32),
		(uint8)((u & 0x0000FF000000) >> 24),
		(uint8)((u & 0x000000FF0000) >> 16),
		(uint8)((u & 0x00000000FF00) >> 8),
		(uint8)((u & 0x0000000000FF) >> 0),
	})
}

// WriteU56 writes the given uint64 as seven bytes to the underlying writer.
func (w *Writer) WriteU56(u uint64) (int, error) {
	return w.writer.Write([]byte{
		(uint8)((u & 0xFF000000000000) >> 48),
		(uint8)((u & 0x00FF0000000000) >> 40),
		(uint8)((u & 0x0000FF00000000) >> 32),
		(uint8)((u & 0x000000FF000000) >> 24),
		(uint8)((u & 0x00000000FF0000) >> 16),
		(uint8)((u & 0x0000000000FF00) >> 8),
		(uint8)((u & 0x000000000000FF) >> 0),
	})
}

// WriteU64 writes the given uint64 as eight bytes to the underlying writer.
func (w *Writer) WriteU64(u uint64) (int, error) {
	return w.writer.Write([]byte{
		(uint8)((u & 0xFF00000000000000) >> 56),
		(uint8)((u & 0x00FF000000000000) >> 48),
		(uint8)((u & 0x0000FF0000000000) >> 40),
		(uint8)((u & 0x000000FF00000000) >> 32),
		(uint8)((u & 0x00000000FF000000) >> 24),
		(uint8)((u & 0x0000000000FF0000) >> 16),
		(uint8)((u & 0x000000000000FF00) >> 8),
		(uint8)((u & 0x00000000000000FF) >> 0),
	})
}
