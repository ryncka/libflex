package flex

// Numeric represents a numeric type, which can be signed, unsigned, or a float.
type Numeric interface {
	Signed | Unsigned | Float
}

// Signed represents a signed integer type.
type Signed interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

// Unsigned represents an unsigned integer type.
type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

// Float represents a floating-point type.
type Float interface {
	~float32 | ~float64
}
