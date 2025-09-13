# Overview

The bytecode of illvm is stored as a `[]byte`, and each opcode within that slice
may take an additional `control` byte to change the specifics of its behavior.
E.g. an opcode may take a control byte which indicates that the argument is
stored as a 4-bit value directly in the control, that it is stored as a
(possibly) multi-byte immediate within the bytecode, or that it should be taken
from the stack.

# Control Code Schemes

There are a few shared control code schemes that opcodes use.

# OpCodes

## Notation

Within the table there are multiple columns which use a notation to denote
variables and effects.

If the row does not have values for ID, Category, and Code then it is a
continuation row; that is, it 'inherits' the values for those fields from the
row above which does have values for it.

If a row is missing the ID and Category but has a Code value, the code is an
alias of the previous rows code.

### Control bits.

The control field of an opcode has a few special considerations.

* If the column is blank then that opcode does not consume a control byte.
* If any row for an opcode is non-blank, _all_ rows must be non-blank.
* When non-blank, the field must consist of 8 bit entries:
  * If a bit is significant it will be noted as `0`, `1`, or a variable letter;
    e.g. `N`.
  * If a bit is insignificant (ignored) it will be noted as `-`. Best practice
    is to set these to 0 in bytecode, but they are discarded no matter their
    value.
  * To fully show that it is a bitfield notation, all non-blank entries must
    start with `0b`

#### Examples

* `0b10101010`: All bits are significant and exact values; the row is only
  applicable when the control field is _exactly_ this value.
* `0b1-------`: The first bit is set to 1, all other bits are ignored.
* `0b0VVVVVVV`: The first bit is unset, all other bits contribute to variable
  `V`.
* `0bTTNNVVVV`: The first two bits are `T`, the next two bits `N`, and the last
  4 are `V`.

### Immediates

Immediates may be denoted in any other column using `iN` as the value, where `N`
is the index of the immediate, starting from 0. E.g. if an opcode has an
immediates column of "u8,u8" then the immediates would effectively be "i0,i1"
where i0 corresponds to the first u8 and i1 corresponds to the second u8.

### References

If a variable is being used as a reference, it must be of the form `TYPE($VAR)`,
where the value of `TYPE` denotes the type of the reference and `VAR` is the
value being read. E.g. `lstr($i0)` indicates a local string with ID in the first
immediate. `str($i0,$i1)` indicates a module string with module ID in the first
immediate and string ID in the second.

### Special Variables

There are a few special variables that may be used in the table.

 * `%pc` - the current 'program counter' within the current context. This is
   evaluated _after_ the opcode is run. That is, it does not point to the
   current opcode, but the start of the next opcode.

### Stack Changes

Stack changes are denoted by the notation `[] -> []` where the first `[]` is the
state of the stack prior to the opcode and the second `[]` is the state of the
stack after the opcode. `[]` always denotes a completely empty stack. Where the
contents of the stack are unspecified `..` is used. That is, `[..]` is a stack
with unknown contents, where the contents are unspecified. `[..]` may represent
an empty stack as well. That is, `[..]` may contain any number of items
including 0, while `[]` is always the zero length stack. E.g. a "Clear" opcode
would represent the stack as `[..] -> []`, that is it takes an arbitrary stack
and the resulting stack is empty.

As the stack in illvm is segmented into frames, the bounds of the current frame
may be denoted in the stack notation with a `|` field. E.g. a stack of
`[.., a | b ]` has two frames referenced, with the variable `a` in the prior
frame and variable `b` in the current frame.

This allows denoting when a new frame is created.

In general, all stack changes are bound to the current frame. With a few
exceptions, any operation that would attempt to operate outside the current
frame result in a VM fault. E.g. if the stack is `[.. | a, b, c]` and `Pop 5` is
the opcode to run, the VM will fault as the current frame has fewer than 5
elements.

For a variable number of elements in a stack, the values will be denoted as a
variable itself with the special notation `s0..sN`. This may also be used for
large runs of variables. `s0..s10` denotes 11 variables, `s0`, `s1`, `s2`, etc.
all the way up to `s10`. E.g. a varargs function call would look something like
`[.., C, N, s0..sN] -> [.., %pc | N, s0..sN]`.

### Variable Types

Where applicable, the following abbreviations are used for types.

