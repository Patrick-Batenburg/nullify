package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUint(t *testing.T) {
	testData := newUintData()
	nonzero := NewUint(testData.Value, true)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewUint(0, true)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				valid: true,
			},
		},
		zero,
	)

	null := NewUint(testData.Value, false)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestUintFrom(t *testing.T) {
	testData := newUintData()
	nonzero := UintFrom(testData.Value)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := UintFrom(0)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				valid: true,
			},
		},
		zero,
	)
}

func TestUintFromPtr(t *testing.T) {
	testData := newUintData()
	nonzero := UintFromPtr(testData.Ptr)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := UintFromPtr(nil)
	assert.Equal(
		t,
		Uint{},
		null,
	)
}

func TestUintUnmarshalJSON(t *testing.T) {
	testData := newUintData()
	var nonzero Uint
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{},
		null,
	)

	var badType Uint
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Uint{},
		badType,
	)

	var invalid Uint
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
		Uint{},
		invalid,
	)
}

func TestUintUnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Uint
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Uint{},
		null,
	)
}

func TestUintUnmarshalText(t *testing.T) {
	testData := newUintData()
	var nonzero Uint
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{},
		null,
	)
}

func TestUintMarshalJSON(t *testing.T) {
	testData := newUintData()
	nonzero := UintFrom(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewUint(0, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint(0, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestUintMarshalText(t *testing.T) {
	testData := newUintData()
	nonzero := UintFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewUint(0, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint(0, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestUintPointer(t *testing.T) {
	testData := newUintData()
	nonzero := UintFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewUint(0, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestUintIsZero(t *testing.T) {
	testData := newUintData()
	nonzero := UintFrom(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewUint(0, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewUint(0, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestUintSetValue(t *testing.T) {
	testData := newUintData()
	sut := NewUint(0, false)
	assert.Equal(
		t,
		Uint{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestUintScan(t *testing.T) {
	testData := newUintData()
	var nonzero Uint
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint
	err = zero.Scan(0)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{
			NullableImpl: NullableImpl[uint]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint{},
		null,
	)
}
