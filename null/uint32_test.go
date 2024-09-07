package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUint32(t *testing.T) {
	testData := newUint32Data()
	nonzero := NewUint32(testData.Value, true)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewUint32(0, true)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				valid: true,
			},
		},
		zero,
	)

	null := NewUint32(testData.Value, false)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestUint32From(t *testing.T) {
	testData := newUint32Data()
	nonzero := Uint32From(testData.Value)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Uint32From(0)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				valid: true,
			},
		},
		zero,
	)
}

func TestUint32FromPtr(t *testing.T) {
	testData := newUint32Data()
	nonzero := Uint32FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Uint32FromPtr(nil)
	assert.Equal(
		t,
		Uint32{},
		null,
	)
}

func TestUint32UnmarshalJSON(t *testing.T) {
	testData := newUint32Data()
	var nonzero Uint32
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint32
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint32
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{},
		null,
	)

	var badType Uint32
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Uint32{},
		badType,
	)

	var invalid Uint32
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
		Uint32{},
		invalid,
	)
}

func TestUint32UnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Uint32
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Uint32{},
		null,
	)
}

func TestUint32UnmarshalText(t *testing.T) {
	testData := newUint32Data()
	var nonzero Uint32
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint32
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint32
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{},
		null,
	)
}

func TestUint32MarshalJSON(t *testing.T) {
	testData := newUint32Data()
	nonzero := Uint32From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewUint32(0, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint32(0, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestUint32MarshalText(t *testing.T) {
	testData := newUint32Data()
	nonzero := Uint32From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewUint32(0, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint32(0, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestUint32Pointer(t *testing.T) {
	testData := newUint32Data()
	nonzero := Uint32From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewUint32(0, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestUint32IsZero(t *testing.T) {
	testData := newUint32Data()
	nonzero := Uint32From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewUint32(0, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewUint32(0, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestUint32SetValue(t *testing.T) {
	testData := newUint32Data()
	sut := NewUint32(0, false)
	assert.Equal(
		t,
		Uint32{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestUint32Scan(t *testing.T) {
	testData := newUint32Data()
	var nonzero Uint32
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint32
	err = zero.Scan(0)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{
			NullableImpl: NullableImpl[uint32]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint32
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint32{},
		null,
	)
}
