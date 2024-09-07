package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/Patrick-Batenburg/nullify/convert"
)

// GenericNullable is an interface that provides methods for working with
// nullable values. A nullable value is a value that can be either valid
// or invalid (null), allowing for the representation of the absence of a value.
//
// The interface defines methods for:
//
// - Checking if a value is valid or invalid
//
// - Comparing two nullable values
//
// - Retrieving the inner value or its zero value
//
// - Marshaling and unmarshaling to/from JSON and text formats
//
// - Scanning from SQL values
//
// - Setting and retrieving the value as a pointer or directly
//
// - Converting to database driver valuer
//
// Implementations of this interface should provide logic to handle the
// validity of the value, perform marshaling/unmarshaling, work with
// SQL data scanning, and support the driver.Valuer interface for database interactions.
type GenericNullable[T any] interface {
	// Equal returns true if both values are equal and both are valid.
	Equal(other NullableImpl[T]) bool

	// IsValid returns true if the value is valid.
	IsValid() bool

	// IsZero returns true if the value is zero.
	IsZero() bool

	// MarshalJSON implements the json.Marshaler interface.
	MarshalJSON() ([]byte, error)

	// MarshalText implements the encoding.TextMarshaler interface.
	MarshalText() ([]byte, error)

	// MustValue returns the inner value if valid, otherwise panics if accessing an invalid value.
	MustValue() T

	// Ptr returns a pointer to the value, or nil if invalid.
	Ptr() *T

	// Scan implements the sql.Scanner interface.
	Scan(src any) (err error)

	// SetValue sets the value and marks it as valid.
	SetValue(value T)

	// UnmarshalJSON implements the json.Unmarshaler interface.
	UnmarshalJSON(data []byte) (err error)

	// UnmarshalText implements the encoding.TextUnmarshaler interface.
	UnmarshalText(data []byte) (err error)

	// Value implements the driver.Valuer interface.
	Value() (driver.Value, error)

	// ValueOrZero returns the inner value if valid, otherwise the zero value of T.
	ValueOrZero() T
}

// Nullable is an interface that provides methods for working with
// nullable values. A nullable value is a value that can be either valid
// or invalid (null), allowing for the representation of the absence of a value.
//
// The interface defines methods for:
//
// - Checking if a value is valid or invalid
//
// - Marshaling and unmarshaling to/from JSON and text formats
//
// - Scanning from SQL values
//
// - Converting to database driver valuer
//
// Implementations of this interface should provide logic to handle the
// validity of the value, perform marshaling/unmarshaling, work with
// SQL data scanning, and support the driver.Valuer interface for database interactions.
type Nullable interface {
	// IsValid returns true if the value is valid.
	IsValid() bool

	// IsZero returns true if the value is zero.
	IsZero() bool

	// MarshalJSON implements the json.Marshaler interface.
	MarshalJSON() ([]byte, error)

	// MarshalText implements the encoding.TextMarshaler interface.
	MarshalText() ([]byte, error)

	// Scan implements the sql.Scanner interface.
	Scan(src any) (err error)

	// UnmarshalJSON implements the json.Unmarshaler interface.
	UnmarshalJSON(data []byte) (err error)

	// UnmarshalText implements the encoding.TextUnmarshaler interface.
	UnmarshalText(data []byte) (err error)

	// Value implements the driver.Valuer interface.
	Value() (driver.Value, error)
}

