package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"time"

	"github.com/google/uuid"
)

var (
	// ZeroString is the zero value for the string type.
	ZeroString = ""

	// NullString is a string representation of null.
	NullString = "null"

	// ZeroIntegerString is a string representation of 0.
	ZeroIntegerString = "0"

	// TrueString is a string representation of true.
	TrueString = "true"

	// FalseString is a string representation of false.
	FalseString = "false"

	// ZeroUUIDString is a string representation of an empty uuid with all zeros.
	ZeroUUIDString = uuid.Nil.String()

	// ZeroInt is the zero value for the int type, which is at least 32 bits in size.
	ZeroInt int = 0

	// ZeroInt8 is the zero value for the int8 type.
	ZeroInt8 int8 = 0

	// ZeroInt16 is the zero value for the int16 type.
	ZeroInt16 int16 = 0

	// ZeroInt32 is the zero value for the int32 type.
	ZeroInt32 int32 = 0

	// ZeroInt64 is the zero value for the int64 type.
	ZeroInt64 int64 = 0

	// ZeroUint is the zero value for the uint type, which is at least 32 bits in size.
	ZeroUint uint = 0

	// ZeroUint8 is the zero value for the uint8 type.
	ZeroUint8 uint8 = 0

	// ZeroUint16 is the zero value for the uint16 type.
	ZeroUint16 uint16 = 0

	// ZeroUint32 is the zero value for the uint32 type
	ZeroUint32 uint32 = 0

	// ZeroUint64 is the zero value for the uint64 type.
	ZeroUint64 uint64 = 0

	// ZeroFloat32 is the zero value for the float32 type.
	ZeroFloat32 float32 = 0

	// ZeroFloat64 is the zero value for the float64 type.
	ZeroFloat64 float64 = 0

	// ZeroByte is the zero value for the byte type.
	ZeroByte byte = 0

	// ZeroBytes is the zero value for the []byte type, which is a nil byte slice.
	ZeroBytes = []byte(nil)

	// ZeroBool is the zero value for the bool type.
	ZeroBool = false

	// ZeroTime is the zero value for the time.Time type.
	ZeroTime = time.Time{}

	// EmptyBytes is an empty byte slice.
	EmptyBytes = []byte{}

	// NullStringBytes is a byte slice of a "null" string value.
	NullStringBytes = []byte(NullString)

	// ZeroIntegerStringBytes is a byte slice of a "0" string value.
	ZeroIntegerStringBytes = []byte(ZeroIntegerString)

	// TrueStringBytes is a byte slice of a "true" string value.
	TrueStringBytes = []byte(TrueString)

	// FalseStringBytes is a byte slice of a "false" string value.
	FalseStringBytes = []byte(FalseString)

	// ZeroStringBytes is a byte slice of an empty string value.
	ZeroStringBytes = []byte(ZeroString)
)
