package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"bytes"
	"encoding/json"
)

// Bytes is a NullableImpl []byte.
type Bytes struct {
	NullableImpl[[]byte]
}

// NewBytes creates a new Bytes
func NewBytes(value []byte, valid bool) Bytes {
	return Bytes{
		NullableImpl: New(value, valid),
	}
}

// BytesFrom creates a new Bytes that will always be valid.
func BytesFrom(value []byte) Bytes {
	return Bytes{
		NullableImpl: From(value),
	}
}

// BytesFromPtr creates a new Bytes that will be null if the value is nil.
func BytesFromPtr(value *[]byte) Bytes {
	return Bytes{
		NullableImpl: FromPtr(value),
	}
}

// MarshalJSON implements json.Marshaler.
func (n Bytes) MarshalJSON() ([]byte, error) {
	if len(n.value) == 0 || n.value == nil {
		return NullStringBytes, nil
	}

	return n.value, nil
}

// MarshalText implements encoding.TextMarshaler.
func (n Bytes) MarshalText() ([]byte, error) {
	if !n.IsValid() {
		return nil, nil
	}

	return n.value, nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Bytes) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, NullStringBytes) {
		n.value = ZeroBytes
		n.valid = false

		return nil
	}

	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return NewUnmarshalError(data, n, err)
	}

	n.value = []byte(value)
	n.valid = true

	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *Bytes) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		n.value = ZeroBytes
		n.valid = false
	} else {
		n.value = append(n.value[0:0], text...)
		n.valid = true
	}

	return nil
}