// Ensure NullableImpl implements Nullable interface
var (
	_ GenericNullable[bool]      = (*NullableImpl[bool])(nil)
	_ GenericNullable[bool]      = (*Bool)(nil)
	_ GenericNullable[byte]      = (*Byte)(nil)
	_ GenericNullable[[]byte]    = (*Bytes)(nil)
	_ GenericNullable[float32]   = (*Float32)(nil)
	_ GenericNullable[float64]   = (*Float64)(nil)
	_ GenericNullable[int]       = (*Int)(nil)
	_ GenericNullable[int8]      = (*Int8)(nil)
	_ GenericNullable[int16]     = (*Int16)(nil)
	_ GenericNullable[int32]     = (*Int32)(nil)
	_ GenericNullable[int64]     = (*Int64)(nil)
	_ GenericNullable[[]byte]    = (*JSON)(nil)
	_ GenericNullable[string]    = (*String)(nil)
	_ GenericNullable[time.Time] = (*Time)(nil)
	_ GenericNullable[uint]      = (*Uint)(nil)
	_ GenericNullable[uint8]     = (*Uint8)(nil)
	_ GenericNullable[uint16]    = (*Uint16)(nil)
	_ GenericNullable[uint32]    = (*Uint32)(nil)
	_ GenericNullable[uint64]    = (*Uint64)(nil)
	_ GenericNullable[uuid.UUID] = (*UUID)(nil)

	_ Nullable = (*NullableImpl[bool])(nil)
	_ Nullable = (*Bool)(nil)
	_ Nullable = (*Byte)(nil)
	_ Nullable = (*Bytes)(nil)
	_ Nullable = (*Float32)(nil)
	_ Nullable = (*Float64)(nil)
	_ Nullable = (*Int)(nil)
	_ Nullable = (*Int8)(nil)
	_ Nullable = (*Int16)(nil)
	_ Nullable = (*Int32)(nil)
	_ Nullable = (*Int64)(nil)
	_ Nullable = (*JSON)(nil)
	_ Nullable = (*String)(nil)
	_ Nullable = (*Time)(nil)
	_ Nullable = (*Uint)(nil)
	_ Nullable = (*Uint8)(nil)
	_ Nullable = (*Uint16)(nil)
	_ Nullable = (*Uint32)(nil)
	_ Nullable = (*Uint64)(nil)
	_ Nullable = (*UUID)(nil)
)

// NullableImpl represents a NullableImpl value of any type.
type NullableImpl[T any] struct {
	value T
	valid bool
}

// New creates a new NullableImpl with a specified value and validity.
func New[T any](value T, valid bool) NullableImpl[T] {
	return NullableImpl[T]{
		value: value,
		valid: valid,
	}
}

// From creates a new NullableImpl that is always valid.
func From[T any](value T) NullableImpl[T] {
	return New(value, true)
}

// FromPtr creates a new NullableImpl that is null if the pointer is nil.
func FromPtr[T any](ptr *T) NullableImpl[T] {
	if ptr == nil {
		var zero T

		return New(zero, false)
	}

	return From(*ptr)
}

// Equal returns true if both values are equal and both are valid.
func (n NullableImpl[T]) Equal(other NullableImpl[T]) bool {
	return n.IsValid() == other.IsValid() &&
		n.IsValid() &&
		reflect.DeepEqual(n.value, other.value)
}

// IsValid returns true if the value is valid.
func (n NullableImpl[T]) IsValid() bool {
	return n.valid
}

// IsZero returns true if the value is zero.
func (n NullableImpl[T]) IsZero() bool {
	var zero T

	return reflect.DeepEqual(n.value, zero)
}

