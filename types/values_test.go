package types_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/types"
	"github.com/tvarney/illvm/types/typeid"
	"github.com/tvarney/testerr"
)

func TestValue(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name   string
		v      types.Value
		id     typeid.ID
		size   int
		upcast typeid.ID
	}{
		{"Uint8", types.Uint8(10), typeid.Uint8, 1, typeid.Uint64},
		{"Uint16", types.Uint16(10), typeid.Uint16, 2, typeid.Uint64},
		{"Uint32", types.Uint32(10), typeid.Uint32, 4, typeid.Uint64},
		{"Uint64", types.Uint64(10), typeid.Uint64, 8, typeid.Uint64},
		{"Int8", types.Int8(10), typeid.Int8, 1, typeid.Int64},
		{"Int16", types.Int16(10), typeid.Int16, 2, typeid.Int64},
		{"Int32", types.Int32(10), typeid.Int32, 4, typeid.Int64},
		{"Int64", types.Int64(10), typeid.Int64, 8, typeid.Int64},
		{"Float32", types.Float32(1.0), typeid.Float32, 4, typeid.Float64},
		{"Float64", types.Float64(1.0), typeid.Float64, 8, typeid.Float64},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.id, test.v.ID())
			assert.Equal(t, test.size, test.v.Size())
			assert.Equal(t, test.upcast, test.v.Upcast().ID())
		})
	}
}

