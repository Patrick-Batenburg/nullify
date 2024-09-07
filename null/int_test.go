package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInt(t *testing.T) {
	testData := newIntData()
	nonzero := NewInt(testData.Value, true)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewInt(ZeroInt, true)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				valid: true,
			},
		},
		zero,
	)

	null := NewInt(testData.Value, false)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestIntFrom(t *testing.T) {
	testData := newIntData()
	nonzero := IntFrom(testData.Value)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := IntFrom(ZeroInt)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				valid: true,
			},
		},
		zero,
	)
}

func TestIntFromPtr(t *testing.T) {
	testData := newIntData()
	nonzero := IntFromPtr(testData.Ptr)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := IntFromPtr(nil)
	assert.Equal(
		t,
		Int{},
		null,
	)
}

func TestIntUnmarshalJSON(t *testing.T) {
	testData := newIntData()
	var nonzero Int
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				valid: true,
			},
		},
		zero,
	)

	var null Int
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{},
		null,
	)

	var badType Int
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Int{},
		badType,
	)

	var invalid Int
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
		Int{},
		invalid,
	)
}

func TestIntUnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Int
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Int{},
		null,
	)
}

func TestIntUnmarshalText(t *testing.T) {
	testData := newIntData()
	var nonzero Int
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				valid: true,
			},
		},
		zero,
	)

	var null Int
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{},
		null,
	)
}

func TestIntMarshalJSON(t *testing.T) {
	testData := newIntData()
	nonzero := IntFrom(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewInt(ZeroInt, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt(ZeroInt, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestIntMarshalText(t *testing.T) {
	testData := newIntData()
	nonzero := IntFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewInt(ZeroInt, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt(ZeroInt, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestIntPointer(t *testing.T) {
	testData := newIntData()
	nonzero := IntFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewInt(ZeroInt, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestIntIsZero(t *testing.T) {
	testData := newIntData()
	nonzero := IntFrom(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewInt(ZeroInt, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewInt(ZeroInt, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestIntSetValue(t *testing.T) {
	testData := newIntData()
	sut := NewInt(ZeroInt, false)
	assert.Equal(
		t,
		Int{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestIntScan(t *testing.T) {
	testData := newIntData()
	var nonzero Int
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int
	err = zero.Scan(ZeroInt)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{
			NullableImpl: NullableImpl[int]{
				valid: true,
			},
		},
		zero,
	)

	var null Int
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int{},
		null,
	)
}
