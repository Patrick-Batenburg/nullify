package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Int8 is an NullableImpl int8.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int8 struct {
	NullableImpl[int8]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewInt8(value int8, valid bool, options ...IntegerOption) Int8 {
	n := Int8{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func Int8From(value int8, options ...IntegerOption) Int8 {
	n := Int8{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// Int8FromPtr creates a new Int8 that will be null if the value is nil.
func Int8FromPtr(value *int8, options ...IntegerOption) Int8 {
	n := Int8{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer interface.
func (n Int8) Value() (driver.Value, error) {
	if !n.IsValid() {
		return nil, nil
	}

	value, err := integerValuerChecker(n.value, n.valuerType)

	if err != nil {
		return value, NewValuerError(n, err)
	}

	return value, nil
}

// setValuerType sets valuerType
func (n *Int8) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