// MarshalJSON implements the json.Marshaler interface.
func (n NullableImpl[T]) MarshalJSON() ([]byte, error) {
	if !n.IsValid() {
		return NullStringBytes, nil
	}

	data, err := json.Marshal(n.value)

	if err != nil {
		return data, NewMarshalError(n, err)
	}

	return data, nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (n NullableImpl[T]) MarshalText() ([]byte, error) {
	if !n.IsValid() {
		return ZeroStringBytes, nil
	}

	switch v := any(n.value).(type) {
	case []byte:
		return v, nil
	case string:
		return []byte(v), nil
	case int:
		return []byte(strconv.FormatInt(int64(v), 10)), nil
	case int8:
		return []byte(strconv.FormatInt(int64(v), 10)), nil
	case int16:
		return []byte(strconv.FormatInt(int64(v), 10)), nil
	case int32:
		return []byte(strconv.FormatInt(int64(v), 10)), nil
	case int64:
		return []byte(strconv.FormatInt(v, 10)), nil
	case uint:
		return []byte(strconv.FormatUint(uint64(v), 10)), nil
	case uint8:
		return []byte(strconv.FormatUint(uint64(v), 10)), nil
	case uint16:
		return []byte(strconv.FormatUint(uint64(v), 10)), nil
	case uint32:
		return []byte(strconv.FormatUint(uint64(v), 10)), nil
	case uint64:
		return []byte(strconv.FormatUint(v, 10)), nil
	case float32:
		return []byte(strconv.FormatFloat(float64(v), 'f', -1, 32)), nil
	case float64:
		return []byte(strconv.FormatFloat(v, 'f', -1, 64)), nil
	case bool:
		if v {
			return TrueStringBytes, nil
		}

		return FalseStringBytes, nil
	default:
		// Fallback to JSON marshaling for complex types
		data, err := json.Marshal(n.value)

		if err != nil {
			return data, NewMarshalError(n, err)
		}

		return data, nil
	}
}

// MustValue returns the inner value if valid, otherwise panics if accessing an invalid value.
func (n NullableImpl[T]) MustValue() T {
	if !n.IsValid() {
		panic(ErrCannotMustValue)
	}

	return n.value
}

// Ptr returns a pointer to the value, or nil if invalid.
func (n NullableImpl[T]) Ptr() *T {
	if !n.IsValid() {
		return nil
	}

	return &n.value
}

// Scan implements the sql.Scanner interface.
func (n *NullableImpl[T]) Scan(src any) error {
	if src == nil {
		var zero T
		n.value = zero
		n.valid = false

		return nil
	}

	err := convert.ConvertAssign(&n.value, src)

	if err != nil {
		return NewScannerError(src, n, err)
	}

	n.valid = true

	return nil
}

// SetValue sets the value and marks it as valid.
func (n *NullableImpl[T]) SetValue(value T) {
	n.value = value
	n.valid = true
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *NullableImpl[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, NullStringBytes) {
		var zero T
		n.value = zero
		n.valid = false

		return nil
	}

	err := json.Unmarshal(data, &n.value)

	if err != nil {
		return NewUnmarshalError(data, n, err)
	}

	n.valid = true

	return nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (n *NullableImpl[T]) UnmarshalText(text []byte) error {
	if len(text) == 0 || bytes.Equal(text, ZeroStringBytes) {
		var zero T
		n.value = zero
		n.valid = false

		return nil
	}

	stringValue := string(text)

	switch v := any(n.value).(type) {
	case string:
		n.value = any(stringValue).(T)
	case int:
		value, err := strconv.ParseInt(stringValue, 10, 0)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(int(value)).(T)
	case int8:
		value, err := strconv.ParseInt(stringValue, 10, 8)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(int8(value)).(T)
	case int16:
		value, err := strconv.ParseInt(stringValue, 10, 16)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(int16(value)).(T)
	case int32:
		value, err := strconv.ParseInt(stringValue, 10, 32)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(int32(value)).(T)
	case int64:
		value, err := strconv.ParseInt(stringValue, 10, 64)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(int64(value)).(T)
	case uint:
		value, err := strconv.ParseUint(stringValue, 10, 0)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(uint(value)).(T)
	case uint8:
		value, err := strconv.ParseUint(stringValue, 10, 8)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(uint8(value)).(T)
	case uint16:
		value, err := strconv.ParseUint(stringValue, 10, 16)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(uint16(value)).(T)
	case uint32:
		value, err := strconv.ParseUint(stringValue, 10, 32)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(uint32(value)).(T)
	case uint64:
		value, err := strconv.ParseUint(stringValue, 10, 64)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(uint64(value)).(T)
	case float32:
		value, err := strconv.ParseFloat(stringValue, 32)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(float32(value)).(T)
	case float64:
		value, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(float64(value)).(T)
	case bool:
		value, err := strconv.ParseBool(stringValue)

		if err != nil {
			return NewUnmarshalError(v, n, err)
		}

		n.value = any(value).(T)
	default:
		return NewUnmarshalError(v, n)
	}

	n.valid = true

	return nil
}

// Value implements the driver.Valuer interface.
func (n NullableImpl[T]) Value() (driver.Value, error) {
	if !n.IsValid() {
		return nil, nil
	}

	return n.value, nil
}

// ValueOrZero returns the inner value if valid, otherwise the zero value of T.
func (n NullableImpl[T]) ValueOrZero() T {
	if !n.IsValid() {
		var zero T

		return zero
	}

	return n.value
}

// ptr returns the pointer of value.
func ptr[T any](value T) *T {
	return &value
}
