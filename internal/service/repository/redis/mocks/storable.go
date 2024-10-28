// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	redis "github.com/go-redis/redis/v8"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockStorable is an autogenerated mock type for the Storable type
type MockStorable struct {
	mock.Mock
}

type MockStorable_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorable) EXPECT() *MockStorable_Expecter {
	return &MockStorable_Expecter{mock: &_m.Mock}
}

// Del provides a mock function with given fields: ctx, keys
func (_m *MockStorable) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Del")
	}

	var r0 *redis.IntCmd
	if rf, ok := ret.Get(0).(func(context.Context, ...string) *redis.IntCmd); ok {
		r0 = rf(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.IntCmd)
		}
	}

	return r0
}

// MockStorable_Del_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Del'
type MockStorable_Del_Call struct {
	*mock.Call
}

// Del is a helper method to define mock.On call
//   - ctx context.Context
//   - keys ...string
func (_e *MockStorable_Expecter) Del(ctx interface{}, keys ...interface{}) *MockStorable_Del_Call {
	return &MockStorable_Del_Call{Call: _e.mock.On("Del",
		append([]interface{}{ctx}, keys...)...)}
}

func (_c *MockStorable_Del_Call) Run(run func(ctx context.Context, keys ...string)) *MockStorable_Del_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockStorable_Del_Call) Return(_a0 *redis.IntCmd) *MockStorable_Del_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorable_Del_Call) RunAndReturn(run func(context.Context, ...string) *redis.IntCmd) *MockStorable_Del_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key
func (_m *MockStorable) Get(ctx context.Context, key string) *redis.StringCmd {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *redis.StringCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *redis.StringCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StringCmd)
		}
	}

	return r0
}

// MockStorable_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockStorable_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockStorable_Expecter) Get(ctx interface{}, key interface{}) *MockStorable_Get_Call {
	return &MockStorable_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *MockStorable_Get_Call) Run(run func(ctx context.Context, key string)) *MockStorable_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStorable_Get_Call) Return(_a0 *redis.StringCmd) *MockStorable_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorable_Get_Call) RunAndReturn(run func(context.Context, string) *redis.StringCmd) *MockStorable_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: ctx, key, value, expiration
func (_m *MockStorable) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	ret := _m.Called(ctx, key, value, expiration)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 *redis.StatusCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StatusCmd)
		}
	}

	return r0
}

// MockStorable_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockStorable_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
//   - value interface{}
//   - expiration time.Duration
func (_e *MockStorable_Expecter) Set(ctx interface{}, key interface{}, value interface{}, expiration interface{}) *MockStorable_Set_Call {
	return &MockStorable_Set_Call{Call: _e.mock.On("Set", ctx, key, value, expiration)}
}

func (_c *MockStorable_Set_Call) Run(run func(ctx context.Context, key string, value interface{}, expiration time.Duration)) *MockStorable_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *MockStorable_Set_Call) Return(_a0 *redis.StatusCmd) *MockStorable_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorable_Set_Call) RunAndReturn(run func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd) *MockStorable_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorable creates a new instance of MockStorable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorable(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorable {
	mock := &MockStorable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}