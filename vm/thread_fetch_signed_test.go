package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/vm"
	"github.com/tvarney/testerr"
)

const (
	negBaseI8  = -0x80
	negBaseI16 = -0x8000
	negBaseI24 = -0x800000
	negBaseI32 = -0x80000000
	negBaseI40 = -0x8000000000
	negBaseI48 = -0x800000000000
	negBaseI56 = -0x80000000000000
	negBaseI64 = -0x8000000000000000
)

func TestThreadFetchSigned(t *testing.T) {
	t.Parallel()

	data := []uint8{
		0x80, 0x01, 0x82, 0x03, 0x84, 0x05, 0x86, 0x07, 0x88, 0x09, 0x8A,
		0x0B, 0x8C, 0x0D, 0x8E, 0x0F,
	}

	for _, test := range []struct {
		name     string
		data     []uint8
		size     int
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"FetchUnderflow0", data, 0, 0, 0, 0, testerr.Is(vm.ImmediateFetchSizeError{Bytes: 0})},
		{"FetchUnderflow1", data, -1, 4, 4, 0, testerr.Is(vm.ImmediateFetchSizeError{Bytes: -1})},
		{"FetchOverflow", data, 9, 1, 1, 0, testerr.Is(vm.ImmediateFetchSizeError{Bytes: 9})},
		{"Size1/Index0", data, 1, 0, 1, negBaseI8 + 0x00, nilerr},
		{"Size1/Index1", data, 1, 1, 2, 0x01, nilerr},
		{"Size1/Index2", data, 1, 2, 3, negBaseI8 + 0x02, nilerr},
		{"Size1/Index3", data, 1, 3, 4, 0x03, nilerr},
		{"Size1/Index4", data, 1, 4, 5, negBaseI8 + 0x04, nilerr},
		{"Size1/Index5", data, 1, 5, 6, 0x05, nilerr},
		{"Size1/Index6", data, 1, 6, 7, negBaseI8 + 0x06, nilerr},
		{"Size1/Index7", data, 1, 7, 8, 0x07, nilerr},
		{"Size1/Index8", data, 1, 8, 9, negBaseI8 + 0x08, nilerr},
		{"Size1/Index9", data, 1, 9, 10, 0x09, nilerr},
		{"Size1/Index10", data, 1, 10, 11, negBaseI8 + 0x0A, nilerr},
		{"Size1/Index11", data, 1, 11, 12, 0x0B, nilerr},
		{"Size1/Index12", data, 1, 12, 13, negBaseI8 + 0x0C, nilerr},
		{"Size1/Index13", data, 1, 13, 14, 0x0D, nilerr},
		{"Size1/Index14", data, 1, 14, 15, negBaseI8 + 0x0E, nilerr},
		{"Size1/Index15", data, 1, 15, 16, 0x0F, nilerr},
		{"Size1/Index16", data, 1, 16, 16, 0, errRead1},
		{"Size1/PCUnderflow", data, 1, -1, -1, 0, errRead1},
		{"Size2/Index0", data, 2, 0, 2, negBaseI16 + 0x0001, nilerr},
		{"Size2/Index1", data, 2, 1, 3, 0x0182, nilerr},
		{"Size2/Index2", data, 2, 2, 4, negBaseI16 + 0x0203, nilerr},
		{"Size2/Index3", data, 2, 3, 5, 0x0384, nilerr},
		{"Size2/Index4", data, 2, 4, 6, negBaseI16 + 0x0405, nilerr},
		{"Size2/Index5", data, 2, 5, 7, 0x0586, nilerr},
		{"Size2/Index6", data, 2, 6, 8, negBaseI16 + 0x0607, nilerr},
		{"Size2/Index7", data, 2, 7, 9, 0x0788, nilerr},
		{"Size2/Index8", data, 2, 8, 10, negBaseI16 + 0x0809, nilerr},
		{"Size2/Index9", data, 2, 9, 11, 0x098A, nilerr},
		{"Size2/Index10", data, 2, 10, 12, negBaseI16 + 0x0A0B, nilerr},
		{"Size2/Index11", data, 2, 11, 13, 0x0B8C, nilerr},
		{"Size2/Index12", data, 2, 12, 14, negBaseI16 + 0x0C0D, nilerr},
		{"Size2/Index13", data, 2, 13, 15, 0x0D8E, nilerr},
		{"Size2/Index14", data, 2, 14, 16, negBaseI16 + 0x0E0F, nilerr},
		{"Size2/Index15", data, 2, 15, 15, 0, errRead2},
		{"Size2/Index16", data, 2, 16, 16, 0, errRead2},
		{"Size2/PCUnderflow", data, 2, -1, -1, 0, errRead2},
		{"Size3/Index0", data, 3, 0, 3, negBaseI24 + 0x000182, nilerr},
		{"Size3/Index1", data, 3, 1, 4, 0x018203, nilerr},
		{"Size3/Index2", data, 3, 2, 5, negBaseI24 + 0x020384, nilerr},
		{"Size3/Index3", data, 3, 3, 6, 0x038405, nilerr},
		{"Size3/Index4", data, 3, 4, 7, negBaseI24 + 0x040586, nilerr},
		{"Size3/Index5", data, 3, 5, 8, 0x058607, nilerr},
		{"Size3/Index6", data, 3, 6, 9, negBaseI24 + 0x060788, nilerr},
		{"Size3/Index7", data, 3, 7, 10, 0x078809, nilerr},
		{"Size3/Index8", data, 3, 8, 11, negBaseI24 + 0x08098A, nilerr},
		{"Size3/Index9", data, 3, 9, 12, 0x098A0B, nilerr},
		{"Size3/Index10", data, 3, 10, 13, negBaseI24 + 0x0A0B8C, nilerr},
		{"Size3/Index11", data, 3, 11, 14, 0x0B8C0D, nilerr},
		{"Size3/Index12", data, 3, 12, 15, negBaseI24 + 0x0C0D8E, nilerr},
		{"Size3/Index13", data, 3, 13, 16, 0x0D8E0F, nilerr},
		{"Size3/Index14", data, 3, 14, 14, 0, errRead3},
		{"Size3/Index15", data, 3, 15, 15, 0, errRead3},
		{"Size3/Index16", data, 3, 16, 16, 0, errRead3},
		{"Size3/PCUnderflow", data, 3, -1, -1, 0, errRead3},
		{"Size4/Index0", data, 4, 0, 4, negBaseI32 + 0x00018203, nilerr},
		{"Size4/Index1", data, 4, 1, 5, 0x01820384, nilerr},
		{"Size4/Index2", data, 4, 2, 6, negBaseI32 + 0x02038405, nilerr},
		{"Size4/Index3", data, 4, 3, 7, 0x03840586, nilerr},
		{"Size4/Index4", data, 4, 4, 8, negBaseI32 + 0x04058607, nilerr},
		{"Size4/Index5", data, 4, 5, 9, 0x05860788, nilerr},
		{"Size4/Index6", data, 4, 6, 10, negBaseI32 + 0x06078809, nilerr},
		{"Size4/Index7", data, 4, 7, 11, 0x0788098A, nilerr},
		{"Size4/Index8", data, 4, 8, 12, negBaseI32 + 0x08098A0B, nilerr},
		{"Size4/Index9", data, 4, 9, 13, 0x098A0B8C, nilerr},
		{"Size4/Index10", data, 4, 10, 14, negBaseI32 + 0x0A0B8C0D, nilerr},
		{"Size4/Index11", data, 4, 11, 15, 0x0B8C0D8E, nilerr},
		{"Size4/Index12", data, 4, 12, 16, negBaseI32 + 0x0C0D8E0F, nilerr},
		{"Size4/Index13", data, 4, 13, 13, 0, errRead4},
		{"Size4/Index14", data, 4, 14, 14, 0, errRead4},
		{"Size4/Index15", data, 4, 15, 15, 0, errRead4},
		{"Size4/Index16", data, 4, 16, 16, 0, errRead4},
		{"Size4/PCUnderflow", data, 4, -1, -1, 0, errRead4},
		{"Size5/Index0", data, 5, 0, 5, negBaseI40 + 0x0001820384, nilerr},
		{"Size5/Index1", data, 5, 1, 6, 0x0182038405, nilerr},
		{"Size5/Index2", data, 5, 2, 7, negBaseI40 + 0x0203840586, nilerr},
		{"Size5/Index3", data, 5, 3, 8, 0x0384058607, nilerr},
		{"Size5/Index4", data, 5, 4, 9, negBaseI40 + 0x0405860788, nilerr},
		{"Size5/Index5", data, 5, 5, 10, 0x0586078809, nilerr},
		{"Size5/Index6", data, 5, 6, 11, negBaseI40 + 0x060788098A, nilerr},
		{"Size5/Index7", data, 5, 7, 12, 0x0788098A0B, nilerr},
		{"Size5/Index8", data, 5, 8, 13, negBaseI40 + 0x08098A0B8C, nilerr},
		{"Size5/Index9", data, 5, 9, 14, 0x098A0B8C0D, nilerr},
		{"Size5/Index10", data, 5, 10, 15, negBaseI40 + 0x0A0B8C0D8E, nilerr},
		{"Size5/Index11", data, 5, 11, 16, 0x0B8C0D8E0F, nilerr},
		{"Size5/Index12", data, 5, 12, 12, 0, errRead5},
		{"Size5/Index13", data, 5, 13, 13, 0, errRead5},
		{"Size5/Index14", data, 5, 14, 14, 0, errRead5},
		{"Size5/Index15", data, 5, 15, 15, 0, errRead5},
		{"Size5/Index16", data, 5, 16, 16, 0, errRead5},
		{"Size5/PCUnderflow", data, 5, -1, -1, 0, errRead5},
		{"Size6/Index0", data, 6, 0, 6, negBaseI48 + 0x000182038405, nilerr},
		{"Size6/Index1", data, 6, 1, 7, 0x018203840586, nilerr},
		{"Size6/Index2", data, 6, 2, 8, negBaseI48 + 0x020384058607, nilerr},
		{"Size6/Index3", data, 6, 3, 9, 0x038405860788, nilerr},
		{"Size6/Index4", data, 6, 4, 10, negBaseI48 + 0x040586078809, nilerr},
		{"Size6/Index5", data, 6, 5, 11, 0x05860788098A, nilerr},
		{"Size6/Index6", data, 6, 6, 12, negBaseI48 + 0x060788098A0B, nilerr},
		{"Size6/Index7", data, 6, 7, 13, 0x0788098A0B8C, nilerr},
		{"Size6/Index8", data, 6, 8, 14, negBaseI48 + 0x08098A0B8C0D, nilerr},
		{"Size6/Index9", data, 6, 9, 15, 0x098A0B8C0D8E, nilerr},
		{"Size6/Index10", data, 6, 10, 16, negBaseI48 + 0x0A0B8C0D8E0F, nilerr},
		{"Size6/Index11", data, 6, 11, 11, 0, errRead6},
		{"Size6/Index12", data, 6, 12, 12, 0, errRead6},
		{"Size6/Index13", data, 6, 13, 13, 0, errRead6},
		{"Size6/Index14", data, 6, 14, 14, 0, errRead6},
		{"Size6/Index15", data, 6, 15, 15, 0, errRead6},
		{"Size6/Index16", data, 6, 16, 16, 0, errRead6},
		{"Size6/PCUnderflow", data, 6, -1, -1, 0, errRead6},
		{"Size7/Index0", data, 7, 0, 7, negBaseI56 + 0x00018203840586, nilerr},
		{"Size7/Index1", data, 7, 1, 8, 0x01820384058607, nilerr},
		{"Size7/Index2", data, 7, 2, 9, negBaseI56 + 0x02038405860788, nilerr},
		{"Size7/Index3", data, 7, 3, 10, 0x03840586078809, nilerr},
		{"Size7/Index4", data, 7, 4, 11, negBaseI56 + 0x0405860788098A, nilerr},
		{"Size7/Index5", data, 7, 5, 12, 0x05860788098A0B, nilerr},
		{"Size7/Index6", data, 7, 6, 13, negBaseI56 + 0x060788098A0B8C, nilerr},
		{"Size7/Index7", data, 7, 7, 14, 0x0788098A0B8C0D, nilerr},
		{"Size7/Index8", data, 7, 8, 15, negBaseI56 + 0x08098A0B8C0D8E, nilerr},
		{"Size7/Index9", data, 7, 9, 16, 0x098A0B8C0D8E0F, nilerr},
		{"Size7/Index10", data, 7, 10, 10, 0, errRead7},
		{"Size7/Index11", data, 7, 11, 11, 0, errRead7},
		{"Size7/Index12", data, 7, 12, 12, 0, errRead7},
		{"Size7/Index13", data, 7, 13, 13, 0, errRead7},
		{"Size7/Index14", data, 7, 14, 14, 0, errRead7},
		{"Size7/Index15", data, 7, 15, 15, 0, errRead7},
		{"Size7/Index16", data, 7, 16, 16, 0, errRead7},
		{"Size7/PCUnderflow", data, 7, -1, -1, 0, errRead7},
		{"Size8/Index0", data, 8, 0, 8, negBaseI64 + 0x0001820384058607, nilerr},
		{"Size8/Index1", data, 8, 1, 9, 0x0182038405860788, nilerr},
		{"Size8/Index2", data, 8, 2, 10, negBaseI64 + 0x0203840586078809, nilerr},
		{"Size8/Index3", data, 8, 3, 11, 0x038405860788098A, nilerr},
		{"Size8/Index4", data, 8, 4, 12, negBaseI64 + 0x0405860788098A0B, nilerr},
		{"Size8/Index5", data, 8, 5, 13, 0x05860788098A0B8C, nilerr},
		{"Size8/Index6", data, 8, 6, 14, negBaseI64 + 0x060788098A0B8C0D, nilerr},
		{"Size8/Index7", data, 8, 7, 15, 0x0788098A0B8C0D8E, nilerr},
		{"Size8/Index8", data, 8, 8, 16, negBaseI64 + 0x08098A0B8C0D8E0F, nilerr},
		{"Size8/Index9", data, 8, 9, 9, 0, errRead8},
		{"Size8/Index10", data, 8, 10, 10, 0, errRead8},
		{"Size8/Index11", data, 8, 11, 11, 0, errRead8},
		{"Size8/Index12", data, 8, 12, 12, 0, errRead8},
		{"Size8/Index13", data, 8, 13, 13, 0, errRead8},
		{"Size8/Index14", data, 8, 14, 14, 0, errRead8},
		{"Size8/Index15", data, 8, 15, 15, 0, errRead8},
		{"Size8/Index16", data, 8, 16, 16, 0, errRead8},
		{"Size8/PCUnderflow", data, 8, -1, -1, 0, errRead8},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchSigned(test.size)
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchI8(t *testing.T) {
	t.Parallel()

	data := []uint8{0x80, 0x01, 0x82, 0x03, 0x84}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int8
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 1, -128, nilerr},
		{"Index1", data, 1, 2, 1, nilerr},
		{"Index2", data, 2, 3, -126, nilerr},
		{"Index3", data, 3, 4, 3, nilerr},
		{"Index4", data, 4, 5, -124, nilerr},
		{"Index5", data, 5, 5, 0, errRead1},
		{"PCUnderflow", data, -1, -1, 0, errRead1},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchI8()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchI16(t *testing.T) {
	t.Parallel()

	data := []uint8{0x80, 0x01, 0x82, 0x03, 0x84}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int16
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 2, negBaseI16 + 0x0001, nilerr},
		{"Index1", data, 1, 3, 0x0182, nilerr},
		{"Index2", data, 2, 4, negBaseI16 + 0x0203, nilerr},
		{"Index3", data, 3, 5, 0x0384, nilerr},
		{"Index4", data, 4, 4, 0, errRead2},
		{"Index5", data, 5, 5, 0, errRead2},
		{"PCUnderflow", data, -1, -1, 0, errRead2},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchI16()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchI24(t *testing.T) {
	t.Parallel()

	data := []uint8{0x80, 0x01, 0x82, 0x03, 0x84, 0x05, 0x86, 0x07}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int32
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 3, negBaseI24 + 0x000182, nilerr},
		{"Index1", data, 1, 4, 0x018203, nilerr},
		{"Index2", data, 2, 5, negBaseI24 + 0x020384, nilerr},
		{"Index3", data, 3, 6, 0x038405, nilerr},
		{"Index4", data, 4, 7, negBaseI24 + 0x040586, nilerr},
		{"Index5", data, 5, 8, 0x058607, nilerr},
		{"Index6", data, 6, 6, 0, errRead3},
		{"Index7", data, 7, 7, 0, errRead3},
		{"Index8", data, 8, 8, 0, errRead3},
		{"PCUnderflow", data, -1, -1, 0, errRead3},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchI24()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchI32(t *testing.T) {
	t.Parallel()

	data := []uint8{0x80, 0x01, 0x82, 0x03, 0x84, 0x05, 0x86, 0x07}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int32
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 4, negBaseI32 + 0x00018203, nilerr},
		{"Index1", data, 1, 5, 0x01820384, nilerr},
		{"Index2", data, 2, 6, negBaseI32 + 0x02038405, nilerr},
		{"Index3", data, 3, 7, 0x03840586, nilerr},
		{"Index4", data, 4, 8, negBaseI32 + 0x04058607, nilerr},
		{"Index5", data, 5, 5, 0, errRead4},
		{"Index6", data, 6, 6, 0, errRead4},
		{"Index7", data, 7, 7, 0, errRead4},
		{"Index8", data, 8, 8, 0, errRead4},
		{"PCUnderflow", data, -1, -1, 0, errRead4},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchI32()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchI40(t *testing.T) {
	t.Parallel()

	data := []uint8{
		0x80, 0x01, 0x82, 0x03, 0x84, 0x05, 0x86, 0x07, 0x88, 0x09, 0x8A, 0x0B,
	}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 5, negBaseI40 + 0x0001820384, nilerr},
		{"Index1", data, 1, 6, 0x0182038405, nilerr},
		{"Index2", data, 2, 7, negBaseI40 + 0x0203840586, nilerr},
		{"Index3", data, 3, 8, 0x0384058607, nilerr},
		{"Index4", data, 4, 9, negBaseI40 + 0x0405860788, nilerr},
		{"Index5", data, 5, 10, 0x0586078809, nilerr},
		{"Index6", data, 6, 11, negBaseI40 + 0x060788098A, nilerr},
		{"Index7", data, 7, 12, 0x0788098A0B, nilerr},
		{"Index8", data, 8, 8, 0, errRead5},
		{"Index9", data, 9, 9, 0, errRead5},
		{"Index10", data, 10, 10, 0, errRead5},
		{"Index11", data, 11, 11, 0, errRead5},
		{"Index12", data, 12, 12, 0, errRead5},
		{"PCUnderflow", data, -1, -1, 0, errRead5},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchI40()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchI48(t *testing.T) {
	t.Parallel()

	data := []uint8{
		0x80, 0x01, 0x82, 0x03, 0x84, 0x05, 0x86, 0x07, 0x88, 0x09, 0x8A, 0x0B,
	}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 6, negBaseI48 + 0x000182038405, nilerr},
		{"Index1", data, 1, 7, 0x018203840586, nilerr},
		{"Index2", data, 2, 8, negBaseI48 + 0x020384058607, nilerr},
		{"Index3", data, 3, 9, 0x038405860788, nilerr},
		{"Index4", data, 4, 10, negBaseI48 + 0x040586078809, nilerr},
		{"Index5", data, 5, 11, 0x05860788098A, nilerr},
		{"Index6", data, 6, 12, negBaseI48 + 0x060788098A0B, nilerr},
		{"Index7", data, 7, 7, 0, errRead6},
		{"Index8", data, 8, 8, 0, errRead6},
		{"Index9", data, 9, 9, 0, errRead6},
		{"Index10", data, 10, 10, 0, errRead6},
		{"Index11", data, 11, 11, 0, errRead6},
		{"Index12", data, 12, 12, 0, errRead6},
		{"PCUnderflow", data, -1, -1, 0, errRead6},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchI48()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchI56(t *testing.T) {
	t.Parallel()

	data := []uint8{
		0x80, 0x01, 0x82, 0x03, 0x84, 0x05, 0x86, 0x07, 0x88, 0x09, 0x8A, 0x0B,
	}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 7, negBaseI56 + 0x00018203840586, nilerr},
		{"Index1", data, 1, 8, 0x01820384058607, nilerr},
		{"Index2", data, 2, 9, negBaseI56 + 0x02038405860788, nilerr},
		{"Index3", data, 3, 10, 0x03840586078809, nilerr},
		{"Index4", data, 4, 11, negBaseI56 + 0x0405860788098A, nilerr},
		{"Index5", data, 5, 12, 0x05860788098A0B, nilerr},
		{"Index6", data, 6, 6, 0, errRead7},
		{"Index7", data, 7, 7, 0, errRead7},
		{"Index8", data, 8, 8, 0, errRead7},
		{"Index9", data, 9, 9, 0, errRead7},
		{"Index10", data, 10, 10, 0, errRead7},
		{"Index11", data, 11, 11, 0, errRead7},
		{"Index12", data, 12, 12, 0, errRead7},
		{"PCUnderflow", data, -1, -1, 0, errRead7},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchI56()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}

