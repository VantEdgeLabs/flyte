// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CacheItem is an autogenerated mock type for the CacheItem type
type CacheItem struct {
	mock.Mock
}

type CacheItem_Expecter struct {
	mock *mock.Mock
}

func (_m *CacheItem) EXPECT() *CacheItem_Expecter {
	return &CacheItem_Expecter{mock: &_m.Mock}
}

// ID provides a mock function with given fields:
func (_m *CacheItem) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CacheItem_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type CacheItem_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *CacheItem_Expecter) ID() *CacheItem_ID_Call {
	return &CacheItem_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *CacheItem_ID_Call) Run(run func()) *CacheItem_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CacheItem_ID_Call) Return(_a0 string) *CacheItem_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheItem_ID_Call) RunAndReturn(run func() string) *CacheItem_ID_Call {
	_c.Call.Return(run)
	return _c
}

// NewCacheItem creates a new instance of CacheItem. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCacheItem(t interface {
	mock.TestingT
	Cleanup(func())
}) *CacheItem {
	mock := &CacheItem{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
