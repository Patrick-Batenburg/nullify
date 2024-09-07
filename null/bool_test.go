package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBool(t *testing.T) {
	testData := newBoolData()
	nonzero := NewBool(testData.Value, true)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewBool(false, true)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				valid: true,
			},
		},
		zero,
	)

	null := NewBool(testData.Value, false)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestBoolFrom(t *testing.T) {
	testData := newBoolData()
	nonzero := BoolFrom(testData.Value)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := BoolFrom(false)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				valid: true,
			},
		},
		zero,
	)
}

func TestBoolFromPtr(t *testing.T) {
	testData := newBoolData()
	nonzero := BoolFromPtr(testData.Ptr)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := BoolFromPtr(nil)
	assert.Equal(
		t,
		Bool{},
		null,
	)
}

func TestBoolUnmarshalJSON(t *testing.T) {
	testData := newBoolData()
	var nonzero Bool
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Bool
	err = json.Unmarshal(FalseStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				valid: true,
			},
		},
		zero,
	)

	var null Bool
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{},
		null,
	)

	var badType Bool
	err = json.Unmarshal(ZeroIntegerStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Bool{},
		badType,
	)

	var invalid Bool
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
		Bool{},
		invalid,
	)
}

func TestBoolUnmarshalText(t *testing.T) {
	testData := newBoolData()
	var nonzero Bool
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Bool
	err = zero.UnmarshalText(FalseStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				valid: true,
			},
		},
		zero,
	)

	var null Bool
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{},
		null,
	)
}

func TestBoolMarshalJSON(t *testing.T) {
	testData := newBoolData()
	nonzero := BoolFrom(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewBool(false, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		FalseString,
		string(data),
	)

	null := NewBool(true, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestBoolMarshalText(t *testing.T) {
	testData := newBoolData()
	nonzero := BoolFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewBool(false, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		FalseString,
		string(data),
	)

	null := NewBool(false, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestBoolPointer(t *testing.T) {
	testData := newBoolData()
	nonzero := BoolFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewBool(false, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestBoolIsZero(t *testing.T) {
	nonzero := BoolFrom(true)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewBool(false, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewBool(false, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestBoolSetValue(t *testing.T) {
	testData := newBoolData()
	sut := NewBool(false, false)
	assert.Equal(
		t,
		Bool{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestBoolScan(t *testing.T) {
	testData := newBoolData()
	var nonzero Bool
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Bool
	err = zero.Scan(false)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{
			NullableImpl: NullableImpl[bool]{
				valid: true,
			},
		},
		zero,
	)

	var null Bool
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bool{},
		null,
	)
}
