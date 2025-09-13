# Overview

The VM of illvm is strongly typed and mostly statically typed. The main
exception to the static typing is the conversion between some primitive types
that may happen automatically, e.g. conversion between a u32 and an i32 may be
done for example when adding a u32 and an i32, resulting in an i32.

# Types

| Name | Stack | Stored |Immediate | Description
|------|-------|--------|----------|------------
| u8   | ❌    | ✅    | ✅       | An 8-bit unsigned value, `[0,255]`
| u16  | ❌    | ✅    | ✅       | A 16-bit unsigned value, `[0,65535]`
| u24  | ❌    | ❌    | ✅       | A 24-bit unsigned value, `[0,16777215]`, immediates only
| u32  | ❌    | ✅    | ✅       | A 32-bit unsigned value, `[0,4294967295] `
| u40  | ❌    | ❌    | ✅       | A 40-bit unsigned value, `[0,1099511627775]`, immediates only
| u48  | ❌    | ❌    | ✅       | A 48-bit unsigned value, `[0,281474976710655]`, immediates only
| u56  | ❌    | ❌    | ✅       | A 56-bit unsigned value, `[0,72057594037927935]`, immediates only
| u64  | ✅    | ✅    | ✅       | A 64-bit unsigned value, `[0,18446744073709551615]`
| i8   | ❌    | ✅    | ✅       | An 8-bit signed value, `[-128,127]`
| i16  | ❌    | ✅    | ✅       | A 16-bit signed value, `[-32768,32767]`
| i24  | ❌    | ❌    | ✅       | A 24-bit signed value, `[-8388608,8388607]`, immediates only
| i32  | ❌    | ✅    | ✅       | A 32-bit signed value, `[-2147483648,2147483647]`
| i40  | ❌    | ❌    | ✅       | A 40-bit signed value, `[-549755813888,549755813887]`, immediates only
| i48  | ❌    | ❌    | ✅       | A 48-bit signed value, `[-140737488355328,140737488355327]`, immediates only
| i56  | ❌    | ❌    | ✅       | A 56-bit signed value, `[-36028797018963968,36028797018963967]`, immediates only
| i64  | ✅    | ✅    | ✅       | A 64-bit signed value, `[-9223372036854775808,9223372036854775807]`
| f32  | ✅    | ✅    | ✅       | A 32-bit floating point value
| f64  | ✅    | ✅    | ✅       | A 64-bit floating point value
| bool | ❌    | ✅    | ❌       | A true/false value, stored as a u8 with `0` for `false`, `1` for `true`
| str  | ✅    | ✅    | ❌       | A string
| list | ✅    | ✅    | ❌       | An array of values of a given type.
| map  | ✅    | ✅    | ❌       | A mapping between two types (go `map`, python `dict`)
| type | ✅    | ✅    | ❌       | A type definition stored in the VM.
| obj  | ✅    | ✅    | ❌       | An instance of a type (`class`, `struct`)

For integer ranges, the values are inclusive.