// Code generated by mockery v1.0.0. DO NOT EDIT.

package appnet

import (
	context "context"
	net "net"

	mock "github.com/stretchr/testify/mock"
)

// MockNetworker is an autogenerated mock type for the Networker type
type MockNetworker struct {
	mock.Mock
}

// Dial provides a mock function with given fields: addr
func (_m *MockNetworker) Dial(addr Addr) (net.Conn, error) {
	ret := _m.Called(addr)

	var r0 net.Conn
	if rf, ok := ret.Get(0).(func(Addr) net.Conn); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Conn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Addr) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DialContext provides a mock function with given fields: ctx, addr
func (_m *MockNetworker) DialContext(ctx context.Context, addr Addr) (net.Conn, error) {
	ret := _m.Called(ctx, addr)

	var r0 net.Conn
	if rf, ok := ret.Get(0).(func(context.Context, Addr) net.Conn); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Conn)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, Addr) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Listen provides a mock function with given fields: addr
func (_m *MockNetworker) Listen(addr Addr) (net.Listener, error) {
	ret := _m.Called(addr)

	var r0 net.Listener
	if rf, ok := ret.Get(0).(func(Addr) net.Listener); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Addr) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListenContext provides a mock function with given fields: ctx, addr
func (_m *MockNetworker) ListenContext(ctx context.Context, addr Addr) (net.Listener, error) {
	ret := _m.Called(ctx, addr)

	var r0 net.Listener
	if rf, ok := ret.Get(0).(func(context.Context, Addr) net.Listener); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Listener)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, Addr) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
