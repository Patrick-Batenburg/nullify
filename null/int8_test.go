package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInt8(t *testing.T) {
	testData := newInt8Data()
	nonzero := NewInt8(testData.Value, true)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewInt8(ZeroInt8, true)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				valid: true,
			},
		},
		zero,
	)

	null := NewInt8(testData.Value, false)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestInt8From(t *testing.T) {
	testData := newInt8Data()
	nonzero := Int8From(testData.Value)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Int8From(ZeroInt8)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				valid: true,
			},
		},
		zero,
	)
}

func TestInt8FromPtr(t *testing.T) {
	testData := newInt8Data()
	nonzero := Int8FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Int8FromPtr(nil)
	assert.Equal(
		t,
		Int8{},
		null,
	)
}

func TestInt8UnmarshalJSON(t *testing.T) {
	testData := newInt8Data()
	var nonzero Int8
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int8
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				valid: true,
			},
		},
		zero,
	)

	var null Int8
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{},
		null,
	)

	var badType Int8
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Int8{},
		badType,
	)

	var invalid Int8
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
		Int8{},
		invalid,
	)
}

func TestInt8UnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Int8
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Int8{},
		null,
	)
}

func TestInt8UnmarshalText(t *testing.T) {
	testData := newInt8Data()
	var nonzero Int8
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int8
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				valid: true,
			},
		},
		zero,
	)

	var null Int8
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{},
		null,
	)
}

func TestInt8MarshalJSON(t *testing.T) {
	testData := newInt8Data()
	nonzero := Int8From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewInt8(ZeroInt8, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt8(ZeroInt8, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestInt8MarshalText(t *testing.T) {
	testData := newInt8Data()
	nonzero := Int8From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewInt8(ZeroInt8, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt8(ZeroInt8, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestInt8Pointer(t *testing.T) {
	testData := newInt8Data()
	nonzero := Int8From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewInt8(ZeroInt8, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestInt8IsZero(t *testing.T) {
	testData := newInt8Data()
	nonzero := Int8From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewInt8(ZeroInt8, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewInt8(ZeroInt8, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestInt8SetValue(t *testing.T) {
	testData := newInt8Data()
	sut := NewInt8(ZeroInt8, false)
	assert.Equal(
		t,
		Int8{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestInt8Scan(t *testing.T) {
	testData := newInt8Data()
	var nonzero Int8
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int8
	err = zero.Scan(ZeroInt8)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{
			NullableImpl: NullableImpl[int8]{
				valid: true,
			},
		},
		zero,
	)

	var null Int8
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int8{},
		null,
	)
}
