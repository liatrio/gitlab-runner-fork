// Code generated by mockery v1.1.0. DO NOT EDIT.

package vault

import (
	api "github.com/hashicorp/vault/api"
	mock "github.com/stretchr/testify/mock"
)

// mockApiClientSys is an autogenerated mock type for the apiClientSys type
type mockApiClientSys struct {
	mock.Mock
}

// Health provides a mock function with given fields:
func (_m *mockApiClientSys) Health() (*api.HealthResponse, error) {
	ret := _m.Called()

	var r0 *api.HealthResponse
	if rf, ok := ret.Get(0).(func() *api.HealthResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.HealthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}