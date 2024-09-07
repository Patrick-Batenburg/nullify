package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUint16(t *testing.T) {
	testData := newUint16Data()
	nonzero := NewUint16(testData.Value, true)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewUint16(0, true)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				valid: true,
			},
		},
		zero,
	)

	null := NewUint16(testData.Value, false)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestUint16From(t *testing.T) {
	testData := newUint16Data()
	nonzero := Uint16From(testData.Value)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Uint16From(0)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				valid: true,
			},
		},
		zero,
	)
}

func TestUint16FromPtr(t *testing.T) {
	testData := newUint16Data()
	nonzero := Uint16FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Uint16FromPtr(nil)
	assert.Equal(
		t,
		Uint16{},
		null,
	)
}

func TestUint16UnmarshalJSON(t *testing.T) {
	testData := newUint16Data()
	var nonzero Uint16
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint16
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint16
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{},
		null,
	)

	var badType Uint16
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Uint16{},
		badType,
	)

	var invalid Uint16
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
		Uint16{},
		invalid,
	)
}

func TestUint16UnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Uint16
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Uint16{},
		null,
	)
}

func TestUint16UnmarshalText(t *testing.T) {
	testData := newUint16Data()
	var nonzero Uint16
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint16
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint16
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{},
		null,
	)
}

func TestUint16MarshalJSON(t *testing.T) {
	testData := newUint16Data()
	nonzero := Uint16From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewUint16(0, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint16(0, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestUint16MarshalText(t *testing.T) {
	testData := newUint16Data()
	nonzero := Uint16From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewUint16(0, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint16(0, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestUint16Pointer(t *testing.T) {
	testData := newUint16Data()
	nonzero := Uint16From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewUint16(0, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestUint16IsZero(t *testing.T) {
	testData := newUint16Data()
	nonzero := Uint16From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewUint16(0, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewUint16(0, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestUint16SetValue(t *testing.T) {
	testData := newUint16Data()
	sut := NewUint16(0, false)
	assert.Equal(
		t,
		Uint16{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestUint16Scan(t *testing.T) {
	testData := newUint16Data()
	var nonzero Uint16
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint16
	err = zero.Scan(0)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{
			NullableImpl: NullableImpl[uint16]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint16
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint16{},
		null,
	)
}
