package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"errors"
	"fmt"
)

var (
	ErrCannotMarshal = errors.New("null: cannot marshal type")

	ErrCannotUnmarshal     = errors.New("null: cannot unmarshal type")
	ErrCannotUnmarshalByte = errors.New(
		"null: cannot convert to byte, data length is greater than one",
	)

	ErrCannotScan = errors.New("null: cannot scan type")

	ErrCannotValue                  = errors.New("null: cannot value type")
	ErrValuerCheckerAssertionFailed = errors.New("null: valuer checker assertion failed")
	ErrValuerCheckerTypeUnsupported = errors.New("null: valuer checker type unsupported")
	ErrValuerCheckerIntegerOverflow = errors.New("null: valuer checker integer overflow detected")

	ErrCannotMustValue = errors.New("null: cannot must value for type")
	ErrDestinationNil  = errors.New("null: destination pointer is nil")

	ErrCannotNewUUID = errors.New("null: uuid value must be a string or implement fmt.Stringer")
)

// MarshalError represents an error that occurs during marshaling.
// It contains the original error and the source type that caused the failure.
type MarshalError struct {
	err        error
	sourceType any
}

// NewMarshalError creates a new MarshalError.
// sourceType indicates the type of the value being marshaled.
// It accepts multiple errors and wraps them together.
func NewMarshalError(sourceType any, errors ...error) error {
	var err = fmt.Errorf("%w %T", ErrCannotMarshal, sourceType)

	for _, item := range errors {
		err = fmt.Errorf("%w; %w", err, item)
	}

	return MarshalError{
		err:        err,
		sourceType: sourceType,
	}
}

// Error returns the string representation of the MarshalError.
func (e MarshalError) Error() string {
	return e.err.Error()
}

// Unwrap returns the underlying error for unwrapping.
func (e MarshalError) Unwrap() error {
	return e.err
}

// UnmarshalError represents an error that occurs during unmarshaling.
// It contains the original error, source type, and target type.
type UnmarshalError struct {
	err        error
	sourceType any
	targetType any
}

// NewUnmarshalError creates a new UnmarshalError.
// sourceType represents the type being unmarshaled from, and targetType is the type being unmarshaled into.
// It accepts multiple errors and wraps them together.
func NewUnmarshalError(sourceType any, targetType any, errors ...error) error {
	var err = fmt.Errorf("%w %T into %T", ErrCannotUnmarshal, sourceType, targetType)

	for _, item := range errors {
		err = fmt.Errorf("%w; %w", err, item)
	}

	return UnmarshalError{
		err:        err,
		sourceType: sourceType,
		targetType: targetType,
	}
}

// Error returns the string representation of the UnmarshalError.
func (e UnmarshalError) Error() string {
	return e.err.Error()
}

// Unwrap returns the underlying error for unwrapping.
func (e UnmarshalError) Unwrap() error {
	return e.err
}

// ScannerError represents an error that occurs during scanning.
// It contains the original error, source type, and target type.
type ScannerError struct {
	err        error
	sourceType any
	targetType any
}

// NewScannerError creates a new ScannerError.
// sourceType is the type being scanned from, and targetType is the type being scanned into.
// It accepts multiple errors and wraps them together.
func NewScannerError(sourceType any, targetType any, errors ...error) error {
	err := fmt.Errorf("%w %T into %T", ErrCannotScan, sourceType, targetType)

	for _, item := range errors {
		err = fmt.Errorf("%w; %w", err, item)
	}

	return ScannerError{
		err:        err,
		sourceType: sourceType,
		targetType: targetType,
	}
}

// Error returns the string representation of the ScannerError.
func (e ScannerError) Error() string {
	return e.err.Error()
}

// Unwrap returns the underlying error for unwrapping.
func (e ScannerError) Unwrap() error {
	return e.err
}

// ValuerError represents an error that occurs when attempting to retrieve a value.
// It contains the original error and the source type.
type ValuerError struct {
	err        error
	sourceType any
}

// NewValuerError creates a new ValuerError.
// sourceType represents the type that caused the error.
// It accepts multiple errors and wraps them together.
func NewValuerError(sourceType any, errors ...error) error {
	err := fmt.Errorf("%w %T", ErrCannotValue, sourceType)

	for _, item := range errors {
		err = fmt.Errorf("%w; %w", err, item)
	}

	return ValuerError{
		err:        err,
		sourceType: sourceType,
	}
}

// Error returns the string representation of the ValuerError.
func (e ValuerError) Error() string {
	return e.err.Error()
}

// Unwrap returns the underlying error for unwrapping.
func (e ValuerError) Unwrap() error {
	return e.err
}
