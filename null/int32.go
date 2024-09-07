package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Int32 is an NullableImpl int32.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int32 struct {
	NullableImpl[int32]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewInt32(value int32, valid bool, options ...IntegerOption) Int32 {
	n := Int32{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func Int32From(value int32, options ...IntegerOption) Int32 {
	n := Int32{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// Int32FromPtr creates a new Int32 that will be null if the value is nil.
func Int32FromPtr(value *int32, options ...IntegerOption) Int32 {
	n := Int32{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer interface.
func (n Int32) Value() (driver.Value, error) {
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
func (n *Int32) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