func TestDowncast(t *testing.T) {
	t.Parallel()

	const (
		Uint8   = typeid.Uint8
		Uint16  = typeid.Uint16
		Uint32  = typeid.Uint32
		Uint64  = typeid.Uint64
		Int8    = typeid.Int8
		Int16   = typeid.Int16
		Int32   = typeid.Int32
		Int64   = typeid.Int64
		Float32 = typeid.Float32
		Float64 = typeid.Float64
	)

	nilErr := testerr.Nil()
	castErr := func(from, to typeid.ID) testerr.ExpectedError {
		return testerr.Is(types.CastError{From: from, To: to})
	}

	for _, test := range []struct {
		name     string
		value    types.StackValue
		to       typeid.ID
		expected types.Value
		err      testerr.ExpectedError
	}{
		// Uint64: Uint8
		{"Uint64/Uint8/Normal", u64(100), Uint8, u8(100), nilErr},
		{"Uint64/Uint8/Overflow", u64(257), Uint8, u8(1), nilErr},
		// Uint64: Uint16
		{"Uint64/Uint16/Normal", u64(1024), Uint16, u16(1024), nilErr},
		{"Uint64/Uint16/Overflow", u64(65537), Uint16, u16(1), nilErr},
		// Uint64: Uint32
		{"Uint64/Uint32/Normal", u64(100000), Uint32, u32(100000), nilErr},
		// Uint64: Uint64 (self)
		{"Uint64/Uint64", u64(50000000000), Uint64, u64(50000000000), nilErr},
		// Uint64: Int8
		{"Uint64/Int8/Normal", u64(12), Int8, i8(12), nilErr},
		// Uint64: Int16
		{"Uint64/Int16/Normal", u64(2048), Int16, i16(2048), nilErr},
		// Uint64: Int32
		{"Uint64/Int32/Normal", u64(101000), Int32, i32(101000), nilErr},
		// Uint64: Int64
		{"Uint64/Int64/Normal", u64(98765432), Int64, i64(98765432), nilErr},
		// Uint64: Float32
		{"Uint64/Float32", u64(32), Float32, f32(32.0), nilErr},
		// Uint64: Float64
		{"Uint64/Float64", u64(999999), Float64, f64(999999.0), nilErr},
		// Uint64: Other (errors)
		{"Uint64/Boolean", u64(0), typeid.Boolean, nil, castErr(Uint64, typeid.Boolean)},
		{"Uint64/String", u64(0), typeid.String, nil, castErr(Uint64, typeid.String)},
		{"Uint64/List", u64(0), typeid.String, nil, castErr(Uint64, typeid.String)},
		{"Uint64/Map", u64(0), typeid.String, nil, castErr(Uint64, typeid.String)},
		{"Uint64/Struct", u64(0), typeid.String, nil, castErr(Uint64, typeid.String)},
		{"Uint64/Class", u64(0), typeid.String, nil, castErr(Uint64, typeid.String)},
		{"Uint64/Function", u64(0), typeid.String, nil, castErr(Uint64, typeid.String)},
		{"Uint64/Method", u64(0), typeid.String, nil, castErr(Uint64, typeid.String)},

		// Int64: Uint8
		{"Int64/Uint8/Normal", i64(12), Uint8, u8(12), nilErr},
		{"Int64/Uint8/Underflow", i64(-2), Uint8, u8(254), nilErr},
		{"Int64/Uint8/Overflow", i64(300), Uint8, u8(44), nilErr},
		// Int64: Uint16
		{"Int64/Uint16/Normal", i64(1030), Uint16, u16(1030), nilErr},
		// Int64: Uint32
		{"Int64/Uint32/Normal", i64(987123), Uint32, u32(987123), nilErr},
		// Int64: Uint64
		{"Int64/Uint64/Normal", i64(99), Uint64, u64(99), nilErr},
		// Int64: Int8
		{"Int64/Int8/Normal", i64(100), Int8, i8(100), nilErr},
		// Int64: Int16
		{"Int64/Int16/Normal", i64(500), Int16, i16(500), nilErr},
		// Int64: Int32
		{"Int64/Int32/Normal", i64(500000), Int32, i32(500000), nilErr},
		// Int64: Int64 (self)
		{"Int64/Int64", i64(1234567890), Int64, i64(1234567890), nilErr},
		// Int64: Float32
		{"Int64/Float32", i64(123), Float32, f32(123.0), nilErr},
		// Int64: Float64
		{"Int64/Float64", i64(321), Float64, f64(321.0), nilErr},
		// Int64: Other (errors)
		{"Int64/Boolean", i64(0), typeid.Boolean, nil, castErr(Int64, typeid.Boolean)},
		{"Int64/String", i64(0), typeid.String, nil, castErr(Int64, typeid.String)},
		{"Int64/List", i64(0), typeid.String, nil, castErr(Int64, typeid.String)},
		{"Int64/Map", i64(0), typeid.String, nil, castErr(Int64, typeid.String)},
		{"Int64/Struct", i64(0), typeid.String, nil, castErr(Int64, typeid.String)},
		{"Int64/Class", i64(0), typeid.String, nil, castErr(Int64, typeid.String)},
		{"Int64/Function", i64(0), typeid.String, nil, castErr(Int64, typeid.String)},
		{"Int64/Method", i64(0), typeid.String, nil, castErr(Int64, typeid.String)},
		// Float64: Uint8
		{"Float64/Uint8/Negative", f64(-1.0), Uint8, u8(255), nilErr},
		{"Float64/Uint8/Positive", f64(1.0), Uint8, u8(1), nilErr},
		{"Float64/Uint8/Fractional", f64(1.5), Uint8, u8(1), nilErr},
		{"Float64/Float64/Uint8/Overflow", f64(300), Uint8, u8(44), nilErr},
		// Float64: Uint16
		{"Float64/Uint16/Negative", f64(-1.0), Uint16, u16(65535), nilErr},
		{"Float64/Uint16/Positive", f64(13.0), Uint16, u16(13), nilErr},
		{"Float64/Uint16/Fractional", f64(82.921), Uint16, u16(82), nilErr},
		{"Float64/Uint16/Overflow", f64(99999), Uint16, u16(99999 - math.MaxUint16 - 1), nilErr},
		// Float64: Uint32
		{"Float64/Uint32/Negative", f64(-2), Uint32, u32(math.MaxUint32 - 1), nilErr},
		{"Float64/Uint32/Positive", f64(1881), Uint32, u32(1881), nilErr},
		{"Float64/Uint32/Fractional", f64(18.82), Uint32, u32(18), nilErr},
		{"Float64/Uint32/Overflow", f64(5000000000), Uint32, u32(5000000000 - math.MaxUint32 - 1), nilErr},
		// Float64: Uint64
		{"Float64/Uint64/Negative", f64(-3), Uint64, u64(math.MaxUint64 - 2), nilErr},
		{"Float64/Uint64/Positive", f64(998998), Uint64, u64(998998), nilErr},
		{"Float64/Uint64/Fractional", f64(1234.5678), Uint64, u64(1234), nilErr},
		{"Float64/Uint64/Overflow", f64(18446744073709551615000.5), typeid.Uint64, u64(0x8000000000000000), nilErr},
		// Float64: Int8
		{"Float64/Int8/Negative", f64(-2.0), typeid.Int8, i8(-2), nilErr},
		{"Float64/Int8/Positive", f64(3.0), typeid.Int8, i8(3), nilErr},
		{"Float64/Int8/Fractional", f64(1.23), typeid.Int8, i8(1), nilErr},
		// Float64: Int16
		{"Float64/Int16/Negative", f64(-32), typeid.Int16, i16(-32), nilErr},
		{"Float64/Int16/Positive", f64(23), typeid.Int16, i16(23), nilErr},
		{"Float64/Int16/Fractional", f64(10.01), typeid.Int16, i16(10), nilErr},
		// Float64: Int32
		{"Float64/Int32/Negative", f64(-14), typeid.Int32, i32(-14), nilErr},
		{"Float64/Int32/Positive", f64(41), typeid.Int32, i32(41), nilErr},
		{"Float64/Int32/Fractional", f64(-1.9), typeid.Int32, i32(-1), nilErr},
		// Float64: Int64
		{"Float64/Int64/Negative", f64(-101), typeid.Int64, i64(-101), nilErr},
		{"Float64/Int64/Positive", f64(333), typeid.Int64, i64(333), nilErr},
		{"Float64/Int64/Fractional", f64(18.18), typeid.Int64, i64(18), nilErr},
		// Float64: Float32
		// comparisons here are spotty, so values should be chosen to avoid rounding
		{"Float64/Float32/Negative", f64(-101.0), typeid.Float32, f32(-101.0), nilErr},
		{"Float64/Float32/Positive", f64(32.0), typeid.Float32, f32(32.0), nilErr},
		{"Float64/Float32/Fractional", f64(1.5), typeid.Float32, f32(1.5), nilErr},
		// Float64: Float64 (self)
		{"Float64/Float64", f64(1.1), typeid.Float64, f64(1.1), nilErr},
		// Float64: Other (errors)
		{"Float64/Boolean", f64(0.0), typeid.Boolean, nil, castErr(Float64, typeid.Boolean)},
		{"Float64/String", f64(1.2), typeid.String, nil, castErr(Float64, typeid.String)},
		{"Float64/List", f64(1.2), typeid.List, nil, castErr(Float64, typeid.List)},
		{"Float64/Map", f64(1.2), typeid.Map, nil, castErr(Float64, typeid.Map)},
		{"Float64/Struct", f64(1.2), typeid.Struct, nil, castErr(Float64, typeid.Struct)},
		{"Float64/Class", f64(1.2), typeid.Class, nil, castErr(Float64, typeid.Class)},
		{"Float64/Function", f64(1.2), typeid.Function, nil, castErr(Float64, typeid.Function)},
		{"Float64/Method", f64(1.2), typeid.Method, nil, castErr(Float64, typeid.Method)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual, err := test.value.Downcast(test.to)
			test.err.Require(t, err)

			if actual != nil && test.expected != nil {
				require.Equal(t, test.expected.ID(), actual.ID())
			}

			require.Equal(t, test.expected, actual)
		})
	}
}

// Helper functions
// ================

func f64(v float64) types.Float64 {
	return types.Float64(v)
}

func f32(v float32) types.Float32 {
	return types.Float32(v)
}

func u8(v uint8) types.Uint8 {
	return types.Uint8(v)
}

func u16(v uint16) types.Uint16 {
	return types.Uint16(v)
}

func u32(v uint32) types.Uint32 {
	return types.Uint32(v)
}

func u64(v uint64) types.Uint64 {
	return types.Uint64(v)
}

func i8(v int8) types.Int8 {
	return types.Int8(v)
}

func i16(v int16) types.Int16 {
	return types.Int16(v)
}

func i32(v int32) types.Int32 {
	return types.Int32(v)
}

func i64(v int64) types.Int64 {
	return types.Int64(v)
}