| Name | Bytes | Type                    | Notes
| -----|-------|-------------------------|------
| u8   |     1 | unsigned 8-bit integer  |
| u16  |     2 | unsigned 16-bit integer |
| u24  |     3 | unsigned 24-bit integer | Only valid for immediates
| u32  |     4 | unsigned 32-bit integer |
| u40  |     5 | unsigned 40-bit integer | Only valid for immediates
| u48  |     6 | unsigned 48-bit integer | Only valid for immediates
| u56  |     7 | unsigned 56-bit integer | Only valid for immediates
| u64  |     8 | unsigned 64-bit integer |
| uN   |     N | unsigned N-byte integer | N must be defined in another column. The variable may be any single-letter; e.g. uM, uA, etc. The value may only be 1-8 inclusive.
| i8   |     1 | signed 8-bit integer    |
| i16  |     2 | signed 16-bit integer   |
| i24  |     3 | signed 24-bit integer   | Only valid for immediates
| i32  |     4 | signed 32-bit integer   |
| i40  |     3 | signed 40-bit integer   | Only valid for immediates
| i48  |     3 | signed 48-bit integer   | Only valid for immediates
| i46  |     3 | signed 56-bit integer   | Only valid for immediates
| i64  |     4 | signed 64-bit integer   |
| iN   |     N | signed N-byte integer   | N must be defined in another column. The variable may be any single-letter; e.g. iM, iA, etc. The value may only be 1-8 inclusive.
| f32  |     4 | 32-bit floating point   |
| f64  |     8 | 64-bit floating point   | 
| fN   |     N | N-byte floating point   | N must be defined in another column. The variable may be any single-letter; e.g. fM, fA, etc. The value must be 4 or 8.
| lstr |     4 | A 'local' string ID     | An optimized form of `str` where the module ID is 0 (self)
| str  |     6 | A general string ID     | The value consists of a 2-byte value for the module and a 4-byte index into a constant string table for said module.


## Table

| ID     | Category | Code  |   Control    | Immediates | Stack                                | Side Effect | Notes
|--------|----------|-------|--------------|------------|--------------------------------------|-------------|-------------
| `0x00` | Misc     | NoOp  |              |            |                                      |             |
| `0x01` | Stack    | Push  | `0b1VVVVVVV` |            | `[..]->[..,V]`                       |             |
|        |          |       | `0b0000NNNN` | `uN`       | `[..]->[..,i0]`                      |             | `N` must be 1-8.
|        |          |       | `0b0001NNNN` | `iN`       | `[..]->[..,i0]`                      |             | `N` must be 1-8.
|        |          |       | `0b0010NNNN` | `fN`       | `[..]->[..,i0]`                      |             | `N` must be 4 or 8.
| `0x02` | Stack    | Pop   | `0b1VVVVVVV` |            | `[..,s0..sV]->[..]`                  |             |
|        |          |       | `0b00------` |            | `[..,s0..sN,N]->[..]`                |             | `N` must be a signed or unsigned integer.
|        |          | Clear | `0b01------` |            | `[..]->[]`                           |             |
| `0x03` | Stack    | Swap  | `0b00AAABBB` |            | `[..,$A,..,$B,..]->[..,$B,..,$A,..]` |             | A and B are indicies from the last element of the stack. $A and $B are unordered.
|        |          |       | `0b01AAABBB` |            | `[..,$A,..,$B,..]->[..,$B,..,$A,..]` |             | A and B are indicies from the first element of the stack. $A and $B are unordered.
|        |          |       | `0b100-NNNN` | `uN`       | `[..,$i0,..,V]->[..,V,..,$i0]`       |             | N must be 1-8. i0 is an index from the last element of the stack.
|        |          |       | `0b110-NNNN` | `uN`       | `[..,$i0,..,V]->[..,V,..,$i0]`       |             | N must be 1-8. i0 is an index from the first element of the stack.
|        |          |       | `0b101-NNNN` | `uN,uN`    | `[..,$i0,..,$i1]->[..,$i1,..,$i0]`   |             | N must be 1-8. i0 and i1 are indices from the last element of the stack.
|        |          |       | `0b111-NNNN` | `uN,uN`    | `[..,$i0,..,$i1]->[..,$i1,..,$i0]`   |             | N must be 1-8. i0 and i1 are indices from the first element of the stack.

# Details
## Misc OpCodes
### NoOp

| Name    | Value
|---------|------
| ID      | `0x00`
| Control | No
| Aliases |

`NoOp` has no effect on the vm. The stack is not changed, no extra bytes are
read, and no side effects happen.

## Stack OpCodes
### Push

| Name    | Value
|---------|------
| ID      | `0x01`
| Control | Yes
| Aliases |

`Push` pushes a new value onto the stack. This value must be an immediate of some
sort, and only basic numeric types are valid.

The control byte can be segmented as `0bITTTNNNN`, where `I` is if an immediate
value is stored in the opcode. If `I` is `1` then V is constructed as
`0bTTTTNNNN`, that is a 7-bit unsigned integer. Otherwise, `0bTTTT` denotes the
type of the immediate to push and `0bNNNN` is the size of the immediate.

Valid values of `T` are 0, 1, and 2. If `T` is `0` (`0b0000`) the immediate is
an unsigned integer of byte-size `N`, with valid values of `N` being `[1,8]`.

If `T` is `1` (`0b0001`) the immediate is a signed integer of byte-size `N`,
with valid values of `N` being `[1,8]`.

If `T` is `2` (`0b0010`) the immediate is a floating point number of byte-size
`N`, with valid values of `N` being `4` or `8`.

Any other value of `T` results in a VM fault. If the size `N` is not a valid
value for the type selection of `T` it also results in a VM fault.

### Pop

| Name    | Value
|---------|------
| ID      | `0x02`
| Control | Yes
| Aliases | `Clear`

`Pop` removes a number of items from the stack.

### Swap

| Name    | Value
|---------|------
| ID      | `0x03`
| Control | Yes
| Aliases |

`Swap` swaps two items within the current stack frame.