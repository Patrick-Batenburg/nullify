package null

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/itlightning/dateparse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTime(t *testing.T) {
	testData := newTimeData()
	nonzero := NewTime(testData.Value, true)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		nonzero,
	)

	nonzeroWithLayout := NewTime(testData.Value, true, WithTimeLayout(time.DateTime))
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.DateTime,
			isStrictLayout: true,
		},
		nonzeroWithLayout,
	)

	nonzeroWithLenientParsing := NewTime(testData.Value, true, WithTimeLenientParsing())
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: false,
		},
		nonzeroWithLenientParsing,
	)

	zero := NewTime(ZeroTime, true)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		zero,
	)

	null := NewTime(testData.Value, false)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		null,
	)
}

func TestTimeFrom(t *testing.T) {
	testData := newTimeData()
	nonzero := TimeFrom(testData.Value)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		nonzero,
	)

	nonzeroWithLayout := TimeFrom(testData.Value, WithTimeLayout(time.DateTime))
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.DateTime,
			isStrictLayout: true,
		},
		nonzeroWithLayout,
	)

	nonzeroWithLenientParsing := TimeFrom(testData.Value, WithTimeLenientParsing())
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: false,
		},
		nonzeroWithLenientParsing,
	)

	zero := TimeFrom(ZeroTime)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		zero,
	)
}

func TestTimeFromPtr(t *testing.T) {
	testData := newTimeData()
	nonzero := TimeFromPtr(testData.Ptr)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		nonzero,
	)

	nonzeroWithLayout := TimeFromPtr(testData.Ptr, WithTimeLayout(time.DateTime))
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.DateTime,
			isStrictLayout: true,
		},
		nonzeroWithLayout,
	)

	nonzeroWithLenientParsing := TimeFromPtr(testData.Ptr, WithTimeLenientParsing())
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: false,
		},
		nonzeroWithLenientParsing,
	)

	null := TimeFromPtr(nil)
	assert.Equal(
		t,
		Time{
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		null,
	)
}

func TestTimeUnmarshalJSON(t *testing.T) {
	testData := newTimeData()
	var nonzero Time
	err := json.Unmarshal(testData.JSONBytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	nonzero = NewTime(ZeroTime, false, WithTimeStrictParsing())
	err = json.Unmarshal(testData.JSONBytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		nonzero,
	)

	nonzero = NewTime(ZeroTime, false, WithTimeLenientParsing())
	err = json.Unmarshal(testData.JSONBytes, &nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout: time.RFC3339,
		},
		nonzero,
	)

	var zero Time
	err = json.Unmarshal([]byte(strconv.Quote(ZeroTime.String())), &zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				valid: true,
			},
		},
		zero,
	)

	var null Time
	err = json.Unmarshal(NullStringBytes, &null)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{},
		null,
	)

	var badType Time
	err = json.Unmarshal(ZeroIntegerStringBytes, &badType)
	require.ErrorIs(t, err, ErrCannotUnmarshal)
	assert.Equal(
		t,
		Time{},
		badType,
	)

	var invalid Time
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
		Time{},
		invalid,
	)
}

func TestTimeUnmarshalText(t *testing.T) {
	testData := newTimeData()
	var nonzero Time
	err := nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	nonzero = NewTime(ZeroTime, false, WithTimeStrictParsing())
	err = nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		nonzero,
	)

	nonzero = NewTime(ZeroTime, false, WithTimeLenientParsing())
	err = nonzero.UnmarshalText(testData.Bytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout: time.RFC3339,
		},
		nonzero,
	)

	var zero Time
	err = zero.UnmarshalText([]byte(ZeroTime.String()))
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				valid: true,
			},
		},
		zero,
	)

	var null Time
	err = null.UnmarshalText(ZeroStringBytes)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{},
		null,
	)
}

func TestTimeMarshalJSON(t *testing.T) {
	testData := newTimeData()
	nonzero := TimeFrom(testData.Value)
	data, err := json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.JSON,
		string(data),
	)

	nonzero = TimeFrom(testData.Value, WithTimeDefaultLayout())
	data, err = json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.JSON,
		string(data),
	)

	nonzero = TimeFrom(testData.Value, WithTimeLayout(time.DateOnly))
	data, err = json.Marshal(nonzero)
	require.NoError(t, err)
	assert.Equal(
		t,
		strconv.Quote(testData.Value.Format(time.DateOnly)),
		string(data),
	)

	zero := NewTime(ZeroTime, true)
	data, err = json.Marshal(zero)
	require.NoError(t, err)
	assert.Equal(
		t,
		strconv.Quote(ZeroTime.Format(time.RFC3339)),
		string(data),
	)

	null := NewTime(gofakeit.Date(), false)
	data, err = json.Marshal(null)
	require.NoError(t, err)
	assert.Equal(
		t,
		NullString,
		string(data),
	)
}

