package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Int is an NullableImpl int.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int struct {
	NullableImpl[int]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewInt(value int, valid bool, options ...IntegerOption) Int {
	n := Int{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func IntFrom(value int, options ...IntegerOption) Int {
	n := Int{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// IntFromPtr creates a new Int that will be null if the value is nil.
func IntFromPtr(value *int, options ...IntegerOption) Int {
	n := Int{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer interface.
func (n Int) Value() (driver.Value, error) {
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
func (n *Int) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
