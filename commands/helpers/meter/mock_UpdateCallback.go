// Code generated by mockery v2.14.0. DO NOT EDIT.

package meter

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockUpdateCallback is an autogenerated mock type for the UpdateCallback type
type MockUpdateCallback struct {
	mock.Mock
}

// Execute provides a mock function with given fields: written, since, done
func (_m *MockUpdateCallback) Execute(written uint64, since time.Duration, done bool) {
	_m.Called(written, since, done)
}

type mockConstructorTestingTNewMockUpdateCallback interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUpdateCallback creates a new instance of MockUpdateCallback. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUpdateCallback(t mockConstructorTestingTNewMockUpdateCallback) *MockUpdateCallback {
	mock := &MockUpdateCallback{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}