// Code generated by mockery. DO NOT EDIT.

package null

import (
	driver "database/sql/driver"

	mock "github.com/stretchr/testify/mock"
)

// MockNullable is an autogenerated mock type for the Nullable type
type MockNullable struct {
	mock.Mock
}

type MockNullable_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNullable) EXPECT() *MockNullable_Expecter {
	return &MockNullable_Expecter{mock: &_m.Mock}
}

// IsValid provides a mock function with given fields:
func (_m *MockNullable) IsValid() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsValid")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockNullable_IsValid_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsValid'
type MockNullable_IsValid_Call struct {
	*mock.Call
}

// IsValid is a helper method to define mock.On call
func (_e *MockNullable_Expecter) IsValid() *MockNullable_IsValid_Call {
	return &MockNullable_IsValid_Call{Call: _e.mock.On("IsValid")}
}

func (_c *MockNullable_IsValid_Call) Run(run func()) *MockNullable_IsValid_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNullable_IsValid_Call) Return(_a0 bool) *MockNullable_IsValid_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNullable_IsValid_Call) RunAndReturn(run func() bool) *MockNullable_IsValid_Call {
	_c.Call.Return(run)
	return _c
}

// IsZero provides a mock function with given fields:
func (_m *MockNullable) IsZero() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsZero")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockNullable_IsZero_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsZero'
type MockNullable_IsZero_Call struct {
	*mock.Call
}

// IsZero is a helper method to define mock.On call
func (_e *MockNullable_Expecter) IsZero() *MockNullable_IsZero_Call {
	return &MockNullable_IsZero_Call{Call: _e.mock.On("IsZero")}
}

func (_c *MockNullable_IsZero_Call) Run(run func()) *MockNullable_IsZero_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNullable_IsZero_Call) Return(_a0 bool) *MockNullable_IsZero_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNullable_IsZero_Call) RunAndReturn(run func() bool) *MockNullable_IsZero_Call {
	_c.Call.Return(run)
	return _c
}

// MarshalJSON provides a mock function with given fields:
func (_m *MockNullable) MarshalJSON() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalJSON")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNullable_MarshalJSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalJSON'
type MockNullable_MarshalJSON_Call struct {
	*mock.Call
}

// MarshalJSON is a helper method to define mock.On call
func (_e *MockNullable_Expecter) MarshalJSON() *MockNullable_MarshalJSON_Call {
	return &MockNullable_MarshalJSON_Call{Call: _e.mock.On("MarshalJSON")}
}

func (_c *MockNullable_MarshalJSON_Call) Run(run func()) *MockNullable_MarshalJSON_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNullable_MarshalJSON_Call) Return(_a0 []byte, _a1 error) *MockNullable_MarshalJSON_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNullable_MarshalJSON_Call) RunAndReturn(run func() ([]byte, error)) *MockNullable_MarshalJSON_Call {
	_c.Call.Return(run)
	return _c
}

// MarshalText provides a mock function with given fields:
func (_m *MockNullable) MarshalText() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalText")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNullable_MarshalText_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalText'
type MockNullable_MarshalText_Call struct {
	*mock.Call
}

// MarshalText is a helper method to define mock.On call
func (_e *MockNullable_Expecter) MarshalText() *MockNullable_MarshalText_Call {
	return &MockNullable_MarshalText_Call{Call: _e.mock.On("MarshalText")}
}

func (_c *MockNullable_MarshalText_Call) Run(run func()) *MockNullable_MarshalText_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNullable_MarshalText_Call) Return(_a0 []byte, _a1 error) *MockNullable_MarshalText_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNullable_MarshalText_Call) RunAndReturn(run func() ([]byte, error)) *MockNullable_MarshalText_Call {
	_c.Call.Return(run)
	return _c
}

// Scan provides a mock function with given fields: src
func (_m *MockNullable) Scan(src interface{}) error {
	ret := _m.Called(src)

	if len(ret) == 0 {
		panic("no return value specified for Scan")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(src)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNullable_Scan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Scan'
type MockNullable_Scan_Call struct {
	*mock.Call
}

// Scan is a helper method to define mock.On call
//   - src interface{}
func (_e *MockNullable_Expecter) Scan(src interface{}) *MockNullable_Scan_Call {
	return &MockNullable_Scan_Call{Call: _e.mock.On("Scan", src)}
}

func (_c *MockNullable_Scan_Call) Run(run func(src interface{})) *MockNullable_Scan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockNullable_Scan_Call) Return(err error) *MockNullable_Scan_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockNullable_Scan_Call) RunAndReturn(run func(interface{}) error) *MockNullable_Scan_Call {
	_c.Call.Return(run)
	return _c
}

// UnmarshalJSON provides a mock function with given fields: data
func (_m *MockNullable) UnmarshalJSON(data []byte) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalJSON")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNullable_UnmarshalJSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalJSON'
type MockNullable_UnmarshalJSON_Call struct {
	*mock.Call
}

// UnmarshalJSON is a helper method to define mock.On call
//   - data []byte
func (_e *MockNullable_Expecter) UnmarshalJSON(data interface{}) *MockNullable_UnmarshalJSON_Call {
	return &MockNullable_UnmarshalJSON_Call{Call: _e.mock.On("UnmarshalJSON", data)}
}

func (_c *MockNullable_UnmarshalJSON_Call) Run(run func(data []byte)) *MockNullable_UnmarshalJSON_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockNullable_UnmarshalJSON_Call) Return(err error) *MockNullable_UnmarshalJSON_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockNullable_UnmarshalJSON_Call) RunAndReturn(run func([]byte) error) *MockNullable_UnmarshalJSON_Call {
	_c.Call.Return(run)
	return _c
}

// UnmarshalText provides a mock function with given fields: data
func (_m *MockNullable) UnmarshalText(data []byte) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalText")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNullable_UnmarshalText_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalText'
type MockNullable_UnmarshalText_Call struct {
	*mock.Call
}

// UnmarshalText is a helper method to define mock.On call
//   - data []byte
func (_e *MockNullable_Expecter) UnmarshalText(data interface{}) *MockNullable_UnmarshalText_Call {
	return &MockNullable_UnmarshalText_Call{Call: _e.mock.On("UnmarshalText", data)}
}

func (_c *MockNullable_UnmarshalText_Call) Run(run func(data []byte)) *MockNullable_UnmarshalText_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockNullable_UnmarshalText_Call) Return(err error) *MockNullable_UnmarshalText_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *MockNullable_UnmarshalText_Call) RunAndReturn(run func([]byte) error) *MockNullable_UnmarshalText_Call {
	_c.Call.Return(run)
	return _c
}

// Value provides a mock function with given fields:
func (_m *MockNullable) Value() (driver.Value, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Value")
	}

	var r0 driver.Value
	var r1 error
	if rf, ok := ret.Get(0).(func() (driver.Value, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() driver.Value); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Value)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNullable_Value_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Value'
type MockNullable_Value_Call struct {
	*mock.Call
}

// Value is a helper method to define mock.On call
func (_e *MockNullable_Expecter) Value() *MockNullable_Value_Call {
	return &MockNullable_Value_Call{Call: _e.mock.On("Value")}
}

func (_c *MockNullable_Value_Call) Run(run func()) *MockNullable_Value_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockNullable_Value_Call) Return(_a0 driver.Value, _a1 error) *MockNullable_Value_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNullable_Value_Call) RunAndReturn(run func() (driver.Value, error)) *MockNullable_Value_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNullable creates a new instance of MockNullable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNullable(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNullable {
	mock := &MockNullable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
