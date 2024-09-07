package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Uint is an NullableImpl uint.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Uint struct {
	NullableImpl[uint]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewUint(value uint, valid bool, options ...IntegerOption) Uint {
	n := Uint{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func UintFrom(value uint, options ...IntegerOption) Uint {
	n := Uint{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// UintFromPtr creates a new Uint that will be null if the value is nil.
func UintFromPtr(value *uint, options ...IntegerOption) Uint {
	n := Uint{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer interface.
func (n Uint) Value() (driver.Value, error) {
	if !n.IsValid() {
		return nil, nil
	}

	return integerValuerChecker(n.value, n.valuerType)
}

// setValuerType sets valuerType
func (n *Uint) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
