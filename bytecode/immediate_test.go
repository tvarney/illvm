package bytecode_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/bytecode"
	"github.com/tvarney/testerr"
)

func TestFetchU8(t *testing.T) {
	t.Parallel()

	data1 := []uint8{1, 2, 3, 4, 5}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint8
		errval   testerr.ExpectedError
	}{
		{"Index0", data1, 0, 1, 1, testerr.Nil()},
		{"Index1", data1, 1, 2, 2, testerr.Nil()},
		{"Index2", data1, 2, 3, 3, testerr.Nil()},
		{"Index3", data1, 3, 4, 4, testerr.Nil()},
		{"Index4", data1, 4, 5, 5, testerr.Nil()},
		{"PCOverflow", data1, 5, 5, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCUnderflow", data1, -1, -1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchU8(test.data, test.pc)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestFetchU16(t *testing.T) {
	t.Parallel()

	data := []uint8{0x01, 0x02, 0x03, 0x04, 0x05}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint16
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 2, 0x0102, testerr.Nil()},
		{"Index1", data, 1, 3, 0x0203, testerr.Nil()},
		{"Index2", data, 2, 4, 0x0304, testerr.Nil()},
		{"Index3", data, 3, 5, 0x0405, testerr.Nil()},
		{"PCOverflow1", data, 4, 4, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow2", data, 5, 5, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCUnderflow", data, -1, -1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchU16(test.data, test.pc)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestFetchU24(t *testing.T) {
	t.Parallel()

	data := []uint8{0x01, 0x02, 0x03, 0x04, 0x05}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint32
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 3, 0x010203, testerr.Nil()},
		{"Index1", data, 1, 4, 0x020304, testerr.Nil()},
		{"Index2", data, 2, 5, 0x030405, testerr.Nil()},
		{"PCOverflow1", data, 3, 3, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow2", data, 4, 4, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow3", data, 5, 5, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCUnderflow", data, -1, -1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchU24(test.data, test.pc)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestFetchU32(t *testing.T) {
	t.Parallel()

	data := []uint8{0x01, 0x02, 0x03, 0x04, 0x05}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint32
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 4, 0x01020304, testerr.Nil()},
		{"Index1", data, 1, 5, 0x02030405, testerr.Nil()},
		{"PCOverflow1", data, 2, 2, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow2", data, 3, 3, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow3", data, 4, 4, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow4", data, 5, 5, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCUnderflow", data, -1, -1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchU32(test.data, test.pc)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestFetchU40(t *testing.T) {
	t.Parallel()

	data := []uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 5, 0x0102030405, testerr.Nil()},
		{"Index1", data, 1, 6, 0x0203040506, testerr.Nil()},
		{"Index2", data, 2, 7, 0x0304050607, testerr.Nil()},
		{"Index3", data, 3, 8, 0x0405060708, testerr.Nil()},
		{"PCOverflow1", data, 4, 4, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow2", data, 5, 5, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow3", data, 6, 6, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow4", data, 7, 7, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow5", data, 8, 8, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCUnderflow", data, -1, -1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchU40(test.data, test.pc)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestFetchU48(t *testing.T) {
	t.Parallel()

	data := []uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 6, 0x010203040506, testerr.Nil()},
		{"Index1", data, 1, 7, 0x020304050607, testerr.Nil()},
		{"Index2", data, 2, 8, 0x030405060708, testerr.Nil()},
		{"PCOverflow1", data, 3, 3, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow2", data, 4, 4, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow3", data, 5, 5, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow4", data, 6, 6, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow5", data, 7, 7, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow6", data, 8, 8, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCUnderflow", data, -1, -1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchU48(test.data, test.pc)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestFetchU56(t *testing.T) {
	t.Parallel()

	data := []uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 7, 0x01020304050607, testerr.Nil()},
		{"Index1", data, 1, 8, 0x02030405060708, testerr.Nil()},
		{"PCOverflow1", data, 2, 2, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow2", data, 3, 3, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow3", data, 4, 4, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow4", data, 5, 5, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow5", data, 6, 6, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow6", data, 7, 7, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow7", data, 8, 8, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCUnderflow", data, -1, -1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchU56(test.data, test.pc)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestFetchU64(t *testing.T) {
	t.Parallel()

	data := []uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 8, 0x0102030405060708, testerr.Nil()},
		{"PCOverflow1", data, 1, 1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow2", data, 2, 2, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow3", data, 3, 3, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow4", data, 4, 4, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow5", data, 5, 5, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow6", data, 6, 6, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow7", data, 7, 7, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCOverflow8", data, 8, 8, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
		{"PCUnderflow", data, -1, -1, 0, testerr.Is(bytecode.ErrNotEnoughBytes)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchU64(test.data, test.pc)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestFetchN(t *testing.T) {
	t.Parallel()

	data := []uint8{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}

	errNotEnough := testerr.Is(bytecode.ErrNotEnoughBytes)

	for _, test := range []struct {
		name     string
		data     []uint8
		size     int
		pc       int
		newpc    int
		expected uint64
		errval   testerr.ExpectedError
	}{
		{"CountUnderflow", data, -1, 1, 1, 0, testerr.Is(bytecode.ErrFetchUnderflow)},
		{"CountOverflow", data, 9, 1, 1, 0, testerr.Is(bytecode.ErrFetchOverflow)},
		{"OneByte", data, 1, 0, 1, 0x01, testerr.Nil()},
		{"OneBytePCOverflow", data, 1, 10, 10, 0, errNotEnough},
		{"TwoBytes", data, 2, 5, 7, 0x0607, testerr.Nil()},
		{"TwoBytesPCOverflow1", data, 2, 9, 9, 0, errNotEnough},
		{"TwoBytesPCOverflow2", data, 2, 10, 10, 0, errNotEnough},
		{"ThreeBytes", data, 3, 6, 9, 0x070809, testerr.Nil()},
		{"ThreeBytesPCOverflow1", data, 3, 7, 7, 0, errNotEnough},
		{"ThreeBytesPCOverflow2", data, 3, 8, 8, 0, errNotEnough},
		{"ThreeBytesPCOverflow3", data, 3, 9, 9, 0, errNotEnough},
		{"FourBytes", data, 4, 2, 6, 0x03040506, testerr.Nil()},
		{"FourBytesPCOverflow1", data, 4, 6, 6, 0, errNotEnough},
		{"FourBytesPCOverflow2", data, 4, 7, 7, 0, errNotEnough},
		{"FourBytesPCOverflow3", data, 4, 8, 8, 0, errNotEnough},
		{"FourBytesPCOverflow4", data, 4, 9, 9, 0, errNotEnough},
		{"FiveBytes", data, 5, 1, 6, 0x0203040506, testerr.Nil()},
		{"FiveBytesPCOverflow1", data, 5, 5, 5, 0, errNotEnough},
		{"FiveBytesPCOverflow2", data, 5, 6, 6, 0, errNotEnough},
		{"FiveBytesPCOverflow3", data, 5, 7, 7, 0, errNotEnough},
		{"FiveBytesPCOverflow4", data, 5, 8, 8, 0, errNotEnough},
		{"FiveBytesPCOverflow5", data, 5, 9, 9, 0, errNotEnough},
		{"SixBytes", data, 6, 2, 8, 0x030405060708, testerr.Nil()},
		{"SixBytesPCOverflow1", data, 6, 4, 4, 0, errNotEnough},
		{"SixBytesPCOverflow2", data, 6, 5, 5, 0, errNotEnough},
		{"SixBytesPCOverflow3", data, 6, 6, 6, 0, errNotEnough},
		{"SixBytesPCOverflow4", data, 6, 7, 7, 0, errNotEnough},
		{"SixBytesPCOverflow5", data, 6, 8, 8, 0, errNotEnough},
		{"SixBytesPCOverflow6", data, 6, 9, 9, 0, errNotEnough},
		{"SevenBytes", data, 7, 0, 7, 0x01020304050607, testerr.Nil()},
		{"SevenBytesPCOverflow1", data, 7, 3, 3, 0, errNotEnough},
		{"SevenBytesPCOverflow2", data, 7, 4, 4, 0, errNotEnough},
		{"SevenBytesPCOverflow3", data, 7, 5, 5, 0, errNotEnough},
		{"SevenBytesPCOverflow4", data, 7, 6, 6, 0, errNotEnough},
		{"SevenBytesPCOverflow5", data, 7, 7, 7, 0, errNotEnough},
		{"SevenBytesPCOverflow6", data, 7, 8, 8, 0, errNotEnough},
		{"SevenBytesPCOverflow7", data, 7, 9, 9, 0, errNotEnough},
		{"EightBytes", data, 8, 0, 8, 0x0102030405060708, testerr.Nil()},
		{"EightBytesPCOverflow1", data, 8, 2, 2, 0, errNotEnough},
		{"EightBytesPCOverflow2", data, 8, 3, 3, 0, errNotEnough},
		{"EightBytesPCOverflow3", data, 8, 4, 4, 0, errNotEnough},
		{"EightBytesPCOverflow4", data, 8, 5, 5, 0, errNotEnough},
		{"EightBytesPCOverflow5", data, 8, 6, 6, 0, errNotEnough},
		{"EightBytesPCOverflow6", data, 8, 7, 7, 0, errNotEnough},
		{"EightBytesPCOverflow7", data, 8, 8, 8, 0, errNotEnough},
		{"EightBytesPCOverflow8", data, 8, 9, 9, 0, errNotEnough},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			newpc, value, err := bytecode.FetchN(test.data, test.pc, test.size)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, newpc, "new pc value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}
