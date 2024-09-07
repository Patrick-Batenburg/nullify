package null // import "github.com/Patrick-Batenburg/nullify/null"

// String is a NullableImpl string. It supports SQL and JSON serialization.
// It will marshal to null if null. Blank string input will be considered null.
type String struct {
	NullableImpl[string]
}

// NewString creates a new String
func NewString(value string, valid bool) String {
	return String{
		NullableImpl: New(value, valid),
	}
}

// StringFrom creates a new String that will always be valid.
func StringFrom(value string) String {
	return String{
		NullableImpl: From(value),
	}
}

// StringFromPtr creates a new String that will be null if the value is nil.
func StringFromPtr(value *string) String {
	return String{
		NullableImpl: FromPtr(value),
	}
}
