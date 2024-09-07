package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Uint16 is an NullableImpl uint16.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Uint16 struct {
	NullableImpl[uint16]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewUint16(value uint16, valid bool, options ...IntegerOption) Uint16 {
	n := Uint16{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func Uint16From(value uint16, options ...IntegerOption) Uint16 {
	n := Uint16{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// Uint16FromPtr creates a new Uint16 that will be null if the value is nil.
func Uint16FromPtr(value *uint16, options ...IntegerOption) Uint16 {
	n := Uint16{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer int16erface.
func (n Uint16) Value() (driver.Value, error) {
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
func (n *Uint16) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
