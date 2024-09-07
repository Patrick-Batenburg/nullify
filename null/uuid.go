package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

// UUID is a NullableImpl uuid.UUID. It supports SQL and JSON serialization.
// It will marshal to null if null.
type UUID struct {
	NullableImpl[uuid.UUID]
}

// NewUUID creates a new UUID.
// value can be any uuid type with a String method or a string value.
// Will panic if given value cannot be parsed as an uuid.
func NewUUID(value any, valid bool) UUID {
	return UUID{
		NullableImpl: New(newUUID(value), valid),
	}
}

// NewRandomUUID creates a new random UUID
func NewRandomUUID() UUID {
	return UUID{
		NullableImpl: New(uuid.New(), true),
	}
}

// UUIDFrom creates a new UUID that will always be valid.
// value can be any uuid type with a String method.
// Will panic if given value cannot be parsed as an uuid.
func UUIDFrom(value any) UUID {
	return UUID{
		NullableImpl: From(newUUID(value)),
	}
}

// UUIDFromPtr creates a new UUID that will be null if the value is nil.
// value can be any uuid type with a String method.
// Will panic if the given value cannot be parsed as an uuid, unless the value is nil.
func UUIDFromPtr(value any) UUID {
	return UUID{
		NullableImpl: FromPtr(newUUIDFromPtr(value)),
	}
}

// MarshalJSON implements json.Marshaler.
func (n UUID) MarshalJSON() (data []byte, err error) {
	if !n.IsValid() {
		return NullStringBytes, err
	}

	data, err = n.value.MarshalText()

	if err != nil {
		return data, err
	}

	return []byte(strconv.Quote(string(data))), err
}

// MarshalText implements encoding.TextMarshaler.
func (n UUID) MarshalText() ([]byte, error) {
	if !n.IsValid() {
		return EmptyBytes, nil
	}

	return n.value.MarshalText()
}

// Scan implements the sql.Scanner interface.
func (n *UUID) Scan(src any) error {
	switch src.(type) {
	case nil:
		n.value = uuid.Nil
		n.valid = false

		return nil
	}

	err := n.value.Scan(src)

	if err != nil {
		return NewScannerError(src, n, err)
	}

	n.valid = true

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *UUID) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, NullStringBytes) {
		n.value = uuid.Nil
		n.valid = false

		return nil
	}

	var err error
	n.value, err = uuid.ParseBytes(data)

	if err != nil {
		return NewUnmarshalError(data, n, err)
	}

	n.valid = true

	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *UUID) UnmarshalText(text []byte) (err error) {
	if len(text) == 0 {
		n.value = uuid.Nil
		n.valid = false

		return nil
	}

	n.value, err = uuid.ParseBytes(text)

	if err != nil {
		return NewUnmarshalError(text, n, err)
	}

	n.valid = true

	return nil
}

// Value implements the driver.Valuer interface.
func (n UUID) Value() (driver.Value, error) {
	if !n.IsValid() {
		return nil, nil
	}

	return n.value.String(), nil
}

// newUUID creates a new uuid.UUID.
// value can be either a string or any type implementing fmt.Stringer.
// Will panic if the given value cannot be parsed as an UUID.
func newUUID(value any) uuid.UUID {
	var str string

	switch v := value.(type) {
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		panic(ErrCannotNewUUID)
	}

	return uuid.MustParse(str)
}

// newUUIDFromPtr creates a new *uuid.UUID.
// value can be either a string or any type implementing fmt.Stringer.
// Will panic if the given value cannot be parsed as an uuid, unless the value is nil.
func newUUIDFromPtr(value any) (uuidPtr *uuid.UUID) {
	var strPtr *string

	switch v := value.(type) {
	case *string:
		strPtr = v
	case fmt.Stringer:
		strPtr = ptr(v.String())
	case nil:
		return uuidPtr
	default:
		panic(ErrCannotNewUUID)
	}

	if strPtr != nil {
		uuidPtr = ptr(uuid.MustParse(*strPtr))
	}

	return uuidPtr
}
