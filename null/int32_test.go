package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInt32(t *testing.T) {
	testData := newInt32Data()
	nonzero := NewInt32(testData.Value, true)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewInt32(ZeroInt32, true)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				valid: true,
			},
		},
		zero,
	)

	null := NewInt32(testData.Value, false)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestInt32From(t *testing.T) {
	testData := newInt32Data()
	nonzero := Int32From(testData.Value)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Int32From(ZeroInt32)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				valid: true,
			},
		},
		zero,
	)
}

func TestInt32FromPtr(t *testing.T) {
	testData := newInt32Data()
	nonzero := Int32FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Int32FromPtr(nil)
	assert.Equal(
		t,
		Int32{},
		null,
	)
}

func TestInt32UnmarshalJSON(t *testing.T) {
	testData := newInt32Data()
	var nonzero Int32
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int32
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				valid: true,
			},
		},
		zero,
	)

	var null Int32
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{},
		null,
	)

	var badType Int32
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Int32{},
		badType,
	)

	var invalid Int32
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
		Int32{},
		invalid,
	)
}

func TestInt32UnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Int32
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Int32{},
		null,
	)
}

func TestInt32UnmarshalText(t *testing.T) {
	testData := newInt32Data()
	var nonzero Int32
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int32
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				valid: true,
			},
		},
		zero,
	)

	var null Int32
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{},
		null,
	)
}

func TestInt32MarshalJSON(t *testing.T) {
	testData := newInt32Data()
	nonzero := Int32From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewInt32(ZeroInt32, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt32(ZeroInt32, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestInt32MarshalText(t *testing.T) {
	testData := newInt32Data()
	nonzero := Int32From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewInt32(ZeroInt32, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt32(ZeroInt32, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestInt32Pointer(t *testing.T) {
	testData := newInt32Data()
	nonzero := Int32From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewInt32(ZeroInt32, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestInt32IsZero(t *testing.T) {
	testData := newInt32Data()
	nonzero := Int32From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewInt32(ZeroInt32, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewInt32(ZeroInt32, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestInt32SetValue(t *testing.T) {
	testData := newInt32Data()
	sut := NewInt32(ZeroInt32, false)
	assert.Equal(
		t,
		Int32{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestInt32Scan(t *testing.T) {
	testData := newInt32Data()
	var nonzero Int32
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int32
	err = zero.Scan(ZeroInt32)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{
			NullableImpl: NullableImpl[int32]{
				valid: true,
			},
		},
		zero,
	)

	var null Int32
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int32{},
		null,
	)
}
