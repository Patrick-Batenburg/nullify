package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"bytes"
	"encoding/json"
)

// JSON is a NullableImpl []byte that contains JSON.
//
// You might want to use this in the case where you have lets say a NullableImpl
// JSON column in postgres for instance, where there is one layer of null for
// the postgres column, and then you also have the opportunity to have null
// as a value contained in the json. When unmarshalling json however you
// cannot set 'null' as a value.
type JSON struct {
	Bytes
}

// NewJSON creates a new JSON
func NewJSON(value []byte, valid bool) JSON {
	return JSON{
		Bytes: Bytes{
			NullableImpl: New(value, valid),
		},
	}
}

// JSONFrom creates a new JSON that will always be valid.
func JSONFrom(value []byte) JSON {
	return JSON{
		Bytes: Bytes{
			NullableImpl: From(value),
		},
	}
}

// JSONFromPtr creates a new JSON that will be null if the value is nil.
func JSONFromPtr(value *[]byte) JSON {
	return JSON{
		Bytes: Bytes{
			NullableImpl: FromPtr(value),
		},
	}
}

// Marshal will marshal the passed in object,
// and store it in the JSON member on the JSON object.
func (n *JSON) Marshal(data any) error {
	switch v := data.(type) {
	case string:
		if v == ZeroString {
			n.Bytes.value = ZeroBytes
			n.Bytes.valid = false

			return nil
		}

		var rawJSON json.RawMessage
		err := json.Unmarshal([]byte(v), &rawJSON)

		if err != nil {
			return NewUnmarshalError(data, n, err)
		}

		n.Bytes.value = []byte(v)
		n.Bytes.valid = true

		return nil
	case []byte:
		if len(v) == 0 {
			n.Bytes.value = ZeroBytes
			n.Bytes.valid = false

			return nil
		}

		var rawJSON json.RawMessage
		err := json.Unmarshal(v, &rawJSON)

		if err != nil {
			return NewUnmarshalError(data, n, err)
		}

		n.Bytes.value = v
		n.Bytes.valid = true

		return nil
	case nil:
		n.Bytes.value = ZeroBytes
		n.Bytes.valid = false

		return nil
	}

	jsonBytes, err := json.Marshal(data)

	if err != nil {
		return NewUnmarshalError(data, n, err)
	}

	n.Bytes.value = jsonBytes
	n.Bytes.valid = true

	return nil
}

// MarshalJSON implements json.Marshaler.
func (n JSON) MarshalJSON() ([]byte, error) {
	if len(n.value) == 0 {
		return NullStringBytes, nil
	}

	return n.value, nil
}

// MarshalText implements encoding.TextMarshaler.
func (n JSON) MarshalText() ([]byte, error) {
	if !n.IsValid() {
		return ZeroBytes, nil
	}

	return n.value, nil
}

// Unmarshal will unmarshal your JSON stored in
// your JSON object and store the result in the
// value pointed to by dest.
func (n JSON) Unmarshal(dest any) error {
	if dest == nil {
		return NewUnmarshalError(dest, n, ErrDestinationNil)
	}

	// Call our implementation of
	// JSON MarshalJSON through json.Marshal
	// to get the value of the JSON object
	data, err := json.Marshal(n)

	if err != nil {
		return NewUnmarshalError(dest, n, err)
	}

	err = json.Unmarshal(data, dest)

	if err != nil {
		return NewUnmarshalError(dest, n, err)
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Example if you have a struct with a nullify.JSON called v:
//
//	{"v": null} -> calls UnmarshalJSON: !valid
//	{"v": {}}   -> calls UnmarshalJSON: valid (json value is '{}')
//
// If "null" is passed in as json, then the value will be set to nil.
// This way a sql.driver can convert nil to SQL NULL.
func (n *JSON) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullStringBytes) {
		n.value = ZeroBytes
		n.valid = false

		return nil
	}

	n.valid = true
	n.value = make([]byte, len(data))
	copy(n.value, data)

	return nil
}