func TestTimeMarshalText(t *testing.T) {
	testData := newTimeData()
	nonzero := TimeFrom(testData.Value)
	data, err := nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
		nonzero,
	)

	nonzero = TimeFrom(testData.Value, WithTimeDefaultLayout())
	data, err = nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.String,
		string(data),
	)

	nonzero = TimeFrom(testData.Value, WithTimeLayout(time.DateOnly))
	data, err = nonzero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		testData.Value.Format(time.DateOnly),
		string(data),
	)

	zero := NewTime(ZeroTime, true)
	data, err = zero.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroTime.Format(time.RFC3339),
		string(data),
	)

	null := NewTime(gofakeit.Date(), false)
	data, err = null.MarshalText()
	require.NoError(t, err)
	assert.Equal(
		t,
		ZeroString,
		string(data),
	)
}

func TestTimePointer(t *testing.T) {
	testData := newTimeData()
	nonzero := TimeFrom(testData.Value)
	ptr := nonzero.Ptr()
	assert.Equal(
		t,
		testData.Ptr,
		ptr,
		nonzero,
	)

	null := NewTime(gofakeit.Date(), false)
	ptr = null.Ptr()
	assert.Nil(
		t,
		ptr,
	)
}

func TestTimeIsZero(t *testing.T) {
	nonzero := TimeFrom(gofakeit.Date())
	assert.False(
		t,
		nonzero.IsZero(),
	)

	zero := NewTime(ZeroTime, true)
	assert.True(
		t,
		zero.IsZero(),
	)

	null := NewTime(ZeroTime, false)
	assert.True(
		t,
		null.IsZero(),
	)
}

func TestTimeSetValue(t *testing.T) {
	testData := newTimeData()
	sut := NewTime(ZeroTime, false)
	assert.Equal(
		t,
		Time{
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		sut,
	)

	sut.SetValue(testData.Value)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		sut,
	)
}

func TestTimeScan(t *testing.T) {
	testData := newTimeData()
	var nonzero Time
	err := nonzero.Scan(testData.Value)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
		},
		nonzero,
	)

	dateparseOptions := []dateparse.ParserOption{dateparse.PreferMonthFirst(false)}

	nonzero = NewTime(ZeroTime, false, WithTimeParseOptions(dateparseOptions...))
	err = nonzero.Scan(testData.Value.Format(time.RFC3339))
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
			parseOptions:   dateparseOptions,
		},
		nonzero,
	)

	nonzero = NewTime(ZeroTime, false)
	err = nonzero.Scan(testData.Value.Format(time.RFC3339))
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		nonzero,
	)

	nonzero = NewTime(ZeroTime, false)
	err = nonzero.Scan([]byte(testData.Value.Format(time.RFC3339)))
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		nonzero,
	)

	nonzero = NewTime(ZeroTime, false)
	err = nonzero.Scan(testData.Value.Unix())
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: testData.Value,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		nonzero,
	)

	var zero Time
	err = zero.Scan(ZeroTime)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				valid: true,
			},
		},
		zero,
	)

	zero = NewTime(ZeroTime, false)
	err = zero.Scan(ZeroTime.Format(time.RFC3339))
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: ZeroTime,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		zero,
	)

	zero = NewTime(ZeroTime, false)
	err = zero.Scan([]byte(ZeroTime.Format(time.RFC3339)))
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: ZeroTime,
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		zero,
	)

	zero = NewTime(ZeroTime, false)
	err = zero.Scan(ZeroTime.Unix())
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{
			NullableImpl: NullableImpl[time.Time]{
				value: ZeroTime.UTC(),
				valid: true,
			},
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		zero,
	)

	var null Time
	err = null.Scan(nil)
	require.NoError(t, err)
	assert.Equal(
		t,
		Time{},
		null,
	)

	var invalid Time
	err = invalid.Scan(gofakeit.LetterN(math.MaxInt8))
	require.ErrorIs(t, err, ErrCannotScan)
	assert.Equal(
		t,
		Time{},
		invalid,
	)

	invalid = NewTime(ZeroTime, false)
	err = invalid.Scan([]byte(gofakeit.LetterN(math.MaxInt8)))
	require.ErrorIs(t, err, ErrCannotScan)
	assert.Equal(
		t,
		Time{
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		invalid,
	)

	invalid = NewTime(ZeroTime, false)
	err = invalid.Scan(ZeroBool)
	require.ErrorIs(t, err, ErrCannotScan)
	assert.Equal(
		t,
		Time{
			layout:         time.RFC3339,
			isStrictLayout: true,
		},
		invalid,
	)
}
