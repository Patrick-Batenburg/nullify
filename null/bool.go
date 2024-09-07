package null // import "github.com/Patrick-Batenburg/nullify/null"

// Bool is a NullableImpl bool.
// It does not consider false values to be null.
// It will decode to null, not false, if null.
type Bool struct {
	NullableImpl[bool]
}

// NewBool creates a new Bool
func NewBool(value bool, valid bool) Bool {
	return Bool{
		NullableImpl: New(value, valid),
	}
}

// BoolFrom creates a new Bool that will always be valid.
func BoolFrom(value bool) Bool {
	return Bool{
		NullableImpl: From(value),
	}
}

// BoolFromPtr creates a new Bool that will be null if the value is nil.
func BoolFromPtr(value *bool) Bool {
	return Bool{
		NullableImpl: FromPtr(value),
	}
}
