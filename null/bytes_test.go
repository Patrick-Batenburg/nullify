package null

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBytes(t *testing.T) {
	testData := newBytesData()
	nonzero := NewBytes(testData.Value, true)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := NewBytes(ZeroBytes, true)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				valid: true,
			},
		},
		zero,
	)

	null := NewBytes(testData.Value, false)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestBytesFrom(t *testing.T) {
	testData := newBytesData()
	nonzero := BytesFrom(testData.Value)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zero := BytesFrom(ZeroBytes)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				valid: true,
			},
		},
		zero,
	)
}

func TestBytesFromPtr(t *testing.T) {
	testData := newBytesData()
	nonzero := BytesFromPtr(testData.Ptr)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	null := BytesFromPtr(nil)
	assert.Equal(
		t,
		Bytes{},
		null,
	)
}

func TestBytesUnmarshalJSON(t *testing.T) {
	testData := newBytesData()
	var nonzero Bytes
	err := json.Unmarshal(testData.JSONBytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var null Bytes
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bytes{},
		null,
	)

	var invalid Bytes
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
		Bytes{},
		invalid,
	)
}

func TestBytesUnmarshalText(t *testing.T) {
	testData := newBytesData()
	var nonzero Bytes
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var null Bytes
	err = null.UnmarshalText(ZeroBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bytes{},
		null,
	)
}

func TestBytesMarshalJSON(t *testing.T) {
	testData := newBytesData()
	nonzero := BytesFrom(testData.JSONBytes)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.JSONString,
		string(data),
	)

	zero := NewBytes(ZeroStringBytes, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)

	null := NewBytes(ZeroBytes, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestBytesMarshalText(t *testing.T) {
	testData := newBytesData()
	nonzero := BytesFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewBytes(EmptyBytes, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)

	null := NewBytes(nil, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestBytesPointer(t *testing.T) {
	testData := newBytesData()
	nonzero := BytesFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewBytes(testData.Bytes, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestBytesIsZero(t *testing.T) {
	testData := newBytesData()
	nonzero := BytesFrom(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewBytes(ZeroBytes, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewBytes(ZeroBytes, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestBytesSetValue(t *testing.T) {
	testData := newBytesData()
	sut := NewBytes(ZeroBytes, false)
	assert.Equal(
		t,
		Bytes{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestBytesScan(t *testing.T) {
	testData := newBytesData()
	var nonzero Bytes
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var emptyBytes Bytes
	err = emptyBytes.Scan(ZeroString)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bytes{
			NullableImpl: NullableImpl[[]byte]{
				value: EmptyBytes,
				valid: true,
			},
		},
		emptyBytes,
	)

	var null Bytes
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Bytes{},
		null,
	)
}
