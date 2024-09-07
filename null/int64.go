package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Int64 is an NullableImpl int64.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int64 struct {
	NullableImpl[int64]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewInt64(value int64, valid bool, options ...IntegerOption) Int64 {
	n := Int64{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func Int64From(value int64, options ...IntegerOption) Int64 {
	n := Int64{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// Int64FromPtr creates a new Int64 that will be null if the value is nil.
func Int64FromPtr(value *int64, options ...IntegerOption) Int64 {
	n := Int64{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer interface.
func (n Int64) Value() (driver.Value, error) {
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
func (n *Int64) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
