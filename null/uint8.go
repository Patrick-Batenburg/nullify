package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Uint8 is an NullableImpl uint8.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Uint8 struct {
	NullableImpl[uint8]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewUint8(value uint8, valid bool, options ...IntegerOption) Uint8 {
	n := Uint8{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func Uint8From(value uint8, options ...IntegerOption) Uint8 {
	n := Uint8{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// Uint8FromPtr creates a new Uint8 that will be null if the value is nil.
func Uint8FromPtr(value *uint8, options ...IntegerOption) Uint8 {
	n := Uint8{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer int8erface.
func (n Uint8) Value() (driver.Value, error) {
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
func (n *Uint8) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
