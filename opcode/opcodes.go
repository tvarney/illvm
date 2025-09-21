package opcode

type ID uint8

const (
	// NoOp instructs the VM to do nothing on the current cycle.
	NoOp ID = iota

	Push    // [..]         -> [.., V]
	Dupe    // [.., V]      -> [.., V, V]
	Pop     // [.., V]      -> [..]
	Swap    // [.., A, B]   -> [.., B, A]
	Reverse // [.., s0..sN] -> [.., sN..s0]
	Length  // [.., s0..sN] -> [.., s0..sN, N]

	Add    // [.., A, B] -> [.., B + A]
	Sub    // [.., A, B] -> [.., B - A]
	Mul    // [.., A, B] -> [.., B * A]
	Div    // [.., A, B] -> [.., B / A]
	FDiv   // [.., A, B] -> [.., int(B / A)]
	Mod    // [.., A, B] -> [.., B % A]
	DivMod // [.., A, B] -> [.., int(B / A), B % A]

	And // [.., A, B] -> [.., A & B]
	Or  // [.., A, B] -> [.., A | B]
	Xor // [.., A, B] -> [.., A ^ B]
	Not // [.., V]    -> [.., ~V]
)
