package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInt16(t *testing.T) {
	testData := newInt16Data()
	nonzero := NewInt16(testData.Value, true)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewInt16(ZeroInt16, true)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				valid: true,
			},
		},
		zero,
	)

	null := NewInt16(testData.Value, false)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestInt16From(t *testing.T) {
	testData := newInt16Data()
	nonzero := Int16From(testData.Value)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Int16From(ZeroInt16)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				valid: true,
			},
		},
		zero,
	)
}

func TestInt16FromPtr(t *testing.T) {
	testData := newInt16Data()
	nonzero := Int16FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Int16FromPtr(nil)
	assert.Equal(
		t,
		Int16{},
		null,
	)
}

func TestInt16UnmarshalJSON(t *testing.T) {
	testData := newInt16Data()
	var nonzero Int16
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int16
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				valid: true,
			},
		},
		zero,
	)

	var null Int16
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{},
		null,
	)

	var badType Int16
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Int16{},
		badType,
	)

	var invalid Int16
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
		Int16{},
		invalid,
	)
}

func TestInt16UnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Int16
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Int16{},
		null,
	)
}

func TestInt16UnmarshalText(t *testing.T) {
	testData := newInt16Data()
	var nonzero Int16
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int16
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				valid: true,
			},
		},
		zero,
	)

	var null Int16
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{},
		null,
	)
}

func TestInt16MarshalJSON(t *testing.T) {
	testData := newInt16Data()
	nonzero := Int16From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewInt16(ZeroInt16, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt16(ZeroInt16, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestInt16MarshalText(t *testing.T) {
	testData := newInt16Data()
	nonzero := Int16From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewInt16(ZeroInt16, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt16(ZeroInt16, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestInt16Pointer(t *testing.T) {
	testData := newInt16Data()
	nonzero := Int16From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewInt16(ZeroInt16, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestInt16IsZero(t *testing.T) {
	testData := newInt16Data()
	nonzero := Int16From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewInt16(ZeroInt16, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewInt16(ZeroInt16, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestInt16SetValue(t *testing.T) {
	testData := newInt16Data()
	sut := NewInt16(ZeroInt16, false)
	assert.Equal(
		t,
		Int16{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestInt16Scan(t *testing.T) {
	testData := newInt16Data()
	var nonzero Int16
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int16
	err = zero.Scan(ZeroInt16)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{
			NullableImpl: NullableImpl[int16]{
				valid: true,
			},
		},
		zero,
	)

	var null Int16
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int16{},
		null,
	)
}
