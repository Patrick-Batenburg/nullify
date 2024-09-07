package null

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type NullableValue interface {
	Value() (driver.Value, error)
}

type IntegerValuerTestCase[T any, U any] struct {
	actualAndExpectedIntegerFn integerResultFn
	sutFn                      newNullableScanSUTFn
}

type newNullableScanSUTFn func() (NullableValue, integerOption)

func IntegerOptionFrom(IntegerOptionFn IntegerOption) integerOption {
	option := new(integerOption)
	IntegerOptionFn(option)

	return *option
}

func newIntTestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewInt(ZeroInt, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newInt8TestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewInt8(ZeroInt8, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newInt16TestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewInt16(ZeroInt16, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newInt32TestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewInt32(ZeroInt32, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newInt64TestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewInt64(ZeroInt64, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newUintTestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewUint(ZeroUint, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newUint8TestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewUint8(ZeroUint8, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newUint16TestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewUint16(ZeroUint16, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newUint32TestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewUint32(ZeroUint32, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func newUint64TestValuerFn(IntegerOptionFn IntegerOption) newNullableScanSUTFn {
	return func() (NullableValue, integerOption) {
		return NewUint64(ZeroUint64, false, IntegerOptionFn), IntegerOptionFrom(IntegerOptionFn)
	}
}

func TestIntegerValuer(t *testing.T) {
	testCases := map[string]IntegerValuerTestCase[int, int]{
		"TestNewInt WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroInt),
			sutFn:                      newIntTestValuerFn(WithIntValuer()),
		},
		"TestNewInt WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroInt8),
			sutFn:                      newIntTestValuerFn(WithInt8Valuer()),
		},
		"TestNewInt WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroInt16),
			sutFn:                      newIntTestValuerFn(WithInt16Valuer()),
		},
		"TestNewInt WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroInt32),
			sutFn:                      newIntTestValuerFn(WithInt32Valuer()),
		},
		"TestNewInt WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroInt64),
			sutFn:                      newIntTestValuerFn(WithInt64Valuer()),
		},
		"TestNewInt WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroUint),
			sutFn:                      newIntTestValuerFn(WithUintValuer()),
		},
		"TestNewInt WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroUint8),
			sutFn:                      newIntTestValuerFn(WithUint8Valuer()),
		},
		"TestNewInt WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroUint16),
			sutFn:                      newIntTestValuerFn(WithUint16Valuer()),
		},
		"TestNewInt WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroUint32),
			sutFn:                      newIntTestValuerFn(WithUint32Valuer()),
		},
		"TestNewInt WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt, ZeroUint64),
			sutFn:                      newIntTestValuerFn(WithUint64Valuer()),
		},

		"TestNewInt8 WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroInt),
			sutFn:                      newInt8TestValuerFn(WithIntValuer()),
		},
		"TestNewInt8 WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroInt8),
			sutFn:                      newInt8TestValuerFn(WithInt8Valuer()),
		},
		"TestNewInt8 WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroInt16),
			sutFn:                      newInt8TestValuerFn(WithInt16Valuer()),
		},
		"TestNewInt8 WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroInt32),
			sutFn:                      newInt8TestValuerFn(WithInt32Valuer()),
		},
		"TestNewInt8 WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroInt64),
			sutFn:                      newInt8TestValuerFn(WithInt64Valuer()),
		},
		"TestNewInt8 WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroUint),
			sutFn:                      newInt8TestValuerFn(WithUintValuer()),
		},
		"TestNewInt8 WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroUint8),
			sutFn:                      newInt8TestValuerFn(WithUint8Valuer()),
		},
		"TestNewInt8 WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroUint16),
			sutFn:                      newInt8TestValuerFn(WithUint16Valuer()),
		},
		"TestNewInt8 WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroUint32),
			sutFn:                      newInt8TestValuerFn(WithUint32Valuer()),
		},
		"TestNewInt8 WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt8, ZeroUint64),
			sutFn:                      newInt8TestValuerFn(WithUint64Valuer()),
		},
		"TestNewInt16 WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroInt),
			sutFn:                      newInt16TestValuerFn(WithIntValuer()),
		},
		"TestNewInt16 WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroInt8),
			sutFn:                      newInt16TestValuerFn(WithInt16Valuer()),
		},
		"TestNewInt16 WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroInt16),
			sutFn:                      newInt16TestValuerFn(WithInt16Valuer()),
		},
		"TestNewInt16 WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroInt32),
			sutFn:                      newInt16TestValuerFn(WithInt32Valuer()),
		},
		"TestNewInt16 WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroInt64),
			sutFn:                      newInt16TestValuerFn(WithInt64Valuer()),
		},
		"TestNewInt16 WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroUint),
			sutFn:                      newInt16TestValuerFn(WithUintValuer()),
		},
		"TestNewInt16 WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroUint8),
			sutFn:                      newInt16TestValuerFn(WithUint8Valuer()),
		},
		"TestNewInt16 WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroUint16),
			sutFn:                      newInt16TestValuerFn(WithUint16Valuer()),
		},
		"TestNewInt16 WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroUint32),
			sutFn:                      newInt16TestValuerFn(WithUint32Valuer()),
		},
		"TestNewInt16 WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt16, ZeroUint64),
			sutFn:                      newInt16TestValuerFn(WithUint64Valuer()),
		},

		"TestNewInt32 WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroInt),
			sutFn:                      newInt32TestValuerFn(WithIntValuer()),
		},
		"TestNewInt32 WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroInt8),
			sutFn:                      newInt32TestValuerFn(WithInt32Valuer()),
		},
		"TestNewInt32 WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroInt16),
			sutFn:                      newInt32TestValuerFn(WithInt16Valuer()),
		},
		"TestNewInt32 WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroInt32),
			sutFn:                      newInt32TestValuerFn(WithInt32Valuer()),
		},
		"TestNewInt32 WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroInt64),
			sutFn:                      newInt32TestValuerFn(WithInt64Valuer()),
		},
		"TestNewInt32 WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroUint),
			sutFn:                      newInt32TestValuerFn(WithUintValuer()),
		},
		"TestNewInt32 WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroUint8),
			sutFn:                      newInt32TestValuerFn(WithUint8Valuer()),
		},
		"TestNewInt32 WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroUint16),
			sutFn:                      newInt32TestValuerFn(WithUint16Valuer()),
		},
		"TestNewInt32 WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroUint32),
			sutFn:                      newInt32TestValuerFn(WithUint32Valuer()),
		},
		"TestNewInt32 WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt32, ZeroUint64),
			sutFn:                      newInt32TestValuerFn(WithUint64Valuer()),
		},

		"TestNewInt64 WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroInt),
			sutFn:                      newInt64TestValuerFn(WithIntValuer()),
		},
		"TestNewInt64 WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroInt8),
			sutFn:                      newInt64TestValuerFn(WithInt8Valuer()),
		},
		"TestNewInt64 WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroInt16),
			sutFn:                      newInt64TestValuerFn(WithInt16Valuer()),
		},
		"TestNewInt64 WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroInt32),
			sutFn:                      newInt64TestValuerFn(WithInt32Valuer()),
		},
		"TestNewInt64 WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroInt64),
			sutFn:                      newInt64TestValuerFn(WithInt64Valuer()),
		},
		"TestNewInt64 WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroUint),
			sutFn:                      newInt64TestValuerFn(WithUintValuer()),
		},
		"TestNewInt64 WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroUint8),
			sutFn:                      newInt64TestValuerFn(WithUint8Valuer()),
		},
		"TestNewInt64 WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroUint16),
			sutFn:                      newInt64TestValuerFn(WithUint16Valuer()),
		},
		"TestNewInt64 WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroUint32),
			sutFn:                      newInt64TestValuerFn(WithUint32Valuer()),
		},
		"TestNewInt64 WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroInt64, ZeroUint64),
			sutFn:                      newInt64TestValuerFn(WithUint64Valuer()),
		},

		"TestNewUint WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroInt),
			sutFn:                      newUintTestValuerFn(WithIntValuer()),
		},
		"TestNewUint WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroInt8),
			sutFn:                      newUintTestValuerFn(WithInt8Valuer()),
		},
		"TestNewUint WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroInt16),
			sutFn:                      newUintTestValuerFn(WithInt16Valuer()),
		},
		"TestNewUint WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroInt32),
			sutFn:                      newUintTestValuerFn(WithInt32Valuer()),
		},
		"TestNewUint WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroInt64),
			sutFn:                      newUintTestValuerFn(WithInt64Valuer()),
		},
		"TestNewUint WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroUint),
			sutFn:                      newUintTestValuerFn(WithUintValuer()),
		},
		"TestNewUint WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroUint8),
			sutFn:                      newUintTestValuerFn(WithUint8Valuer()),
		},
		"TestNewUint WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroUint16),
			sutFn:                      newUintTestValuerFn(WithUint16Valuer()),
		},
		"TestNewUint WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroUint32),
			sutFn:                      newUintTestValuerFn(WithUint32Valuer()),
		},
		"TestNewUint WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint, ZeroUint64),
			sutFn:                      newUintTestValuerFn(WithUint64Valuer()),
		},

		"TestNewUint8 WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroInt),
			sutFn:                      newUint8TestValuerFn(WithIntValuer()),
		},
		"TestNewUint8 WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroInt8),
			sutFn:                      newUint8TestValuerFn(WithInt8Valuer()),
		},
		"TestNewUint8 WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroInt16),
			sutFn:                      newUint8TestValuerFn(WithInt16Valuer()),
		},
		"TestNewUint8 WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroInt32),
			sutFn:                      newUint8TestValuerFn(WithInt32Valuer()),
		},
		"TestNewUint8 WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroInt64),
			sutFn:                      newUint8TestValuerFn(WithInt64Valuer()),
		},
		"TestNewUint8 WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroUint),
			sutFn:                      newUint8TestValuerFn(WithUintValuer()),
		},
		"TestNewUint8 WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroUint8),
			sutFn:                      newUint8TestValuerFn(WithUint8Valuer()),
		},
		"TestNewUint8 WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroUint16),
			sutFn:                      newUint8TestValuerFn(WithUint16Valuer()),
		},
		"TestNewUint8 WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroUint32),
			sutFn:                      newUint8TestValuerFn(WithUint32Valuer()),
		},
		"TestNewUint8 WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint8, ZeroUint64),
			sutFn:                      newUint8TestValuerFn(WithUint64Valuer()),
		},

		"TestNewUint16 WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroInt),
			sutFn:                      newUint16TestValuerFn(WithIntValuer()),
		},
		"TestNewUint16 WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroInt8),
			sutFn:                      newUint16TestValuerFn(WithInt8Valuer()),
		},
		"TestNewUint16 WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroInt16),
			sutFn:                      newUint16TestValuerFn(WithInt16Valuer()),
		},
		"TestNewUint16 WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroInt32),
			sutFn:                      newUint16TestValuerFn(WithInt32Valuer()),
		},
		"TestNewUint16 WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroInt64),
			sutFn:                      newUint16TestValuerFn(WithInt64Valuer()),
		},
		"TestNewUint16 WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroUint),
			sutFn:                      newUint16TestValuerFn(WithUintValuer()),
		},
		"TestNewUint16 WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroUint8),
			sutFn:                      newUint16TestValuerFn(WithUint8Valuer()),
		},
		"TestNewUint16 WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroUint16),
			sutFn:                      newUint16TestValuerFn(WithUint16Valuer()),
		},
		"TestNewUint16 WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroUint32),
			sutFn:                      newUint16TestValuerFn(WithUint32Valuer()),
		},
		"TestNewUint16 WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint16, ZeroUint64),
			sutFn:                      newUint16TestValuerFn(WithUint64Valuer()),
		},

		"TestNewUint32 WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroInt),
			sutFn:                      newUint32TestValuerFn(WithIntValuer()),
		},
		"TestNewUint32 WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroInt8),
			sutFn:                      newUint32TestValuerFn(WithInt8Valuer()),
		},
		"TestNewUint32 WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroInt16),
			sutFn:                      newUint32TestValuerFn(WithInt16Valuer()),
		},
		"TestNewUint32 WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroInt32),
			sutFn:                      newUint32TestValuerFn(WithInt32Valuer()),
		},
		"TestNewUint32 WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroInt64),
			sutFn:                      newUint32TestValuerFn(WithInt64Valuer()),
		},
		"TestNewUint32 WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroUint),
			sutFn:                      newUint32TestValuerFn(WithUintValuer()),
		},
		"TestNewUint32 WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroUint8),
			sutFn:                      newUint32TestValuerFn(WithUint8Valuer()),
		},
		"TestNewUint32 WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroUint16),
			sutFn:                      newUint32TestValuerFn(WithUint16Valuer()),
		},
		"TestNewUint32 WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroUint32),
			sutFn:                      newUint32TestValuerFn(WithUint32Valuer()),
		},
		"TestNewUint32 WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint32, ZeroUint64),
			sutFn:                      newUint32TestValuerFn(WithUint64Valuer()),
		},

		"TestNewUint64 WithIntValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroInt),
			sutFn:                      newUint64TestValuerFn(WithIntValuer()),
		},
		"TestNewUint64 WithInt8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroInt8),
			sutFn:                      newUint64TestValuerFn(WithInt8Valuer()),
		},
		"TestNewUint64 WithInt16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroInt16),
			sutFn:                      newUint64TestValuerFn(WithInt16Valuer()),
		},
		"TestNewUint64 WithInt32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroInt32),
			sutFn:                      newUint64TestValuerFn(WithInt32Valuer()),
		},
		"TestNewUint64 WithInt64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroInt64),
			sutFn:                      newUint64TestValuerFn(WithInt64Valuer()),
		},
		"TestNewUint64 WithUintValuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroUint),
			sutFn:                      newUint64TestValuerFn(WithUintValuer()),
		},
		"TestNewUint64 WithUint8Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroUint8),
			sutFn:                      newUint64TestValuerFn(WithUint8Valuer()),
		},
		"TestNewUint64 WithUint16Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroUint16),
			sutFn:                      newUint64TestValuerFn(WithUint16Valuer()),
		},
		"TestNewUint64 WithUint32Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroUint32),
			sutFn:                      newUint64TestValuerFn(WithUint32Valuer()),
		},
		"TestNewUint64 WithUint64Valuer": {
			actualAndExpectedIntegerFn: withActualAndExpectedInteger(t, ZeroUint64, ZeroUint64),
			sutFn:                      newUint64TestValuerFn(WithUint64Valuer()),
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, func(t *testing.T) {
			sut, option := testCase.sutFn()
			integerResult := new(integerResult)
			testCase.actualAndExpectedIntegerFn(integerResult)

			switch s := sut.(type) {
			case Int:
				s.SetValue(integerResult.actual.(int))
				assert.Equal(t, Int{
					NullableImpl: NullableImpl[int]{
						value: integerResult.actual.(int),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Int8:
				s.SetValue(integerResult.actual.(int8))
				assert.Equal(t, Int8{
					NullableImpl: NullableImpl[int8]{
						value: integerResult.actual.(int8),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Int16:
				s.SetValue(integerResult.actual.(int16))
				assert.Equal(t, Int16{
					NullableImpl: NullableImpl[int16]{
						value: integerResult.actual.(int16),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Int32:
				s.SetValue(integerResult.actual.(int32))
				assert.Equal(t, Int32{
					NullableImpl: NullableImpl[int32]{
						value: integerResult.actual.(int32),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Int64:
				s.SetValue(integerResult.actual.(int64))
				assert.Equal(t, Int64{
					NullableImpl: NullableImpl[int64]{
						value: integerResult.actual.(int64),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Uint:
				s.SetValue(integerResult.actual.(uint))
				assert.Equal(t, Uint{
					NullableImpl: NullableImpl[uint]{
						value: integerResult.actual.(uint),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Uint8:
				s.SetValue(integerResult.actual.(uint8))
				assert.Equal(t, Uint8{
					NullableImpl: NullableImpl[uint8]{
						value: integerResult.actual.(uint8),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Uint16:
				s.SetValue(integerResult.actual.(uint16))
				assert.Equal(t, Uint16{
					NullableImpl: NullableImpl[uint16]{
						value: integerResult.actual.(uint16),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Uint32:
				s.SetValue(integerResult.actual.(uint32))
				assert.Equal(t, Uint32{
					NullableImpl: NullableImpl[uint32]{
						value: integerResult.actual.(uint32),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			case Uint64:
				s.SetValue(integerResult.actual.(uint64))
				assert.Equal(t, Uint64{
					NullableImpl: NullableImpl[uint64]{
						value: integerResult.actual.(uint64),
						valid: true,
					},
					valuerType: option.valuerType,
				}, s)
				sut = s
			default:
				t.Errorf("invalid type")
			}

			valuer, err := sut.Value()
			require.NoError(t, err)
			assert.EqualValues(t, integerResult.expected, valuer)
		})
	}
}
