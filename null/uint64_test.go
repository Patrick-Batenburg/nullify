package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUint64(t *testing.T) {
	testData := newUint64Data()
	nonzero := NewUint64(testData.Value, true)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewUint64(0, true)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				valid: true,
			},
		},
		zero,
	)

	null := NewUint64(testData.Value, false)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestUint64From(t *testing.T) {
	testData := newUint64Data()
	nonzero := Uint64From(testData.Value)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := Uint64From(0)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				valid: true,
			},
		},
		zero,
	)
}

func TestUint64FromPtr(t *testing.T) {
	testData := newUint64Data()
	nonzero := Uint64FromPtr(testData.Ptr)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := Uint64FromPtr(nil)
	assert.Equal(
		t,
		Uint64{},
		null,
	)
}

func TestUint64UnmarshalJSON(t *testing.T) {
	testData := newUint64Data()
	var nonzero Uint64
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint64
	err = json.Unmarshal(ZeroIntegerStringBytes, &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint64
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{},
		null,
	)

	var badType Uint64
	err = json.Unmarshal(TrueStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Uint64{},
		badType,
	)

	var invalid Uint64
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
		Uint64{},
		invalid,
	)
}

func TestUint64UnmarshalJSONNonIntegerNumber(t *testing.T) {
	testData := newFloat64Data()
	var null Uint64
	err := json.Unmarshal(testData.Bytes, &null)
	require.Error(t, err)
	assert.Equal(
		t,
		Uint64{},
		null,
	)
}

func TestUint64UnmarshalText(t *testing.T) {
	testData := newUint64Data()
	var nonzero Uint64
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint64
	err = zero.UnmarshalText(ZeroIntegerStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint64
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{},
		null,
	)
}

func TestUint64MarshalJSON(t *testing.T) {
	testData := newUint64Data()
	nonzero := Uint64From(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewUint64(0, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint64(0, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestUint64MarshalText(t *testing.T) {
	testData := newUint64Data()
	nonzero := Uint64From(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewUint64(0, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroIntegerString,
		string(data),
	)

	null := NewUint64(0, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestUint64Pointer(t *testing.T) {
	testData := newUint64Data()
	nonzero := Uint64From(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewUint64(0, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestUint64IsZero(t *testing.T) {
	testData := newUint64Data()
	nonzero := Uint64From(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewUint64(0, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewUint64(0, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestUint64SetValue(t *testing.T) {
	testData := newUint64Data()
	sut := NewUint64(0, false)
	assert.Equal(
		t,
		Uint64{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestUint64Scan(t *testing.T) {
	testData := newUint64Data()
	var nonzero Uint64
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero Uint64
	err = zero.Scan(0)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{
			NullableImpl: NullableImpl[uint64]{
				valid: true,
			},
		},
		zero,
	)

	var null Uint64
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Uint64{},
		null,
	)
}