func TestThreadFetchI64(t *testing.T) {
	t.Parallel()

	data := []uint8{
		0x80, 0x01, 0x82, 0x03, 0x84, 0x05, 0x86, 0x07, 0x88, 0x09, 0x8A, 0x0B,
	}

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 8, negBaseI64 + 0x0001820384058607, nilerr},
		{"Index1", data, 1, 9, 0x0182038405860788, nilerr},
		{"Index2", data, 2, 10, negBaseI64 + 0x0203840586078809, nilerr},
		{"Index3", data, 3, 11, 0x038405860788098A, nilerr},
		{"Index4", data, 4, 12, negBaseI64 + 0x0405860788098A0B, nilerr},
		{"Index5", data, 5, 5, 0, errRead8},
		{"Index6", data, 6, 6, 0, errRead8},
		{"Index7", data, 7, 7, 0, errRead8},
		{"Index8", data, 8, 8, 0, errRead8},
		{"Index9", data, 9, 9, 0, errRead8},
		{"Index10", data, 10, 10, 0, errRead8},
		{"Index11", data, 11, 11, 0, errRead8},
		{"Index12", data, 12, 12, 0, errRead8},
		{"PCUnderflow", data, -1, -1, 0, errRead8},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			th := &vm.Thread{Data: test.data, PC: test.pc}
			v, err := th.FetchI64()
			test.errval.Require(t, err)
			require.Equal(t, test.newpc, th.PC, "new PC value")
			require.Equal(t, test.expected, v, "returned value")
		})
	}
}
