package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
)

// Uint64 is an NullableImpl uint64.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Uint64 struct {
	NullableImpl[uint64]

	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// NewInt creates a new Int
func NewUint64(value uint64, valid bool, options ...IntegerOption) Uint64 {
	n := Uint64{
		NullableImpl: New(value, valid),
	}

	n.setValuerType(options...)

	return n
}

// IntFrom creates a new Int that will always be valid.
func Uint64From(value uint64, options ...IntegerOption) Uint64 {
	n := Uint64{
		NullableImpl: From(value),
	}

	n.setValuerType(options...)

	return n
}

// Uint64FromPtr creates a new Uint64 that will be null if the value is nil.
func Uint64FromPtr(value *uint64, options ...IntegerOption) Uint64 {
	n := Uint64{
		NullableImpl: FromPtr(value),
	}

	n.setValuerType(options...)

	return n
}

// Value implements the driver.Valuer int64erface.
func (n Uint64) Value() (driver.Value, error) {
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
func (n *Uint64) setValuerType(options ...IntegerOption) {
	if len(options) > 0 {
		integerOption := new(integerOption)
		options[0](integerOption)
		n.valuerType = integerOption.valuerType
	}
}
