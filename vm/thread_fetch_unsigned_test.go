package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/vm"
	"github.com/tvarney/testerr"
)

var (
	nilerr   = testerr.Nil()
	errRead1 = testerr.Is(vm.FetchNotEnoughBytesError{Bytes: 1})
	errRead2 = testerr.Is(vm.FetchNotEnoughBytesError{Bytes: 2})
	errRead3 = testerr.Is(vm.FetchNotEnoughBytesError{Bytes: 3})
	errRead4 = testerr.Is(vm.FetchNotEnoughBytesError{Bytes: 4})
	errRead5 = testerr.Is(vm.FetchNotEnoughBytesError{Bytes: 5})
	errRead6 = testerr.Is(vm.FetchNotEnoughBytesError{Bytes: 6})
	errRead7 = testerr.Is(vm.FetchNotEnoughBytesError{Bytes: 7})
	errRead8 = testerr.Is(vm.FetchNotEnoughBytesError{Bytes: 8})
)

func TestThreadFetchUnsigned(t *testing.T) {
	t.Parallel()

	data := []uint8{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C,
		0x0D, 0x0E, 0x0F,
	}

	for _, test := range []struct {
		name     string
		data     []uint8
		size     int
		pc       int
		newpc    int
		expected uint64
		errval   testerr.ExpectedError
	}{
		{"FetchUnderflow0", data, 0, 1, 1, 0, testerr.Is(vm.ImmediateFetchSizeError{Bytes: 0})},
		{"FetchUnderflow1", data, -1, 1, 1, 0, testerr.Is(vm.ImmediateFetchSizeError{Bytes: -1})},
		{"FetchOverflow", data, 9, 1, 1, 0, testerr.Is(vm.ImmediateFetchSizeError{Bytes: 9})},
		{"Size1/Index0", data, 1, 0, 1, 0x01, nilerr},
		{"Size1/Index5", data, 1, 5, 6, 0x06, nilerr},
		{"Size1/PCOverflow", data, 1, 20, 20, 0, errRead1},
		{"Size1/PCUnderflow", data, 1, -1, -1, 0, errRead1},
		{"Size2/Index0", data, 2, 0, 2, 0x0102, nilerr},
		{"Size2/Index1", data, 2, 1, 3, 0x0203, nilerr},
		{"Size2/Index2", data, 2, 2, 4, 0x0304, nilerr},
		{"Size2/Index3", data, 2, 3, 5, 0x0405, nilerr},
		{"Size2/Index4", data, 2, 4, 6, 0x0506, nilerr},
		{"Size2/Index5", data, 2, 5, 7, 0x0607, nilerr},
		{"Size2/Index6", data, 2, 6, 8, 0x0708, nilerr},
		{"Size2/Index7", data, 2, 7, 9, 0x0809, nilerr},
		{"Size2/Index8", data, 2, 8, 10, 0x090A, nilerr},
		{"Size2/Index9", data, 2, 9, 11, 0x0A0B, nilerr},
		{"Size2/Index10", data, 2, 10, 12, 0x0B0C, nilerr},
		{"Size2/Index11", data, 2, 11, 13, 0x0C0D, nilerr},
		{"Size2/Index12", data, 2, 12, 14, 0x0D0E, nilerr},
		{"Size2/Index13", data, 2, 13, 15, 0x0E0F, nilerr},
		{"Size2/Index14", data, 2, 14, 14, 0, errRead2},
		{"Size2/Index15", data, 2, 15, 15, 0, errRead2},
		{"Size2/PCUnderflow", data, 2, -1, -1, 0, errRead2},
		{"Size3/Index0", data, 3, 0, 3, 0x010203, nilerr},
		{"Size3/Index1", data, 3, 1, 4, 0x020304, nilerr},
		{"Size3/Index2", data, 3, 2, 5, 0x030405, nilerr},
		{"Size3/Index3", data, 3, 3, 6, 0x040506, nilerr},
		{"Size3/Index4", data, 3, 4, 7, 0x050607, nilerr},
		{"Size3/Index5", data, 3, 5, 8, 0x060708, nilerr},
		{"Size3/Index6", data, 3, 6, 9, 0x070809, nilerr},
		{"Size3/Index7", data, 3, 7, 10, 0x08090A, nilerr},
		{"Size3/Index8", data, 3, 8, 11, 0x090A0B, nilerr},
		{"Size3/Index9", data, 3, 9, 12, 0x0A0B0C, nilerr},
		{"Size3/Index10", data, 3, 10, 13, 0x0B0C0D, nilerr},
		{"Size3/Index11", data, 3, 11, 14, 0x0C0D0E, nilerr},
		{"Size3/Index12", data, 3, 12, 15, 0x0D0E0F, nilerr},
		{"Size3/Index13", data, 3, 13, 13, 0, errRead3},
		{"Size3/Index14", data, 3, 14, 14, 0, errRead3},
		{"Size3/Index15", data, 3, 15, 15, 0, errRead3},
		{"Size3/PCUnderflow", data, 3, -1, -1, 0, errRead3},
		{"Size4/Index0", data, 4, 0, 4, 0x01020304, nilerr},
		{"Size4/Index1", data, 4, 1, 5, 0x02030405, nilerr},
		{"Size4/Index2", data, 4, 2, 6, 0x03040506, nilerr},
		{"Size4/Index3", data, 4, 3, 7, 0x04050607, nilerr},
		{"Size4/Index4", data, 4, 4, 8, 0x05060708, nilerr},
		{"Size4/Index5", data, 4, 5, 9, 0x06070809, nilerr},
		{"Size4/Index6", data, 4, 6, 10, 0x0708090A, nilerr},
		{"Size4/Index7", data, 4, 7, 11, 0x08090A0B, nilerr},
		{"Size4/Index8", data, 4, 8, 12, 0x090A0B0C, nilerr},
		{"Size4/Index9", data, 4, 9, 13, 0x0A0B0C0D, nilerr},
		{"Size4/Index10", data, 4, 10, 14, 0x0B0C0D0E, nilerr},
		{"Size4/Index11", data, 4, 11, 15, 0x0C0D0E0F, nilerr},
		{"Size4/Index12", data, 4, 12, 12, 0, errRead4},
		{"Size4/Index13", data, 4, 13, 13, 0, errRead4},
		{"Size4/Index14", data, 4, 14, 14, 0, errRead4},
		{"Size4/Index15", data, 4, 15, 15, 0, errRead4},
		{"Size4/PCUnderflow", data, 4, -1, -1, 0, errRead4},
		{"Size5/Index0", data, 5, 0, 5, 0x0102030405, nilerr},
		{"Size5/Index1", data, 5, 1, 6, 0x0203040506, nilerr},
		{"Size5/Index2", data, 5, 2, 7, 0x0304050607, nilerr},
		{"Size5/Index3", data, 5, 3, 8, 0x0405060708, nilerr},
		{"FiveBytes/Index4", data, 5, 4, 9, 0x0506070809, nilerr},
		{"Size5/Index5", data, 5, 5, 10, 0x060708090A, nilerr},
		{"Size5/Index6", data, 5, 6, 11, 0x0708090A0B, nilerr},
		{"Size5/Index7", data, 5, 7, 12, 0x08090A0B0C, nilerr},
		{"Size5/Index8", data, 5, 8, 13, 0x090A0B0C0D, nilerr},
		{"Size5/Index9", data, 5, 9, 14, 0x0A0B0C0D0E, nilerr},
		{"Size5/Index10", data, 5, 10, 15, 0x0B0C0D0E0F, nilerr},
		{"Size5/Index11", data, 5, 11, 11, 0, errRead5},
		{"Size5/Index12", data, 5, 12, 12, 0, errRead5},
		{"Size5/Index13", data, 5, 13, 13, 0, errRead5},
		{"Size5/Index14", data, 5, 14, 14, 0, errRead5},
		{"Size5/Index15", data, 5, 15, 15, 0, errRead5},
		{"Size5/PCUnderflow", data, 5, -1, -1, 0, errRead5},
		{"Size6/Index0", data, 6, 0, 6, 0x010203040506, nilerr},
		{"Size6/Index1", data, 6, 1, 7, 0x020304050607, nilerr},
		{"Size6/Index2", data, 6, 2, 8, 0x030405060708, nilerr},
		{"Size6/Index3", data, 6, 3, 9, 0x040506070809, nilerr},
		{"Size6/Index4", data, 6, 4, 10, 0x05060708090A, nilerr},
		{"Size6/Index5", data, 6, 5, 11, 0x060708090A0B, nilerr},
		{"Size6/Index6", data, 6, 6, 12, 0x0708090A0B0C, nilerr},
		{"Size6/Index7", data, 6, 7, 13, 0x08090A0B0C0D, nilerr},
		{"Size6/Index8", data, 6, 8, 14, 0x090A0B0C0D0E, nilerr},
		{"Size6/Index9", data, 6, 9, 15, 0x0A0B0C0D0E0F, nilerr},
		{"Size6/Index10", data, 6, 10, 10, 0, errRead6},
		{"Size6/Index11", data, 6, 11, 11, 0, errRead6},
		{"Size6/Index12", data, 6, 12, 12, 0, errRead6},
		{"Size6/Index13", data, 6, 13, 13, 0, errRead6},
		{"Size6/Index14", data, 6, 14, 14, 0, errRead6},
		{"Size6/Index15", data, 6, 15, 15, 0, errRead6},
		{"Size6/PCUnderflow", data, 6, -1, -1, 0, errRead6},
		{"Size7/Index0", data, 7, 0, 7, 0x01020304050607, nilerr},
		{"Size7/Index1", data, 7, 1, 8, 0x02030405060708, nilerr},
		{"Size7/Index2", data, 7, 2, 9, 0x03040506070809, nilerr},
		{"Size7/Index3", data, 7, 3, 10, 0x0405060708090A, nilerr},
		{"Size7/Index4", data, 7, 4, 11, 0x05060708090A0B, nilerr},
		{"Size7/Index5", data, 7, 5, 12, 0x060708090A0B0C, nilerr},
		{"Size7/Index6", data, 7, 6, 13, 0x0708090A0B0C0D, nilerr},
		{"Size7/Index7", data, 7, 7, 14, 0x08090A0B0C0D0E, nilerr},
		{"Size7/Index8", data, 7, 8, 15, 0x090A0B0C0D0E0F, nilerr},
		{"Size7/Index9", data, 7, 9, 9, 0, errRead7},
		{"Size7/Index10", data, 7, 10, 10, 0, errRead7},
		{"Size7/Index11", data, 7, 11, 11, 0, errRead7},
		{"Size7/Index12", data, 7, 12, 12, 0, errRead7},
		{"Size7/Index13", data, 7, 13, 13, 0, errRead7},
		{"Size7/Index14", data, 7, 14, 14, 0, errRead7},
		{"Size7/Index15", data, 7, 15, 15, 0, errRead7},
		{"Size7/PCUnderflow", data, 7, -1, -1, 0, errRead7},
		{"Size8/Index0", data, 8, 0, 8, 0x0102030405060708, nilerr},
		{"Size8/Index1", data, 8, 1, 9, 0x0203040506070809, nilerr},
		{"Size8/Index2", data, 8, 2, 10, 0x030405060708090A, nilerr},
		{"Size8/Index3", data, 8, 3, 11, 0x0405060708090A0B, nilerr},
		{"Size8/Index4", data, 8, 4, 12, 0x05060708090A0B0C, nilerr},
		{"Size8/Index5", data, 8, 5, 13, 0x060708090A0B0C0D, nilerr},
		{"Size8/Index6", data, 8, 6, 14, 0x0708090A0B0C0D0E, nilerr},
		{"Size8/Index7", data, 8, 7, 15, 0x08090A0B0C0D0E0F, nilerr},
		{"Size8/Index8", data, 8, 8, 8, 0, errRead8},
		{"Size8/Index9", data, 8, 9, 9, 0, errRead8},
		{"Size8/Index10", data, 8, 10, 10, 0, errRead8},
		{"Size8/Index11", data, 8, 11, 11, 0, errRead8},
		{"Size8/Index12", data, 8, 12, 12, 0, errRead8},
		{"Size8/Index13", data, 8, 13, 13, 0, errRead8},
		{"Size8/Index14", data, 8, 14, 14, 0, errRead8},
		{"Size8/Index15", data, 8, 15, 15, 0, errRead8},
		{"Size8/PCUnderflow", data, 8, -1, -1, 0, errRead8},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			value, err := th.FetchUnsigned(test.size)

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestThreadFetchU8(t *testing.T) {
	t.Parallel()

	data := []uint8{1, 2, 3, 4, 5}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected uint8
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 1, 1, nilerr},
		{"Index1", data, 1, 2, 2, nilerr},
		{"Index2", data, 2, 3, 3, nilerr},
		{"Index3", data, 3, 4, 4, nilerr},
		{"Index4", data, 4, 5, 5, nilerr},
		{"PCOverflow", data, 5, 5, 0, errRead1},
		{"PCUnderflow", data, -1, -1, 0, errRead1},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchU8()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchU16(t *testing.T) {
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
		{"Index0", data, 0, 2, 0x0102, nilerr},
		{"Index1", data, 1, 3, 0x0203, nilerr},
		{"Index2", data, 2, 4, 0x0304, nilerr},
		{"Index3", data, 3, 5, 0x0405, nilerr},
		{"PCOverflow1", data, 4, 4, 0, errRead2},
		{"PCOverflow2", data, 5, 5, 0, errRead2},
		{"PCUnderflow1", data, -1, -1, 0, errRead2},
		{"PCUnderflow2", data, -2, -2, 0, errRead2},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchU16()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchU24(t *testing.T) {
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
		{"Index0", data, 0, 3, 0x010203, nilerr},
		{"Index1", data, 1, 4, 0x020304, nilerr},
		{"Index2", data, 2, 5, 0x030405, nilerr},
		{"PCOverflow1", data, 3, 3, 0, errRead3},
		{"PCOverflow2", data, 4, 4, 0, errRead3},
		{"PCOverflow3", data, 5, 5, 0, errRead3},
		{"PCUnderflow", data, -1, -1, 0, errRead3},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			value, err := th.FetchU24()

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestThreadFetchU32(t *testing.T) {
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
		{"Index0", data, 0, 4, 0x01020304, nilerr},
		{"Index1", data, 1, 5, 0x02030405, nilerr},
		{"PCOverflow1", data, 2, 2, 0, errRead4},
		{"PCOverflow2", data, 3, 3, 0, errRead4},
		{"PCOverflow3", data, 4, 4, 0, errRead4},
		{"PCOverflow4", data, 5, 5, 0, errRead4},
		{"PCUnderflow", data, -1, -1, 0, errRead4},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: data, PC: test.pc}
			value, err := th.FetchU32()

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestThreadFetchU40(t *testing.T) {
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
		{"Index0", data, 0, 5, 0x0102030405, nilerr},
		{"Index1", data, 1, 6, 0x0203040506, nilerr},
		{"Index2", data, 2, 7, 0x0304050607, nilerr},
		{"Index3", data, 3, 8, 0x0405060708, nilerr},
		{"PCOverflow1", data, 4, 4, 0, errRead5},
		{"PCOverflow2", data, 5, 5, 0, errRead5},
		{"PCOverflow3", data, 6, 6, 0, errRead5},
		{"PCOverflow4", data, 7, 7, 0, errRead5},
		{"PCOverflow5", data, 8, 8, 0, errRead5},
		{"PCUnderflow", data, -1, -1, 0, errRead5},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			value, err := th.FetchU40()

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestThreadFetchU48(t *testing.T) {
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
		{"Index0", data, 0, 6, 0x010203040506, nilerr},
		{"Index1", data, 1, 7, 0x020304050607, nilerr},
		{"Index2", data, 2, 8, 0x030405060708, nilerr},
		{"PCOverflow1", data, 3, 3, 0, errRead6},
		{"PCOverflow2", data, 4, 4, 0, errRead6},
		{"PCOverflow3", data, 5, 5, 0, errRead6},
		{"PCOverflow4", data, 6, 6, 0, errRead6},
		{"PCOverflow5", data, 7, 7, 0, errRead6},
		{"PCOverflow6", data, 8, 8, 0, errRead6},
		{"PCUnderflow", data, -1, -1, 0, errRead6},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			value, err := th.FetchU48()

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestThreadFetchU56(t *testing.T) {
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
		{"Index0", data, 0, 7, 0x01020304050607, nilerr},
		{"Index1", data, 1, 8, 0x02030405060708, nilerr},
		{"PCOverflow1", data, 2, 2, 0, errRead7},
		{"PCOverflow2", data, 3, 3, 0, errRead7},
		{"PCOverflow3", data, 4, 4, 0, errRead7},
		{"PCOverflow4", data, 5, 5, 0, errRead7},
		{"PCOverflow5", data, 6, 6, 0, errRead7},
		{"PCOverflow6", data, 7, 7, 0, errRead7},
		{"PCOverflow7", data, 8, 8, 0, errRead7},
		{"PCUnderflow", data, -1, -1, 0, errRead7},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			value, err := th.FetchU56()

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}

func TestThreadFetchU64(t *testing.T) {
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
		{"Index0", data, 0, 8, 0x0102030405060708, nilerr},
		{"PCOverflow1", data, 1, 1, 0, errRead8},
		{"PCOverflow2", data, 2, 2, 0, errRead8},
		{"PCOverflow3", data, 3, 3, 0, errRead8},
		{"PCOverflow4", data, 4, 4, 0, errRead8},
		{"PCOverflow5", data, 5, 5, 0, errRead8},
		{"PCOverflow6", data, 6, 6, 0, errRead8},
		{"PCOverflow7", data, 7, 7, 0, errRead8},
		{"PCOverflow8", data, 8, 8, 0, errRead8},
		{"PCUnderflow", data, -1, -1, 0, errRead8},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			value, err := th.FetchU64()

			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, value, "returned value")
		})
	}
}
