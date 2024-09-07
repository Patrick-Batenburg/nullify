package null // import "github.com/Patrick-Batenburg/nullify/null"

// Float64 is a NullableImpl float64.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Float64 struct {
	NullableImpl[float64]
}

// NewFloat64 creates a new Float64
func NewFloat64(value float64, valid bool) Float64 {
	return Float64{
		NullableImpl: New(value, valid),
	}
}

// Float64From creates a new Float64 that will always be valid.
func Float64From(value float64) Float64 {
	return Float64{
		NullableImpl: From(value),
	}
}

// Float64FromPtr creates a new Float64 that will be null if the value is nil.
func Float64FromPtr(value *float64) Float64 {
	return Float64{
		NullableImpl: FromPtr(value),
	}
}
