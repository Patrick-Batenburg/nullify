package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Byte is an NullableImpl byte.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Byte struct {
	NullableImpl[byte]
}

// NewByte creates a new Byte
func NewByte(value byte, valid bool) Byte {
	return Byte{
		NullableImpl: New(value, valid),
	}
}

// ByteFrom creates a new Byte that will always be valid.
func ByteFrom(value byte) Byte {
	return Byte{
		NullableImpl: From(value),
	}
}

// ByteFromPtr creates a new Byte that will be null if the value is nil.
func ByteFromPtr(value *byte) Byte {
	return Byte{
		NullableImpl: FromPtr(value),
	}
}

// MarshalJSON implements json.Marshaler.
func (n Byte) MarshalJSON() ([]byte, error) {
	if !n.IsValid() {
		return NullStringBytes, nil
	}

	if n.IsZero() {
		return []byte(strconv.Quote(ZeroString)), nil
	}

	return []byte(strconv.Quote(string(n.value))), nil
}

// MarshalText implements encoding.TextMarshaler.
func (n Byte) MarshalText() ([]byte, error) {
	if !n.IsValid() {
		return ZeroStringBytes, nil
	}

	if n.IsZero() {
		return []byte(ZeroString), nil
	}

	return []byte{n.value}, nil
}

// Scan implements the sql.Scanner interface.
func (n *Byte) Scan(src any) error {
	if src == nil {
		n.value = ZeroByte
		n.valid = false

		return nil
	}

	switch value := src.(type) {
	case string:
		switch len(value) {
		case 0:
			n.value = ZeroByte
			n.valid = false
		case 1:
			n.value = byte(value[0])
			n.valid = true
		default:
			return NewScannerError(src, n, ErrCannotUnmarshalByte)
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		byteValuer, err := castToUintValuer(value, byte(0))

		if err != nil {
			return NewScannerError(src, n, err)
		}

		n.value = byteValuer.(byte)
		n.valid = true
	default:
		return NewScannerError(value, n)
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Byte) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, NullStringBytes) {
		n.value = 0
		n.valid = false

		return nil
	}

	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return NewUnmarshalError(data, n, err)
	}

	switch len(value) {
	case 0:
		n.value = 0
	case 1:
		n.value = value[0]
	default:
		return NewUnmarshalError(data, n, ErrCannotUnmarshalByte)
	}

	n.valid = true

	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *Byte) UnmarshalText(text []byte) error {
	if text == nil {
		n.value = 0
		n.valid = false

		return nil
	}

	value := string(text)

	switch len(value) {
	case 0:
		n.value = 0
	case 1:
		n.value = value[0]
	default:
		return NewUnmarshalError(text, n, ErrCannotUnmarshalByte)
	}

	n.valid = true

	return nil
}
