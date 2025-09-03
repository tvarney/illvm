package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tvarney/illvm/types"
	"github.com/tvarney/illvm/types/typeid"
)

func TestCastError(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name     string
		from     typeid.ID
		to       typeid.ID
		expected string
	}{
		{"Int8ToInt16", typeid.Int8, typeid.Int16, "unable to cast int8 to int16"},
		{"Int16ToInt8", typeid.Int16, typeid.Int8, "unable to cast int16 to int8"},
		{"Int64ToString", typeid.Int64, typeid.String, "unable to cast int64 to string"},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			e := types.CastError{From: test.from, To: test.to}
			require.Equal(t, test.expected, e.Error())
		})
	}
}
