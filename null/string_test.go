package null

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewString(t *testing.T) {
	testData := newStringData()
	nonzero := NewString(testData.Value, true)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewString(ZeroString, true)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				valid: true,
			},
		},
		zero,
	)

	null := NewString(testData.Value, false)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestStringFrom(t *testing.T) {
	testData := newStringData()
	nonzero := StringFrom(testData.Value)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := StringFrom(ZeroString)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				valid: true,
			},
		},
		zero,
	)
}

func TestStringFromPtr(t *testing.T) {
	testData := newStringData()
	nonzero := StringFromPtr(testData.Ptr)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := StringFromPtr(nil)
	assert.Equal(
		t,
		String{},
		null,
	)
}

func TestStringUnmarshalJSON(t *testing.T) {
	testData := newStringData()
	var nonzero String
	err := json.Unmarshal(testData.JSONBytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero String
	err = json.Unmarshal([]byte(strconv.Quote(ZeroString)), &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				valid: true,
			},
		},
		zero,
	)

	var null String
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		String{},
		null,
	)

	var badType String
	err = json.Unmarshal(ZeroIntegerStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		String{},
		badType,
	)

	var invalid String
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
		String{},
		invalid,
	)
}

func TestStringUnmarshalText(t *testing.T) {
	testData := newStringData()
	var nonzero String
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var null String
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		String{},
		null,
	)
}

func TestStringMarshalJSON(t *testing.T) {
	testData := newStringData()
	nonzero := StringFrom(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.JSONString,
		string(data),
	)

	zero := NewString(ZeroString, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		strconv.Quote(ZeroString),
		string(data),
	)

	null := NewString(gofakeit.LetterN(math.MaxInt8), false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestStringMarshalText(t *testing.T) {
	testData := newStringData()
	nonzero := StringFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.Value,
		string(data),
		nonzero,
	)

	zero := NewString(ZeroString, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)

	null := NewString(gofakeit.LetterN(math.MaxInt8), false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestStringPointer(t *testing.T) {
	testData := newStringData()
	nonzero := StringFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewString(ZeroString, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestStringIsZero(t *testing.T) {
	nonzero := StringFrom(gofakeit.LetterN(math.MaxInt8))
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewString(ZeroString, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewString(ZeroString, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestStringSetValue(t *testing.T) {
	testData := newStringData()
	sut := NewString(ZeroString, false)
	assert.Equal(
		t,
		String{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestStringScan(t *testing.T) {
	testData := newStringData()
	var nonzero String
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero String
	err = zero.Scan(ZeroString)
	require.NoError(t, err)
	assert.Equal(
		t,
		String{
			NullableImpl: NullableImpl[string]{
				valid: true,
			},
		},
		zero,
	)

	var null String
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		String{},
		null,
	)
}
