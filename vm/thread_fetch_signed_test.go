package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/vm"
	"github.com/tvarney/testerr"
)

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

	const negBase = -0x8000

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int16
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 2, negBase + 0x0001, nilerr},
		{"Index1", data, 1, 3, 0x0182, nilerr},
		{"Index2", data, 2, 4, negBase + 0x0203, nilerr},
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

	const negBase = -0x800000

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int32
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 3, negBase + 0x000182, nilerr},
		{"Index1", data, 1, 4, 0x018203, nilerr},
		{"Index2", data, 2, 5, negBase + 0x020384, nilerr},
		{"Index3", data, 3, 6, 0x038405, nilerr},
		{"Index4", data, 4, 7, negBase + 0x040586, nilerr},
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

	const negBase = -0x80000000

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int32
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 4, negBase + 0x00018203, nilerr},
		{"Index1", data, 1, 5, 0x01820384, nilerr},
		{"Index2", data, 2, 6, negBase + 0x02038405, nilerr},
		{"Index3", data, 3, 7, 0x03840586, nilerr},
		{"Index4", data, 4, 8, negBase + 0x04058607, nilerr},
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

	const negBase = -0x8000000000

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 5, negBase + 0x0001820384, nilerr},
		{"Index1", data, 1, 6, 0x0182038405, nilerr},
		{"Index2", data, 2, 7, negBase + 0x0203840586, nilerr},
		{"Index3", data, 3, 8, 0x0384058607, nilerr},
		{"Index4", data, 4, 9, negBase + 0x0405860788, nilerr},
		{"Index5", data, 5, 10, 0x0586078809, nilerr},
		{"Index6", data, 6, 11, negBase + 0x060788098A, nilerr},
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

	const negBase = -0x800000000000

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 6, negBase + 0x000182038405, nilerr},
		{"Index1", data, 1, 7, 0x018203840586, nilerr},
		{"Index2", data, 2, 8, negBase + 0x020384058607, nilerr},
		{"Index3", data, 3, 9, 0x038405860788, nilerr},
		{"Index4", data, 4, 10, negBase + 0x040586078809, nilerr},
		{"Index5", data, 5, 11, 0x05860788098A, nilerr},
		{"Index6", data, 6, 12, negBase + 0x060788098A0B, nilerr},
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

	const negBase = -0x80000000000000

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 7, negBase + 0x00018203840586, nilerr},
		{"Index1", data, 1, 8, 0x01820384058607, nilerr},
		{"Index2", data, 2, 9, negBase + 0x02038405860788, nilerr},
		{"Index3", data, 3, 10, 0x03840586078809, nilerr},
		{"Index4", data, 4, 11, negBase + 0x0405860788098A, nilerr},
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

	const negBase = -0x8000000000000000

	for _, test := range []struct {
		name     string
		data     []uint8
		pc       int
		newpc    int
		expected int64
		errval   testerr.ExpectedError
	}{
		{"Index0", data, 0, 8, negBase + 0x0001820384058607, nilerr},
		{"Index1", data, 1, 9, 0x0182038405860788, nilerr},
		{"Index2", data, 2, 10, negBase + 0x0203840586078809, nilerr},
		{"Index3", data, 3, 11, 0x038405860788098A, nilerr},
		{"Index4", data, 4, 12, negBase + 0x0405860788098A0B, nilerr},
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
