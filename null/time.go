package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"bytes"
	"strconv"
	"time"

	"github.com/itlightning/dateparse"
)

// DateParsePreferMonthFirst is an option that allows preferMonthFirst to be changed from its default
var DateParsePreferMonthFirst = true

// Time is a NullableImpl time.Time. It supports SQL and JSON serialization.
type Time struct {
	NullableImpl[time.Time]

	// layout determines how the Time should be formatted when driver.Valuer,
	// json.Marshaler and encoding.TextMarshaler are called.
	layout string

	// isStrictLayout determines if the layout should be strictly or leniently parsed
	// when sql.Scanner, json.Unmarshaler and encoding.TextUnmarshaler are called.
	isStrictLayout bool

	// parseOptions is a slice of parse options when sql.Scanner,
	// json.Unmarshaler and encoding.TextUnMarshaler are called.
	parseOptions []dateparse.ParserOption
}

// NewTime creates a new Time with RFC3339 layout.
func NewTime(value time.Time, valid bool, options ...TimeOptionFn) Time {
	n := &Time{
		NullableImpl:   New(value, valid),
		layout:         time.RFC3339,
		isStrictLayout: true,
	}

	for _, option := range options {
		option(n)
	}

	return *n
}

// TimeFrom creates a new Time with RFC3339 layout that will always be valid.
func TimeFrom(value time.Time, options ...TimeOptionFn) Time {
	n := &Time{
		NullableImpl:   From(value),
		layout:         time.RFC3339,
		isStrictLayout: true,
	}

	for _, option := range options {
		option(n)
	}

	return *n
}

// TimeFromPtr creates a new Time  with RFC3339 layout that will be null if the value is nil.
func TimeFromPtr(value *time.Time, options ...TimeOptionFn) Time {
	n := &Time{
		NullableImpl:   FromPtr(value),
		layout:         time.RFC3339,
		isStrictLayout: true,
	}

	for _, option := range options {
		option(n)
	}

	return *n
}

// MarshalJSON implements json.Marshaler.
func (n Time) MarshalJSON() ([]byte, error) {
	if !n.IsValid() {
		return NullStringBytes, nil
	}

	return []byte(strconv.Quote(n.Format(WithTimeLayoutFormat()))), nil
}

// MarshalText implements encoding.TextMarshaler.
func (n Time) MarshalText() ([]byte, error) {
	if !n.IsValid() {
		return EmptyBytes, nil
	}

	return []byte(n.Format(WithTimeLayoutFormat())), nil
}

// Scan implements the sql.Scanner interface.
func (n *Time) Scan(src any) (err error) {
	switch v := src.(type) {
	case time.Time:
		n.value = v
	case string:
		n.value, err = n.parseTime(v)

		if err != nil {
			return NewScannerError(v, n, err)
		}
	case []byte:
		n.value, err = n.parseTime(string(v))

		if err != nil {
			return NewScannerError(v, n, err)
		}
	case int64:
		n.value = time.Unix(v, 0).UTC()
		n.valid = true

		return nil
	case nil:
		n.value = ZeroTime
		n.valid = false

		return nil
	default:
		return NewScannerError(v, n)
	}

	n.valid = true

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Time) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, NullStringBytes) {
		n.value = ZeroTime
		n.valid = false

		return nil
	}

	str, err := strconv.Unquote(string(data))

	if err != nil {
		return NewUnmarshalError(data, n, err)
	}

	n.value, err = n.parseTime(string(str))

	if err != nil {
		return NewUnmarshalError(data, n, err)
	}

	n.valid = true

	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *Time) UnmarshalText(text []byte) (err error) {
	if len(text) == 0 {
		n.value = ZeroTime
		n.valid = false

		return nil
	}

	n.value, err = n.parseTime(string(text))

	if err != nil {
		return NewUnmarshalError(text, n, err)
	}

	n.valid = true

	return nil
}

// Format returns a textual representation of the time value formatted according
// to the layout defined by the argument.
// If none option argument is specified then it defaults to RFC3339.
func (n Time) Format(options ...TimeFormatOption) string {
	if len(options) > 0 {
		option := new(timeFormatOption)
		options[0](option)

		if option.useTimeLayout {
			option.layout = n.layout
		}

		if option.layout == ZeroString {
			option.layout = time.RFC3339
		}

		return n.value.Format(option.layout)
	}

	return n.value.Format(time.RFC3339)
}

// parseTime parses a formatted string and returns the time value it represents.
// In strict mode if the date is ambigous mm/dd vs dd/mm it will return an error.
// Otherwise parses an unknown date format and detects the layout.
func (n Time) parseTime(value string) (format time.Time, err error) {
	if n.isStrictLayout {
		return dateparse.ParseStrict(value, n.parseOptions...)
	}

	return dateparse.ParseAny(value, n.parseOptions...)
}
