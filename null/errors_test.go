package null

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalError(t *testing.T) {
	errMock := gofakeit.ErrorValidation()

	sourceType := ZeroString
	err := NewMarshalError(sourceType, errMock)

	expectedErrMsg := fmt.Sprintf("%v %T; %v", ErrCannotMarshal, sourceType, errMock)
	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrMsg)

	var unwrappedErr MarshalError
	require.ErrorAs(t, err, &unwrappedErr)
	require.ErrorIs(t, unwrappedErr, ErrCannotMarshal)
}

func TestUnmarshalError(t *testing.T) {
	errMock := gofakeit.ErrorValidation()

	sourceType := ZeroString
	targetType := ZeroInt16
	err := NewUnmarshalError(sourceType, targetType, errMock)

	expectedErrMsg := fmt.Sprintf(
		"%v %T into %T; %v",
		ErrCannotUnmarshal,
		sourceType,
		targetType,
		errMock,
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrMsg)

	var unwrappedErr UnmarshalError
	require.ErrorAs(t, err, &unwrappedErr)
	require.ErrorIs(t, unwrappedErr, ErrCannotUnmarshal)
}

func TestScannerError(t *testing.T) {
	errMock := gofakeit.ErrorDatabase()

	sourceType := ZeroString
	targetType := ZeroInt16
	err := NewScannerError(sourceType, targetType, errMock)

	expectedErrMsg := fmt.Sprintf(
		"%v %T into %T; %v",
		ErrCannotScan,
		sourceType,
		targetType,
		errMock,
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrMsg)

	var unwrappedErr ScannerError
	require.ErrorAs(t, err, &unwrappedErr)
	require.ErrorIs(t, unwrappedErr, ErrCannotScan)
}

func TestValuerError(t *testing.T) {
	errMock := gofakeit.ErrorDatabase()

	sourceType := ZeroString
	err := NewValuerError(sourceType, errMock)

	expectedErrMsg := fmt.Sprintf("%v %T; %v", ErrCannotValue, sourceType, errMock)
	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedErrMsg)

	var unwrappedErr ValuerError
	require.ErrorAs(t, err, &unwrappedErr)
	require.ErrorIs(t, unwrappedErr, ErrCannotValue)
}
