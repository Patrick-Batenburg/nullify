package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFloat64(t *testing.T) {
	testData := newFloat64Data()
	nonzero := NewFloat64(testData.Value, true)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewFloat64(ZeroFloat64, true)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				valid: true,
			},
		},
		zero,
	)

	null := NewFloat64(testData.Value, false)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestFloat64From(t *testing.T) {
	testData := newFloat64Data()
	nonzero := Float64From(testData.Value)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Float64From(ZeroFloat64)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				valid: true,
			},
		},
		zero,
	)
}

func TestFloat64FromPtr(t *testing.T) {
	testData := newFloat64Data()
	nonzero := Float64FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Float64FromPtr(nil)
	assert.Equal(
		t,
		Float64{},
		null,
	)
}

func TestFloat64UnmarshalJSON(t *testing.T) {
	testData := newFloat64Data()
	var nonzero Float64
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Float64
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				valid: true,
			},
		},
		zero,
	)

	var null Float64
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{},
		null,
	)

	var badType Float64
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Float64{},
		badType,
	)

	var invalid Float64
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
		Float64{},
		invalid,
	)
}

func TestFloat64UnmarshalText(t *testing.T) {
	testData := newFloat64Data()
	var nonzero Float64
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Float64
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				valid: true,
			},
		},
		zero,
	)

	var null Float64
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{},
		null,
	)
}

func TestFloat64MarshalJSON(t *testing.T) {
	testData := newFloat64Data()
	nonzero := Float64From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.JSONString,
		string(data),
	)

	zero := NewFloat64(ZeroFloat64, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewFloat64(ZeroFloat64, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestFloat64MarshalText(t *testing.T) {
	testData := newFloat64Data()
	nonzero := Float64From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewFloat64(ZeroFloat64, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewFloat64(ZeroFloat64, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestFloat64Pointer(t *testing.T) {
	testData := newFloat64Data()
	nonzero := Float64From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewFloat64(ZeroFloat64, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestFloat64IsZero(t *testing.T) {
	testData := newFloat64Data()
	nonzero := Float64From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewFloat64(ZeroFloat64, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewFloat64(ZeroFloat64, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestFloat64SetValue(t *testing.T) {
	testData := newFloat64Data()
	sut := NewFloat64(ZeroFloat64, false)
	assert.Equal(
		t,
		Float64{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestFloat64Scan(t *testing.T) {
	testData := newFloat64Data()
	var nonzero Float64
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Float64
	err = zero.Scan(ZeroFloat64)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{
			NullableImpl: NullableImpl[float64]{
				valid: true,
			},
		},
		zero,
	)

	var null Float64
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float64{},
		null,
	)
}
