package null // import "github.com/Patrick-Batenburg/nullify/null"

// Float32 is a NullableImpl float32.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Float32 struct {
	NullableImpl[float32]
}

// NewFloat32 creates a new Float32
func NewFloat32(value float32, valid bool) Float32 {
	return Float32{
		NullableImpl: New(value, valid),
	}
}

// Float32From creates a new Float32 that will always be valid.
func Float32From(value float32) Float32 {
	return Float32{
		NullableImpl: From(value),
	}
}

// Float32FromPtr creates a new Float32 that will be null if the value is nil.
func Float32FromPtr(value *float32) Float32 {
	return Float32{
		NullableImpl: FromPtr(value),
	}
}
