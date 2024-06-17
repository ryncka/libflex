package flex

type Numeric interface {
	Signed | Unsigned | Float
}

type Signed interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type Unsigned interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint
}

type Float interface {
	~float32 | ~float64
}
