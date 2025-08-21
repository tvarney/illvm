package typeid_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/types/typeid"
)

func TestID(t *testing.T) {
	t.Parallel()

	t.Run("String", func(t *testing.T) {
		t.Parallel()

		for _, test := range []struct {
			name     string
			id       typeid.ID
			expected string
		}{
			{"Void", typeid.Void, "void"},
			{"Uint8", typeid.Uint8, "uint8"},
			{"Uint16", typeid.Uint16, "uint16"},
			{"Uint32", typeid.Uint32, "uint32"},
			{"Uint64", typeid.Uint64, "uint64"},
			{"Int8", typeid.Int8, "int8"},
			{"Int16", typeid.Int16, "int16"},
			{"Int32", typeid.Int32, "int32"},
			{"Int64", typeid.Int64, "int64"},
			{"Float32", typeid.Float32, "float32"},
			{"Float64", typeid.Float64, "float64"},
			{"Boolean", typeid.Boolean, "boolean"},
			{"String", typeid.String, "string"},
			{"List", typeid.List, "list"},
			{"Map", typeid.Map, "map"},
			{"Struct", typeid.Struct, "struct"},
			{"Class", typeid.Class, "class"},
			{"Function", typeid.Function, "function"},
			{"Method", typeid.Method, "method"},
			{"Unknown", typeid.ID(255), "unknown"},
		} {
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()
				require.Equal(t, test.expected, test.id.String())
			})
		}
	})
}
