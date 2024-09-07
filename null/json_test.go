package null

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewJSON(t *testing.T) {
	testData := newJSONData()
	nonzero := NewJSON(testData.Value, true)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Value,
					valid: true,
				},
			},
		},
		nonzero,
	)

	zero := NewJSON(ZeroBytes, true)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					valid: true,
				},
			},
		},
		zero,
	)

	null := NewJSON(testData.Value, false)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Value,
				},
			},
		},
		null,
	)
}

func TestJSONFrom(t *testing.T) {
	testData := newJSONData()
	nonzero := JSONFrom(testData.Value)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Value,
					valid: true,
				},
			},
		},
		nonzero,
	)

	zero := JSONFrom(ZeroBytes)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					valid: true,
				},
			},
		},
		zero,
	)
}

func TestJSONFromPtr(t *testing.T) {
	testData := newJSONData()
	nonzero := JSONFromPtr(testData.Ptr)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Value,
					valid: true,
				},
			},
		},
		nonzero,
	)

	null := JSONFromPtr(nil)
	assert.Equal(
		t,
		JSON{},
		null,
	)
}

func TestJSONUnmarshalAny(t *testing.T) {
	testData := newJSONData()
	nonzero := JSONFrom(testData.Bytes)
	var jsonMap map[string]any

	err := nonzero.Unmarshal(&jsonMap)
	require.NoError(t, err)
	assert.Len(t, jsonMap, 7)

	assert.Equal(t, testData.Struct.ID.value.String(), jsonMap["id"])
	assert.Equal(t, testData.Struct.FirstName.value, jsonMap["firstName"])
	assert.Nil(t, jsonMap["middleName"])
	assert.Equal(t, testData.Struct.LastName.value, jsonMap["lastName"])
	assert.InEpsilon(t, testData.Struct.Age.value, jsonMap["age"], ZeroFloat64)
	assert.Equal(t, testData.Struct.Email.value, jsonMap["email"])
	assert.Equal(t, testData.Struct.CreatedAt.value.Format(time.RFC3339), jsonMap["createdAt"])

	null := NewJSON(NullStringBytes, false)
	jsonMap = nil
	err = null.Unmarshal(&jsonMap)
	require.NoError(t, err)
	assert.Nil(t, jsonMap)

	invalid := NewJSON(invalidJSON, false)
	jsonMap = nil
	err = invalid.Unmarshal(&jsonMap)
	var syntaxErr *json.SyntaxError
	require.ErrorAs(
		t,
		err,
		&syntaxErr,
		"expected error to be of type *json.SyntaxError",
	)
}

func TestJSONUnmarshalJSON(t *testing.T) {
	testData := newJSONData()
	var nonzero JSON
	err := json.Unmarshal(testData.Bytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Value,
					valid: true,
				},
			},
		},
		nonzero,
	)

	var null JSON
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{},
		null,
	)

	var invalid JSON
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
		JSON{},
		invalid,
	)
}

func TestJSONUnmarshalText(t *testing.T) {
	testData := newJSONData()
	var nonzero JSON
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Value,
					valid: true,
				},
			},
		},
		nonzero,
	)

	var null JSON
	err = null.UnmarshalText(ZeroBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{},
		null,
	)
}

func TestJSONMarshalAny(t *testing.T) {
	testData := newJSONData()
	nonzero := NewJSON(nil, false)
	err := nonzero.Marshal(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Bytes,
					valid: true,
				},
			},
		},
		nonzero,
	)

	nonzero = NewJSON(nil, false)
	err = nonzero.Marshal(testData.Struct)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Bytes,
					valid: true,
				},
			},
		},
		nonzero,
	)

	nonzero = NewJSON(nil, false)
	err = nonzero.Marshal(testData.String)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Bytes,
					valid: true,
				},
			},
		},
		nonzero,
	)

	zero := NewJSON(nil, false)
	err = zero.Marshal(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: ZeroBytes,
					valid: false,
				},
			},
		},
		zero,
	)

	zero = NewJSON(nil, false)
	err = zero.Marshal(ZeroString)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: ZeroBytes,
					valid: false,
				},
			},
		},
		zero,
	)

	null := NewJSON(nil, false)
	err = null.Marshal(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: ZeroBytes,
					valid: false,
				},
			},
		},
		null,
	)
}

func TestJSONMarshalJSON(t *testing.T) {
	testData := newJSONData()
	nonzero := JSONFrom(testData.Bytes)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	zero := NewJSON(ZeroStringBytes, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)

	null := NewJSON(ZeroBytes, false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestJSONMarshalText(t *testing.T) {
	testData := newJSONData()
	nonzero := JSONFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	zero := NewJSON(EmptyBytes, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)

	null := NewJSON(nil, false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestJSONPointer(t *testing.T) {
	testData := newJSONData()
	nonzero := JSONFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewJSON(testData.Bytes, false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestJSONIsZero(t *testing.T) {
	testData := newJSONData()
	nonzero := JSONFrom(testData.Value)
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewJSON(ZeroBytes, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewJSON(ZeroBytes, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestJSONSetValue(t *testing.T) {
	testData := newJSONData()
	sut := NewJSON(ZeroBytes, false)
	assert.Equal(
		t,
		JSON{},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Value,
					valid: true,
				},
			},
		},
		sut,
	)
}

func TestJSONScan(t *testing.T) {
	testData := newJSONData()
	var nonzero JSON
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: testData.Value,
					valid: true,
				},
			},
		},
		nonzero,
	)

	var emptyJSON JSON
	err = emptyJSON.Scan(ZeroString)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{
			Bytes: Bytes{
				NullableImpl: NullableImpl[[]byte]{
					value: EmptyBytes,
					valid: true,
				},
			},
		},
		emptyJSON,
	)

	var null JSON
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		JSON{},
		null,
	)
}
