package null

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUUID(t *testing.T) {
	testData := newUUIDData()
	nonzero := NewUUID(testData.Value, true)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	nonzero = NewRandomUUID()
	assert.NotEqual(t, uuid.Nil, nonzero.value)

	zeroFromString := NewUUID(ZeroUUIDString, true)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zeroFromString,
	)

	zeroFromStringer := NewUUID(uuid.Nil, true)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zeroFromStringer,
	)

	null := NewUUID(testData.Value, false)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				value: testData.Value,
			},
		},
		null,
	)
}

func TestUUIDFrom(t *testing.T) {
	testData := newUUIDData()
	nonzero := UUIDFrom(testData.Value)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zeroFromString := UUIDFrom(ZeroUUIDString)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zeroFromString,
	)

	zeroFromStringer := UUIDFrom(uuid.Nil)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zeroFromStringer,
	)
}

func TestUUIDFromPtr(t *testing.T) {
	testData := newUUIDData()
	nonzero := UUIDFromPtr(testData.Ptr)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	zeroFromString := UUIDFromPtr(&ZeroUUIDString)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zeroFromString,
	)

	zeroFromStringer := UUIDFromPtr(&uuid.Nil)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zeroFromStringer,
	)

	null := UUIDFromPtr(nil)
	assert.Equal(
		t,
		UUID{},
		null,
	)
}

func TestUUIDUnmarshalJSON(t *testing.T) {
	testData := newUUIDData()
	var nonzero UUID
	err := json.Unmarshal(testData.JSONBytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero UUID
	err = json.Unmarshal([]byte(strconv.Quote(ZeroUUIDString)), &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zero,
	)

	var null UUID
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{},
		null,
	)

	var badType UUID
	err = json.Unmarshal(ZeroIntegerStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		UUID{},
		badType,
	)

	var invalid UUID
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
		UUID{},
		invalid,
	)
}

func TestUUIDUnmarshalText(t *testing.T) {
	testData := newUUIDData()
	var nonzero UUID
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero UUID
	err = zero.UnmarshalText([]byte(ZeroUUIDString))
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zero,
	)

	var null UUID
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{},
		null,
	)
}

func TestUUIDMarshalJSON(t *testing.T) {
	testData := newUUIDData()
	nonzero := UUIDFrom(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.JSON,
		string(data),
	)

	zero := NewUUID(ZeroUUIDString, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		strconv.Quote(ZeroUUIDString),
		string(data),
	)

	null := NewUUID(gofakeit.UUID(), false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestUUIDMarshalText(t *testing.T) {
	testData := newUUIDData()
	nonzero := UUIDFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewUUID(ZeroUUIDString, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroUUIDString,
		string(data),
	)

	null := NewUUID(gofakeit.UUID(), false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestUUIDPointer(t *testing.T) {
	testData := newUUIDData()
	nonzero := UUIDFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewUUID(gofakeit.UUID(), false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestUUIDIsZero(t *testing.T) {
	nonzero := UUIDFrom(gofakeit.UUID())
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewUUID(ZeroUUIDString, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewUUID(ZeroUUIDString, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestUUIDSetValue(t *testing.T) {
	testData := newUUIDData()
	sut := NewUUID(ZeroUUIDString, false)
	assert.Equal(
		t,
		UUID{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				value: testData.Value,
				valid: true,
			},
		},
		sut,
	)
}

func TestUUIDScan(t *testing.T) {
	testData := newUUIDData()
	var nonzero UUID
	err := nonzero.Scan([]byte(testData.Value.String()))
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	var zero UUID
	err = zero.Scan(ZeroUUIDString)
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{
			NullableImpl: NullableImpl[uuid.UUID]{
				valid: true,
			},
		},
		zero,
	)

	var null UUID
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		UUID{},
		null,
	)
}
