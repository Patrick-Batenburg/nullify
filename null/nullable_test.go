package null

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNullableEqual(t *testing.T) {
	randomIn64 := gofakeit.Int64()
	randomString := gofakeit.LetterN(math.MaxInt8)

	nonzeroInteger := From(randomIn64)
	assert.True(t, nonzeroInteger.Equal(nonzeroInteger))
	assert.True(t, nonzeroInteger.Equal(From(randomIn64)))
	assert.False(t, nonzeroInteger.Equal(From(randomIn64+gofakeit.Int64())))

	nonzeroString := From(randomString)
	assert.True(t, nonzeroString.Equal(nonzeroString))
	assert.True(t, nonzeroString.Equal(From(randomString)))
	assert.False(t, nonzeroString.Equal(From(randomString+gofakeit.Letter())))
}

func TestNullableMustValue(t *testing.T) {
	value := gofakeit.Int64()
	nonzero := From(value)
	assert.Equal(t, value, nonzero.MustValue())

	zero := From(ZeroInt64)
	assert.Equal(t, ZeroInt64, zero.MustValue())

	assert.Panics(t, func() {
		invalid := New(ZeroInt64, false)
		assert.Equal(t, ZeroInt64, invalid.MustValue())
	})
}

type CircularReferenceNode struct {
	Value int
	Next  *CircularReferenceNode
}

func TestNullableMarshalText(t *testing.T) {
	testData := newJSONData()
	nonzeroJSON := From(testData.Bytes)
	data, err := nonzeroJSON.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzeroJSON,
	)

	nonzeroStruct := From(testData.Struct)
	data, err = nonzeroStruct.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzeroStruct,
	)

	nonzeroBool := From(true)
	data, err = nonzeroBool.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		TrueString,
		string(data),
		nonzeroStruct,
	)

	node := &CircularReferenceNode{Value: 1}
	node.Next = node
	invalid := From(node)
	_, err = invalid.MarshalText()
	var syntaxErr *json.UnsupportedValueError
	require.ErrorAs(
		t,
		err,
		&syntaxErr,
		"expected error to be of type *json.UnsupportedValueError",
	)
}

func TestNullableValue(t *testing.T) {
	value := gofakeit.Int64()
	nonzero := From(value)
	driverValue, err := nonzero.Value()
	require.NoError(t, err)
	assert.Equal(t, value, driverValue)

	zero := From(ZeroInt64)
	driverValue, err = zero.Value()
	require.NoError(t, err)
	assert.Equal(t, ZeroInt64, driverValue)

	invalid := New(gofakeit.Int64(), false)
	driverValue, err = invalid.Value()
	require.NoError(t, err)
	assert.Nil(t, driverValue)
}

func TestNullableValueOrZero(t *testing.T) {
	value := gofakeit.Int64()
	nonzero := From(value)
	valueOrZero := nonzero.ValueOrZero()
	assert.Equal(t, value, valueOrZero)

	zero := From(ZeroInt64)
	valueOrZero = zero.ValueOrZero()
	assert.Equal(t, ZeroInt64, valueOrZero)

	invalid := New(gofakeit.Int64(), false)
	valueOrZero = invalid.ValueOrZero()
	assert.Equal(t, ZeroInt64, valueOrZero)
	assert.NotEqual(t, ZeroInt64, invalid.value)
}

func TestNullableUnmarshalText(t *testing.T) {
	testData := newStringData()
	var invalidInt = From(ZeroInt)
	err := invalidInt.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidInt8 = From(ZeroInt8)
	err = invalidInt8.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidInt16 = From(ZeroInt16)
	err = invalidInt16.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidInt32 = From(ZeroInt32)
	err = invalidInt32.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidInt64 = From(ZeroInt64)
	err = invalidInt64.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidUint = From(ZeroUint)
	err = invalidUint.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidUint8 = From(ZeroUint8)
	err = invalidUint8.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidUint16 = From(ZeroUint16)
	err = invalidUint16.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidUint32 = From(ZeroUint32)
	err = invalidUint32.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidUint64 = From(ZeroUint64)
	err = invalidUint64.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidFloat32 = From(ZeroFloat32)
	err = invalidFloat32.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidFloat64 = From(ZeroFloat64)
	err = invalidFloat64.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidBool = From(ZeroBool)
	err = invalidBool.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	var invalidUUID = From(uuid.New())
	err = invalidUUID.UnmarshalText(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotUnmarshal)

	invalidUUID = From(uuid.New())
	err = invalidUUID.UnmarshalText([]byte(gofakeit.UUID()))
	require.ErrorIs(t, err, ErrCannotUnmarshal)
}

func TestNullableScan(t *testing.T) {
	testData := newStringData()
	var invalidInt = From(ZeroInt)
	err := invalidInt.Scan(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotScan)

	var invalidUUID = From(uuid.New())
	err = invalidUUID.Scan(testData.Bytes)
	require.ErrorIs(t, err, ErrCannotScan)

	validUUID := From(uuid.New())
	err = validUUID.Scan([]byte(gofakeit.UUID()))
	require.NoError(t, err)
}
