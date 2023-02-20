// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	httpserver "github.com/bowoBp/csmart/utils/httpserver"
	mock "github.com/stretchr/testify/mock"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

type Router_Expecter struct {
	mock *mock.Mock
}

func (_m *Router) EXPECT() *Router_Expecter {
	return &Router_Expecter{mock: &_m.Mock}
}

// Route provides a mock function with given fields: rgh
func (_m *Router) Route(rgh httpserver.RouteHandler) {
	_m.Called(rgh)
}

// Router_Route_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Route'
type Router_Route_Call struct {
	*mock.Call
}

// Route is a helper method to define mock.On call
//   - rgh httpserver.RouteHandler
func (_e *Router_Expecter) Route(rgh interface{}) *Router_Route_Call {
	return &Router_Route_Call{Call: _e.mock.On("Route", rgh)}
}

func (_c *Router_Route_Call) Run(run func(rgh httpserver.RouteHandler)) *Router_Route_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(httpserver.RouteHandler))
	})
	return _c
}

func (_c *Router_Route_Call) Return() *Router_Route_Call {
	_c.Call.Return()
	return _c
}

func (_c *Router_Route_Call) RunAndReturn(run func(httpserver.RouteHandler)) *Router_Route_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRouter interface {
	mock.TestingT
	Cleanup(func())
}

// NewRouter creates a new instance of Router. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRouter(t mockConstructorTestingTNewRouter) *Router {
	mock := &Router{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
