package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"time"

	"github.com/itlightning/dateparse"
)

// TimeOptionFn is a type alias for a function that modifies an TimeOption.
type TimeOptionFn = func(*Time)

// WithTimeLayout sets a time layout.
func WithTimeParseOptions(options ...dateparse.ParserOption) TimeOptionFn {
	return func(option *Time) {
		option.parseOptions = options
	}
}

// WithTimeLayout sets a time layout.
func WithTimeLayout(layout string) TimeOptionFn {
	return func(option *Time) {
		option.layout = layout
	}
}

// WithTimeDefaultLayout sets the time layout to the RFC3339 format.
func WithTimeDefaultLayout() TimeOptionFn {
	return func(option *Time) {
		option.layout = time.RFC3339
	}
}

// WithTimeLenientParsing enables lenient time paring when
// json.Unmarshaler and encoding.TextUnmarshaler are called.
func WithTimeLenientParsing() TimeOptionFn {
	return func(option *Time) {
		option.isStrictLayout = false
	}
}

// WithTimeStrictParsing enables strict time paring when
// json.Unmarshaler and encoding.TextUnmarshaler are called.
func WithTimeStrictParsing() TimeOptionFn {
	return func(option *Time) {
		option.isStrictLayout = true
	}
}

// TimeFormatOption struct holds format option for formatting a time value.
type timeFormatOption struct {
	// layout defines the format pattern for the time value (e.g. "2006-01-02").
	layout string

	// useTimeLayout indicates whether to use the layout associated with the Time instance.
	// If true, the Time instance's layout will be used instead of any other provided layout.
	useTimeLayout bool
}

// TimeFormatOptionFn is a function type that modifies a formatOption instance.
// It is used to apply different formatting options when formatting a Time value.
type TimeFormatOption = func(*timeFormatOption)

// WithLayoutValue returns a FormatOptionFn that sets a custom layout for formatting the time value.
// This is useful when you want to specify a particular format.
func WithTimeLayoutValue(layout string) TimeFormatOption {
	return func(option *timeFormatOption) {
		option.layout = layout
	}
}

// WithTimeLayout returns a FormatOptionFn that sets the useTimeLayout flag to true.
// When this option is used, the formatter will use the layout associated with the Time instance.
func WithTimeLayoutFormat() TimeFormatOption {
	return func(option *timeFormatOption) {
		option.useTimeLayout = true
	}
}

// WithDefaultLayout returns a FormatOptionFn that sets the layout to the RFC3339 format.
func WithTimeDefaultLayoutFormat() TimeFormatOption {
	return func(option *timeFormatOption) {
		option.layout = time.RFC3339
	}
}
