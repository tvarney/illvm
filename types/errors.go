package types

import (
	"fmt"

	"github.com/tvarney/illvm/types/typeid"
)

// CastError is an error type which indicates that casting from one value type
// to another was not possible.
type CastError struct {
	From typeid.Id
	To   typeid.Id
}

func (e CastError) Error() string {
	return fmt.Sprintf("unable to cast %s to %s", e.From, e.To)
}
