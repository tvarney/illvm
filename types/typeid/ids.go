package typeid

type Id uint8

const (
	Void Id = iota
	Uint8
	Uint16
	Uint32
	Uint64
	Int8
	Int16
	Int32
	Int64
	Float32
	Float64
	Boolean
	String
	List
	Map
	Struct
	Class
	Function
	Method
	// Closure?
)

func (i Id) String() string {
	switch i {
	case Void:
		return "void"
	case Uint8:
		return "uint8"
	case Uint16:
		return "uint16"
	case Uint32:
		return "uint32"
	case Uint64:
		return "uint64"
	case Int8:
		return "int8"
	case Int16:
		return "int16"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	case Boolean:
		return "boolean"
	case String:
		return "string"
	case List:
		return "list"
	case Map:
		return "map"
	case Struct:
		return "struct"
	case Class:
		return "class"
	case Function:
		return "function"
	case Method:
		return "method"
	}

	return "unknown"
}
