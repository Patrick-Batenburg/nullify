package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Uint32 is an NullableImpl uint32.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Uint32 struct {
	NullableImpl[uint32]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewUint32(value uint32, valid bool, options ...IntegerOption) Uint32 {
	n := Uint32{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func Uint32From(value uint32, options ...IntegerOption) Uint32 {
	n := Uint32{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// Uint32FromPtr creates a new Uint32 that will be null if the value is nil.
func Uint32FromPtr(value *uint32, options ...IntegerOption) Uint32 {
	n := Uint32{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer int32erface.
func (n Uint32) Value() (driver.Value, error) {
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
func (n *Uint32) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
