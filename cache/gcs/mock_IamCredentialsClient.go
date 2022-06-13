// Code generated by mockery v1.1.0. DO NOT EDIT.

package gcs

import (
	context "context"

	gax "github.com/googleapis/gax-go/v2"
	credentials "google.golang.org/genproto/googleapis/iam/credentials/v1"

	mock "github.com/stretchr/testify/mock"
)

// MockIamCredentialsClient is an autogenerated mock type for the IamCredentialsClient type
type MockIamCredentialsClient struct {
	mock.Mock
}

// SignBlob provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockIamCredentialsClient) SignBlob(_a0 context.Context, _a1 *credentials.SignBlobRequest, _a2 ...gax.CallOption) (*credentials.SignBlobResponse, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *credentials.SignBlobResponse
	if rf, ok := ret.Get(0).(func(context.Context, *credentials.SignBlobRequest, ...gax.CallOption) *credentials.SignBlobResponse); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*credentials.SignBlobResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *credentials.SignBlobRequest, ...gax.CallOption) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}