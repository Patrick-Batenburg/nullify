package null // import "github.com/Patrick-Batenburg/nullify/null"

// integerOption holds options for configuring integer value handling.
type integerOption struct {
	// valuerType sets the destination type when driver.Valuer is called.
	valuerType any
}

// IntegerOption is a type alias for a function that modifies an IntegerOption.
type IntegerOption = func(*integerOption)

// WithIntegerValuer sets a custom integer type for value conversion.
// This function allows specifying any integer type as the valuerType.
func WithIntegerValuer(value any) IntegerOption {
	return func(option *integerOption) {
		option.valuerType = value
	}
}

// WithIntValuer sets the valuerType to int.
// This function is used when you want to handle integer values as type int.
func WithIntValuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroInt
	}
}

// WithInt8Valuer sets the valuerType to int8.
// This function is used when you want to handle integer values as type int8.
func WithInt8Valuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroInt8
	}
}

// WithInt16Valuer sets the valuerType to int16.
// This function is used when you want to handle integer values as type int16.
func WithInt16Valuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroInt16
	}
}

// WithInt32Valuer sets the valuerType to int32.
// This function is used when you want to handle integer values as type int32.
func WithInt32Valuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroInt32
	}
}

// WithInt64Valuer sets the valuerType to int64.
// This function is used when you want to handle integer values as type int64.
func WithInt64Valuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroInt64
	}
}

// WithUintValuer sets the valuerType to uint.
// This function is used when you want to handle integer values as type uint.
func WithUintValuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroUint
	}
}

// WithUint8Valuer sets the valuerType to uint8.
// This function is used when you want to handle integer values as type uint8.
func WithUint8Valuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroUint8
	}
}

// WithUint16Valuer sets the valuerType to uint16.
// This function is used when you want to handle integer values as type uint16.
func WithUint16Valuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroUint16
	}
}

// WithUint32Valuer sets the valuerType to uint32.
// This function is used when you want to handle integer values as type uint32.
func WithUint32Valuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroUint32
	}
}

// WithUint64Valuer sets the valuerType to uint64.
// This function is used when you want to handle integer values as type uint64.
func WithUint64Valuer() IntegerOption {
	return func(option *integerOption) {
		option.valuerType = ZeroUint64
	}
}
