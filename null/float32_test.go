package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFloat32(t *testing.T) {
	testData := newFloat32Data()
	nonzero := NewFloat32(testData.Value, true)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewFloat32(ZeroFloat32, true)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				valid: true,
			},
		},
		zero,
	)

	null := NewFloat32(testData.Value, false)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestFloat32From(t *testing.T) {
	testData := newFloat32Data()
	nonzero := Float32From(testData.Value)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Float32From(ZeroFloat32)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				valid: true,
			},
		},
		zero,
	)
}

func TestFloat32FromPtr(t *testing.T) {
	testData := newFloat32Data()
	nonzero := Float32FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Float32FromPtr(nil)
	assert.Equal(
		t,
		Float32{},
		null,
	)
}

func TestFloat32UnmarshalJSON(t *testing.T) {
	testData := newFloat32Data()
	var nonzero Float32
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Float32
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				valid: true,
			},
		},
		zero,
	)

	var null Float32
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{},
		null,
	)

	var badType Float32
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Float32{},
		badType,
	)

	var invalid Float32
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
		Float32{},
		invalid,
	)
}

func TestFloat32UnmarshalText(t *testing.T) {
	testData := newFloat32Data()
	var nonzero Float32
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Float32
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				valid: true,
			},
		},
		zero,
	)

	var null Float32
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{},
		null,
	)
}

func TestFloat32MarshalJSON(t *testing.T) {
	testData := newFloat32Data()
	nonzero := Float32From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.JSONString,
		string(data),
	)

	zero := NewFloat32(ZeroFloat32, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewFloat32(ZeroFloat32, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestFloat32MarshalText(t *testing.T) {
	testData := newFloat32Data()
	nonzero := Float32From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewFloat32(ZeroFloat32, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewFloat32(ZeroFloat32, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestFloat32Pointer(t *testing.T) {
	testData := newFloat32Data()
	nonzero := Float32From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewFloat32(ZeroFloat32, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestFloat32IsZero(t *testing.T) {
	testData := newFloat32Data()
	nonzero := Float32From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewFloat32(ZeroFloat32, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewFloat32(ZeroFloat32, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestFloat32SetValue(t *testing.T) {
	testData := newFloat32Data()
	sut := NewFloat32(ZeroFloat32, false)
	assert.Equal(
		t,
		Float32{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestFloat32Scan(t *testing.T) {
	testData := newFloat32Data()
	var nonzero Float32
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Float32
	err = zero.Scan(ZeroFloat32)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{
			NullableImpl: NullableImpl[float32]{
				valid: true,
			},
		},
		zero,
	)

	var null Float32
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Float32{},
		null,
	)
}
