package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUint8(t *testing.T) {
	testData := newUint8Data()
	nonzero := NewUint8(testData.Value, true)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewUint8(0, true)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				valid: true,
			},
		},
		zero,
	)

	null := NewUint8(testData.Value, false)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestUint8From(t *testing.T) {
	testData := newUint8Data()
	nonzero := Uint8From(testData.Value)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Uint8From(0)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				valid: true,
			},
		},
		zero,
	)
}

func TestUint8FromPtr(t *testing.T) {
	testData := newUint8Data()
	nonzero := Uint8FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Uint8FromPtr(nil)
	assert.Equal(
		t,
		Uint8{},
		null,
	)
}

func TestUint8UnmarshalJSON(t *testing.T) {
	testData := newUint8Data()
	var nonzero Uint8
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint8
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint8
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{},
		null,
	)

	var badType Uint8
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Uint8{},
		badType,
	)

	var invalid Uint8
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
		Uint8{},
		invalid,
	)
}

func TestUint8UnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Uint8
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Uint8{},
		null,
	)
}

func TestUint8UnmarshalText(t *testing.T) {
	testData := newUint8Data()
	var nonzero Uint8
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint8
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint8
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{},
		null,
	)
}

func TestUint8MarshalJSON(t *testing.T) {
	testData := newUint8Data()
	nonzero := Uint8From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewUint8(0, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint8(0, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestUint8MarshalText(t *testing.T) {
	testData := newUint8Data()
	nonzero := Uint8From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewUint8(0, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint8(0, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestUint8Pointer(t *testing.T) {
	testData := newUint8Data()
	nonzero := Uint8From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewUint8(0, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestUint8IsZero(t *testing.T) {
	testData := newUint8Data()
	nonzero := Uint8From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewUint8(0, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewUint8(0, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestUint8SetValue(t *testing.T) {
	testData := newUint8Data()
	sut := NewUint8(0, false)
	assert.Equal(
		t,
		Uint8{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestUint8Scan(t *testing.T) {
	testData := newUint8Data()
	var nonzero Uint8
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint8
	err = zero.Scan(0)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{
			NullableImpl: NullableImpl[uint8]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint8
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint8{},
		null,
	)
}
