package vm

import (
	"github.com/tvarney/consterr"
	"github.com/tvarney/illvm/opcode"
	"github.com/tvarney/illvm/types"
)

const (
	// ErrBytecodeOverflow indicates that the virtual machine attempted to read
	// past the end of the bytecode.
	ErrBytecodeOverflow consterr.Error = "no more opcodes"

	// ErrOperationUndefined indicates that the virtual machine encountered an
	// opcode it couldn't execute.
	ErrOperationUndefined consterr.Error = "operation undefined"
)

// Thread is a single execution context of a illvm virtual machine.
type Thread struct {
	Machine *Machine
	Stack   []types.Value
	Data    []uint8

	PC int
}

// Run runs the opcodes in the given bytecode data until an error occurs.
func (t *Thread) Run() error {
	err := t.Step()
	for err == nil {
		err = t.Step()
	}

	return err
}

// RunFor runs the next `step` opcodes.
func (t *Thread) RunFor(steps int) error {
	for range steps {
		if err := t.Step(); err != nil {
			return err
		}
	}

	return nil
}

// Step fetches the next opcode and any associated data and executes it.
func (t *Thread) Step() error {
	if t.PC < 0 || t.PC >= len(t.Data) {
		return ErrBytecodeOverflow
	}

	op := t.Data[t.PC]
	t.PC++

	switch opcode.ID(op) {
	case opcode.NoOp:
		return nil
	default:
		return ErrOperationUndefined
	}
}
