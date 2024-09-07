package null

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewByte(t *testing.T) {
	testData := newByteData()
	nonzero := NewByte(testData.Value, true)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewByte(0, true)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				valid: true,
			},
		},
		zero,
	)

	null := NewByte(testData.Value, false)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestByteFrom(t *testing.T) {
	testData := newByteData()
	nonzero := ByteFrom(testData.Value)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := ByteFrom(0)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				valid: true,
			},
		},
		zero,
	)
}

func TestByteFromPtr(t *testing.T) {
	testData := newByteData()
	nonzero := ByteFromPtr(testData.Ptr)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := ByteFromPtr(nil)
	assert.Equal(
		t,
		Byte{},
		null,
	)
}

func TestByteUnmarshalJSON(t *testing.T) {
	testData := newByteData()
	var nonzero Byte
	err := json.Unmarshal(testData.JSONBytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Byte
	err = json.Unmarshal([]byte(strconv.Quote(ZeroString)), &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				valid: true,
			},
		},
		zero,
	)

	var null Byte
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{},
		null,
	)

	var badType Byte
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Byte{},
		badType,
	)

	var invalid Byte
	err = json.Unmarshal(invalidJSON, &invalid)
	var syntaxErr *json.SyntaxError
	require.ErrorAs(
		t,
		err,
		&syntaxErr,
		"expected error to be of type *json.SyntaxError",
	)
	assert.Equal(
		t,
		Byte{},
		invalid,
	)
}

func TestByteUnmarshalText(t *testing.T) {
	testData := newByteData()
	var nonzero Byte
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Byte
	err = zero.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				valid: true,
			},
		},
		zero,
	)

	var null Byte
	err = null.UnmarshalText(ZeroBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{},
		null,
	)
}

func TestByteMarshalJSON(t *testing.T) {
	testData := newByteData()
	nonzero := ByteFrom(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.JSONString,
		string(data),
	)

	zero := NewByte(0, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		strconv.Quote(ZeroString),
		string(data),
	)

	null := NewByte(0, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestByteMarshalText(t *testing.T) {
	testData := newByteData()
	nonzero := ByteFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewByte(0, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)

	null := NewByte(0, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestBytePointer(t *testing.T) {
	testData := newByteData()
	nonzero := ByteFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewByte(0, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestByteIsZero(t *testing.T) {
	testData := newByteData()
	nonzero := ByteFrom(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewByte(0, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewByte(0, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestByteSetValue(t *testing.T) {
	testData := newByteData()
	sut := NewByte(0, false)
	assert.Equal(
		t,
		Byte{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestByteScan(t *testing.T) {
	testData := newByteData()
	var nonzero Byte
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Byte
	err = zero.Scan(0)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{
			NullableImpl: NullableImpl[byte]{
				valid: true,
			},
		},
		zero,
	)

	var null Byte
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Byte{},
		null,
	)

	var invalid Byte
	err = null.Scan(math.MaxUint8 + 1)
	require.ErrorIs(
		t,
		err,
		ErrValuerCheckerIntegerOverflow,
	)
	assert.Equal(
		t,
		Byte{},
		invalid,
	)
}
