package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInt64(t *testing.T) {
	testData := newInt64Data()
	nonzero := NewInt64(testData.Value, true)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewInt64(ZeroInt64, true)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				valid: true,
			},
		},
		zero,
	)

	null := NewInt64(testData.Value, false)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestInt64From(t *testing.T) {
	testData := newInt64Data()
	nonzero := Int64From(testData.Value)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Int64From(ZeroInt64)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				valid: true,
			},
		},
		zero,
	)
}

func TestInt64FromPtr(t *testing.T) {
	testData := newInt64Data()
	nonzero := Int64FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Int64FromPtr(nil)
	assert.Equal(
		t,
		Int64{},
		null,
	)
}

func TestInt64UnmarshalJSON(t *testing.T) {
	testData := newInt64Data()
	var nonzero Int64
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int64
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				valid: true,
			},
		},
		zero,
	)

	var null Int64
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{},
		null,
	)

	var badType Int64
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Int64{},
		badType,
	)

	var invalid Int64
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
		Int64{},
		invalid,
	)
}

func TestInt64UnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Int64
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Int64{},
		null,
	)
}

func TestInt64UnmarshalText(t *testing.T) {
	testData := newInt64Data()
	var nonzero Int64
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int64
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				valid: true,
			},
		},
		zero,
	)

	var null Int64
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{},
		null,
	)
}

func TestInt64MarshalJSON(t *testing.T) {
	testData := newInt64Data()
	nonzero := Int64From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewInt64(ZeroInt64, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt64(ZeroInt64, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestInt64MarshalText(t *testing.T) {
	testData := newInt64Data()
	nonzero := Int64From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewInt64(ZeroInt64, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewInt64(ZeroInt64, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestInt64Pointer(t *testing.T) {
	testData := newInt64Data()
	nonzero := Int64From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewInt64(ZeroInt64, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestInt64IsZero(t *testing.T) {
	testData := newInt64Data()
	nonzero := Int64From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewInt64(ZeroInt64, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewInt64(ZeroInt64, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestInt64SetValue(t *testing.T) {
	testData := newInt64Data()
	sut := NewInt64(ZeroInt64, false)
	assert.Equal(
		t,
		Int64{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestInt64Scan(t *testing.T) {
	testData := newInt64Data()
	var nonzero Int64
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Int64
	err = zero.Scan(ZeroInt64)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{
			NullableImpl: NullableImpl[int64]{
				valid: true,
			},
		},
		zero,
	)

	var null Int64
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Int64{},
		null,
	)
}
