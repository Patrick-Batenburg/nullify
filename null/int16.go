package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Int16 is an NullableImpl int16.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int16 struct {
	NullableImpl[int16]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewInt16(value int16, valid bool, options ...IntegerOption) Int16 {
	n := Int16{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func Int16From(value int16, options ...IntegerOption) Int16 {
	n := Int16{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// Int16FromPtr creates a new Int16 that will be null if the value is nil.
func Int16FromPtr(value *int16, options ...IntegerOption) Int16 {
	n := Int16{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer interface.
func (n Int16) Value() (driver.Value, error) {
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
func (n *Int16) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
