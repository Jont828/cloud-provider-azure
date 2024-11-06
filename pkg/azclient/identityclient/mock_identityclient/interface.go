// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: identityclient/interface.go
//
// Generated by this command:
//
//	mockgen -package mock_identityclient -source identityclient/interface.go -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt
//

// Package mock_identityclient is a generated GoMock package.
package mock_identityclient

import (
	context "context"
	reflect "reflect"

	armmsi "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, resourceName string) (*armmsi.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName)
	ret0, _ := ret[0].(*armmsi.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, resourceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, resourceName)
}
