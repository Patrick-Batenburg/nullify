package null // import "github.com/Patrick-Batenburg/nullify/null"

import (
	"database/sql/driver"
	"math"
)

// integerValuerChecker attempts to cast value to the type specified by valuerType.
// Returns the cast value or an error if the type assertion fails or if an overflow occurs.
func integerValuerChecker(value any, valuerType any) (driver.Value, error) {
	switch valuerType.(type) {
	case int:
		return castToIntValuer(value, valuerType)
	case int8:
		return castToIntValuer(value, valuerType)
	case int16:
		return castToIntValuer(value, valuerType)
	case int32:
		return castToIntValuer(value, valuerType)
	case int64:
		return castToIntValuer(value, valuerType)
	case uint:
		return castToUintValuer(value, valuerType)
	case uint8:
		return castToUintValuer(value, valuerType)
	case uint16:
		return castToUintValuer(value, valuerType)
	case uint32:
		return castToUintValuer(value, valuerType)
	case uint64:
		return castToUintValuer(value, valuerType)
	default:
		return value, nil
	}
}

// castToIntValuer attempts to cast value to a specified signed integer type.
func castToIntValuer(value any, targetType any) (driver.Value, error) {
	switch v := value.(type) {
	case int:
		return convertIntValuer(v, targetType)
	case int8:
		return convertIntValuer(int(v), targetType)
	case int16:
		return convertIntValuer(int(v), targetType)
	case int32:
		return convertIntValuer(int(v), targetType)
	case int64:
		return convertIntValuer(int(v), targetType)
	case uint:
		return convertIntValuer(int(v), targetType)
	case uint8:
		return convertIntValuer(int(v), targetType)
	case uint16:
		return convertIntValuer(int(v), targetType)
	case uint32:
		return convertIntValuer(int(v), targetType)
	case uint64:
		return convertIntValuer(int(v), targetType)
	default:
		return nil, ErrValuerCheckerTypeUnsupported
	}
}

// castToUintValuer attempts to cast value to a specified unsigned integer type.
func castToUintValuer(value any, targetType any) (driver.Value, error) {
	switch v := value.(type) {
	case int:
		return convertUintValuer(uint(v), targetType)
	case int8:
		return convertUintValuer(uint(v), targetType)
	case int16:
		return convertUintValuer(uint(v), targetType)
	case int32:
		return convertUintValuer(uint(v), targetType)
	case int64:
		return convertUintValuer(uint(v), targetType)
	case uint:
		return convertUintValuer(v, targetType)
	case uint8:
		return convertUintValuer(uint(v), targetType)
	case uint16:
		return convertUintValuer(uint(v), targetType)
	case uint32:
		return convertUintValuer(uint(v), targetType)
	case uint64:
		return convertUintValuer(uint(v), targetType)
	default:
		return nil, ErrValuerCheckerTypeUnsupported
	}
}

// convertIntValuer converts an integer value to the specified target type.
func convertIntValuer(value int, targetType any) (driver.Value, error) {
	switch targetType.(type) {
	case int:
		return int(value), nil
	case int8:
		if value < math.MinInt8 || value > math.MaxInt8 {
			return nil, ErrValuerCheckerIntegerOverflow
		}

		return int8(value), nil
	case int16:
		if value < math.MinInt16 || value > math.MaxInt16 {
			return nil, ErrValuerCheckerIntegerOverflow
		}

		return int16(value), nil
	case int32:
		if value < math.MinInt32 || value > math.MaxInt32 {
			return nil, ErrValuerCheckerIntegerOverflow
		}

		return int32(value), nil
	case int64:
		return int64(value), nil
	default:
		return nil, ErrValuerCheckerTypeUnsupported
	}
}

// convertUintValuer converts an unsigned integer value to the specified target type.
func convertUintValuer(value uint, targetType any) (driver.Value, error) {
	switch targetType.(type) {
	case uint:
		return uint(value), nil
	case uint8:
		if value > math.MaxUint8 {
			return nil, ErrValuerCheckerIntegerOverflow
		}

		return uint8(value), nil
	case uint16:
		if value > math.MaxUint16 {
			return nil, ErrValuerCheckerIntegerOverflow
		}

		return uint16(value), nil
	case uint32:
		if value > math.MaxUint32 {
			return nil, ErrValuerCheckerIntegerOverflow
		}

		return uint32(value), nil
	case uint64:
		return uint64(value), nil
	default:
		return nil, ErrValuerCheckerTypeUnsupported
	}
}
