package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

var (
	invalidJSON = []byte(`:)`)
)

type BoolData struct {
	Value  bool
	Ptr    *bool
	String string
	Bytes  []byte
}

func newBoolData() BoolData {
	value := gofakeit.Bool()
	str := fmt.Sprintf("%t", value)

	return BoolData{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type ByteData struct {
	Value      byte
	Ptr        *byte
	String     string
	JSONString string
	JSONBytes  []byte
	Bytes      []byte
}

func newByteData() ByteData {
	value := byte(gofakeit.Letter()[0])
	str := string(value)
	JSONString := strconv.Quote(str)

	return ByteData{
		Value:      value,
		Ptr:        &value,
		String:     str,
		JSONString: JSONString,
		JSONBytes:  []byte(JSONString),
		Bytes:      []byte(str),
	}
}

type BytesData struct {
	Value      []byte
	Ptr        *[]byte
	String     string
	JSONString string
	JSONBytes  []byte
	Bytes      []byte
}

func newBytesData() BytesData {
	value := []byte(gofakeit.LetterN(math.MaxInt8))
	str := string(value)
	JSONString := strconv.Quote(str)

	return BytesData{
		Value:      value,
		Ptr:        &value,
		String:     str,
		JSONString: JSONString,
		JSONBytes:  []byte(JSONString),
		Bytes:      []byte(str),
	}
}

type Float32Data struct {
	Value      float32
	Ptr        *float32
	String     string
	JSONString string
	Bytes      []byte
}

func newFloat32Data() Float32Data {
	value := gofakeit.Float32Range(0.01, math.MaxFloat32)
	jsonBytes, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	str := string(jsonBytes)

	return Float32Data{
		Value:      value,
		Ptr:        &value,
		String:     strconv.FormatFloat(float64(value), 'f', -1, 32),
		JSONString: str,
		Bytes:      []byte(str),
	}
}

type Float64Data struct {
	Value      float64
	Ptr        *float64
	String     string
	JSONString string
	Bytes      []byte
}

func newFloat64Data() Float64Data {
	value := gofakeit.Float64Range(math.MaxFloat32, math.MaxFloat64)
	jsonBytes, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	str := string(jsonBytes)

	return Float64Data{
		Value:      value,
		Ptr:        &value,
		String:     strconv.FormatFloat(float64(value), 'f', -1, 64),
		JSONString: str,
		Bytes:      []byte(str),
	}
}

type IntData struct {
	Value  int
	Ptr    *int
	String string
	Bytes  []byte
}

func newIntData() IntData {
	value := gofakeit.Int()
	str := fmt.Sprintf("%d", value)

	return IntData{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type Int8Data struct {
	Value  int8
	Ptr    *int8
	String string
	Bytes  []byte
}

func newInt8Data() Int8Data {
	var value int8

	for {
		value = gofakeit.Int8()

		if value != 0 {
			break
		}
	}

	str := fmt.Sprintf("%d", value)

	return Int8Data{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type Int16Data struct {
	Value  int16
	Ptr    *int16
	String string
	Bytes  []byte
}

func newInt16Data() Int16Data {
	var value int16

	if gofakeit.Bool() {
		value = int16(gofakeit.IntRange(math.MaxInt8, math.MaxInt16))
	} else {
		value = int16(gofakeit.IntRange(math.MinInt8, math.MinInt16))
	}

	str := fmt.Sprintf("%d", value)

	return Int16Data{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type Int32Data struct {
	Value  int32
	Ptr    *int32
	String string
	Bytes  []byte
}

func newInt32Data() Int32Data {
	var value int32

	if gofakeit.Bool() {
		value = int32(gofakeit.IntRange(math.MaxInt16, math.MaxInt32))
	} else {
		value = int32(gofakeit.IntRange(math.MinInt16, math.MinInt32))
	}

	str := fmt.Sprintf("%d", value)

	return Int32Data{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type Int64Data struct {
	Value  int64
	Ptr    *int64
	String string
	Bytes  []byte
}

func newInt64Data() Int64Data {
	var value int64

	if gofakeit.Bool() {
		value = int64(gofakeit.IntRange(math.MaxInt32, math.MaxInt64))
	} else {
		value = int64(gofakeit.IntRange(math.MinInt32, math.MinInt64))
	}

	str := fmt.Sprintf("%d", value)

	return Int64Data{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type StringData struct {
	Value      string
	Ptr        *string
	JSONString string
	Bytes      []byte
	JSONBytes  []byte
}

func newStringData() StringData {
	value := gofakeit.LetterN(math.MaxInt8)
	JSONString := strconv.Quote(value)

	return StringData{
		Value:      value,
		Ptr:        &value,
		JSONString: JSONString,
		Bytes:      []byte(value),
		JSONBytes:  []byte(JSONString),
	}
}

type User struct {
	ID         UUID   `json:"id"`
	FirstName  String `json:"firstName"`
	MiddleName String `json:"middleName"`
	LastName   String `json:"lastName"`
	Age        Int8   `json:"age"`
	Email      String `json:"email"`
	CreatedAt  Time   `json:"createdAt"`
}

type JSONData struct {
	Value  []byte
	Ptr    *[]byte
	String string
	Struct User
	Bytes  []byte
}

func newJSONData() JSONData {
	var user User
	user.ID.SetValue(uuid.New())
	user.FirstName.SetValue(gofakeit.FirstName())
	user.LastName.SetValue(gofakeit.LastName())
	user.Age.SetValue(int8(gofakeit.IntRange(0, 100)))
	user.Email.SetValue(gofakeit.Email())
	user.CreatedAt.SetValue(time.Unix(gofakeit.Date().Unix(), 0))

	data := fmt.Sprintf(`{
  "id": "%s",
  "firstName": "%s",
  "middleName": null,
  "lastName": "%s",
  "age": %d,
  "email": "%s",
  "createdAt": "%s"
}`,
		user.ID.value,
		user.FirstName.value,
		user.LastName.value,
		user.Age.value,
		user.Email.value,
		user.CreatedAt.value.Format(time.RFC3339),
	)
	data = strings.ReplaceAll(strings.ReplaceAll(data, "\n", ""), " ", "")
	value := []byte(data)

	return JSONData{
		Value:  value,
		Ptr:    &value,
		String: data,
		Struct: user,
		Bytes:  []byte(data),
	}
}

type TimeData struct {
	Value     time.Time
	Ptr       *time.Time
	String    string
	JSON      string
	Bytes     []byte
	JSONBytes []byte
}

func newTimeData() TimeData {
	value := time.Unix(gofakeit.Date().Unix(), 0).UTC()
	str := value.Format(time.RFC3339)

	return TimeData{
		Value:     value,
		Ptr:       &value,
		String:    str,
		JSON:      strconv.Quote(string(value.Format(time.RFC3339))),
		Bytes:     []byte(str),
		JSONBytes: []byte(strconv.Quote(str)),
	}
}

type UintData struct {
	Value  uint
	Ptr    *uint
	String string
	Bytes  []byte
}

func newUintData() UintData {
	value := gofakeit.UintRange(1, math.MaxUint64)
	str := fmt.Sprintf("%d", value)

	return UintData{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type Uint8Data struct {
	Value  uint8
	Ptr    *uint8
	String string
	Bytes  []byte
}

func newUint8Data() Uint8Data {
	value := uint8(gofakeit.UintRange(1, math.MaxUint8))
	str := fmt.Sprintf("%d", value)

	return Uint8Data{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type Uint16Data struct {
	Value  uint16
	Ptr    *uint16
	String string
	Bytes  []byte
}

func newUint16Data() Uint16Data {
	value := uint16(gofakeit.UintRange(math.MaxUint8, math.MaxUint16))
	str := fmt.Sprintf("%d", value)

	return Uint16Data{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type Uint32Data struct {
	Value  uint32
	Ptr    *uint32
	String string
	Bytes  []byte
}

func newUint32Data() Uint32Data {
	value := uint32(gofakeit.UintRange(math.MaxUint16, math.MaxUint32))
	str := fmt.Sprintf("%d", value)

	return Uint32Data{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type Uint64Data struct {
	Value  uint64
	Ptr    *uint64
	String string
	Bytes  []byte
}

func newUint64Data() Uint64Data {
	value := uint64(gofakeit.UintRange(math.MaxUint32, math.MaxUint64))
	str := fmt.Sprintf("%d", value)

	return Uint64Data{
		Value:  value,
		Ptr:    &value,
		String: str,
		Bytes:  []byte(str),
	}
}

type UUIDData struct {
	Value     uuid.UUID
	Ptr       *uuid.UUID
	String    string
	JSON      string
	Bytes     []byte
	JSONBytes []byte
}

func newUUIDData() UUIDData {
	value := uuid.New()
	str := value.String()

	return UUIDData{
		Value:     value,
		Ptr:       &value,
		String:    str,
		JSON:      strconv.Quote(str),
		Bytes:     []byte(str),
		JSONBytes: []byte(strconv.Quote(str))}
}

type integerResult struct {
	actual   any
	expected driver.Value
}

type integerResultFn = func(*integerResult)

func withActualAndExpectedInteger(t *testing.T, actualType any, expectedType any) integerResultFn {
	t.Helper()

	return func(option *integerResult) {
		t.Helper()
		var randomInt int

		var minValue, maxValue int

		// Determine the smallest and largest possible values to fit both actualType and expectedType
		switch actualType.(type) {
		case int8:
			minValue = math.MinInt8
			maxValue = math.MaxInt8
		case int16:
			minValue = math.MinInt16
			maxValue = math.MaxInt16
		case int32:
			minValue = math.MinInt32
			maxValue = math.MaxInt32
		case int64, int:
			minValue = math.MinInt
			maxValue = math.MaxInt
		case uint8:
			minValue = 0
			maxValue = math.MaxUint8
		case uint16:
			minValue = 0
			maxValue = math.MaxUint16
		case uint32:
			minValue = 0
			maxValue = int(math.MaxUint32)
		case uint64, uint:
			minValue = 0
			maxValue = math.MaxInt
		default:
			t.Fatalf("invalid actual type")
		}

		// Further limit the maxVal based on the expectedType
		switch expectedType.(type) {
		case int8:
			if maxValue > math.MaxInt8 {
				maxValue = math.MaxInt8
			}
			if minValue < math.MinInt8 {
				minValue = math.MinInt8
			}
		case int16:
			if maxValue > math.MaxInt16 {
				maxValue = math.MaxInt16
			}
			if minValue < math.MinInt16 {
				minValue = math.MinInt16
			}
		case int32:
			if maxValue > math.MaxInt32 {
				maxValue = math.MaxInt32
			}
			if minValue < math.MinInt32 {
				minValue = math.MinInt32
			}
		case int64, int:
			// No need to adjust minValue or maxValue for int64 or int
		case uint8:
			if maxValue > math.MaxUint8 {
				maxValue = math.MaxUint8
			}
			if minValue < 0 {
				minValue = 0
			}
		case uint16:
			if maxValue > math.MaxUint16 {
				maxValue = math.MaxUint16
			}
			if minValue < 0 {
				minValue = 0
			}
		case uint32:
			if maxValue > int(math.MaxUint32) {
				maxValue = int(math.MaxUint32)
			}
			if minValue < 0 {
				minValue = 0
			}
		case uint64, uint:
			if minValue < 0 {
				minValue = 0
			}
		default:
			t.Fatalf("invalid expected type")
		}

		randomInt = gofakeit.IntRange(minValue, maxValue)

		switch actualType.(type) {
		case int:
			option.actual = randomInt
		case int8:
			option.actual = int8(randomInt)
		case int16:
			option.actual = int16(randomInt)
		case int32:
			option.actual = int32(randomInt)
		case int64:
			option.actual = int64(randomInt)
		case uint:
			option.actual = uint(randomInt)
		case uint8:
			option.actual = uint8(randomInt)
		case uint16:
			option.actual = uint16(randomInt)
		case uint32:
			option.actual = uint32(randomInt)
		case uint64:
			option.actual = uint64(randomInt)
		default:
			t.Fatalf("invalid actual type")
		}

		switch expectedType.(type) {
		case int:
			option.expected = randomInt
		case int8:
			option.expected = int8(randomInt)
		case int16:
			option.expected = int16(randomInt)
		case int32:
			option.expected = int32(randomInt)
		case int64:
			option.expected = int64(randomInt)
		case uint:
			option.expected = uint(randomInt)
		case uint8:
			option.expected = uint8(randomInt)
		case uint16:
			option.expected = uint16(randomInt)
		case uint32:
			option.expected = uint32(randomInt)
		case uint64:
			option.expected = uint64(randomInt)
		default:
			t.Fatalf("invalid expected type")
		}
	}
}
