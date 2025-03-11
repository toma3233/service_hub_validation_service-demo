package mock

import (
	context "context"
	reflect "reflect"

	v1 "dev.azure.com/service-hub-flg/service_hub_validation/_git/service_hub_validation_service.git/mygreeterv4/api/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockMyGreeterClient is a mock of MyGreeterClient interface.
type MockMyGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *MockMyGreeterClientMockRecorder
	isgomock struct{}
}

// MockMyGreeterClientMockRecorder is the mock recorder for MockMyGreeterClient.
type MockMyGreeterClientMockRecorder struct {
	mock *MockMyGreeterClient
}

// NewMockMyGreeterClient creates a new mock instance.
func NewMockMyGreeterClient(ctrl *gomock.Controller) *MockMyGreeterClient {
	mock := &MockMyGreeterClient{ctrl: ctrl}
	mock.recorder = &MockMyGreeterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMyGreeterClient) EXPECT() *MockMyGreeterClientMockRecorder {
	return m.recorder
}

// CreateResourceGroup mocks base method.
func (m *MockMyGreeterClient) CreateResourceGroup(ctx context.Context, in *v1.CreateResourceGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateResourceGroup", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateResourceGroup indicates an expected call of CreateResourceGroup.
func (mr *MockMyGreeterClientMockRecorder) CreateResourceGroup(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResourceGroup", reflect.TypeOf((*MockMyGreeterClient)(nil).CreateResourceGroup), varargs...)
}

// CreateStorageAccount mocks base method.
func (m *MockMyGreeterClient) CreateStorageAccount(ctx context.Context, in *v1.CreateStorageAccountRequest, opts ...grpc.CallOption) (*v1.CreateStorageAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateStorageAccount", varargs...)
	ret0, _ := ret[0].(*v1.CreateStorageAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStorageAccount indicates an expected call of CreateStorageAccount.
func (mr *MockMyGreeterClientMockRecorder) CreateStorageAccount(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStorageAccount", reflect.TypeOf((*MockMyGreeterClient)(nil).CreateStorageAccount), varargs...)
}

// DeleteResourceGroup mocks base method.
func (m *MockMyGreeterClient) DeleteResourceGroup(ctx context.Context, in *v1.DeleteResourceGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteResourceGroup", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteResourceGroup indicates an expected call of DeleteResourceGroup.
func (mr *MockMyGreeterClientMockRecorder) DeleteResourceGroup(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourceGroup", reflect.TypeOf((*MockMyGreeterClient)(nil).DeleteResourceGroup), varargs...)
}

// DeleteStorageAccount mocks base method.
func (m *MockMyGreeterClient) DeleteStorageAccount(ctx context.Context, in *v1.DeleteStorageAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteStorageAccount", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStorageAccount indicates an expected call of DeleteStorageAccount.
func (mr *MockMyGreeterClientMockRecorder) DeleteStorageAccount(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStorageAccount", reflect.TypeOf((*MockMyGreeterClient)(nil).DeleteStorageAccount), varargs...)
}

// ListResourceGroups mocks base method.
func (m *MockMyGreeterClient) ListResourceGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v1.ListResourceGroupResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceGroups", varargs...)
	ret0, _ := ret[0].(*v1.ListResourceGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceGroups indicates an expected call of ListResourceGroups.
func (mr *MockMyGreeterClientMockRecorder) ListResourceGroups(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceGroups", reflect.TypeOf((*MockMyGreeterClient)(nil).ListResourceGroups), varargs...)
}

// ListStorageAccounts mocks base method.
func (m *MockMyGreeterClient) ListStorageAccounts(ctx context.Context, in *v1.ListStorageAccountRequest, opts ...grpc.CallOption) (*v1.ListStorageAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStorageAccounts", varargs...)
	ret0, _ := ret[0].(*v1.ListStorageAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStorageAccounts indicates an expected call of ListStorageAccounts.
func (mr *MockMyGreeterClientMockRecorder) ListStorageAccounts(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStorageAccounts", reflect.TypeOf((*MockMyGreeterClient)(nil).ListStorageAccounts), varargs...)
}

// ReadResourceGroup mocks base method.
func (m *MockMyGreeterClient) ReadResourceGroup(ctx context.Context, in *v1.ReadResourceGroupRequest, opts ...grpc.CallOption) (*v1.ReadResourceGroupResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadResourceGroup", varargs...)
	ret0, _ := ret[0].(*v1.ReadResourceGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadResourceGroup indicates an expected call of ReadResourceGroup.
func (mr *MockMyGreeterClientMockRecorder) ReadResourceGroup(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadResourceGroup", reflect.TypeOf((*MockMyGreeterClient)(nil).ReadResourceGroup), varargs...)
}

// ReadStorageAccount mocks base method.
func (m *MockMyGreeterClient) ReadStorageAccount(ctx context.Context, in *v1.ReadStorageAccountRequest, opts ...grpc.CallOption) (*v1.ReadStorageAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadStorageAccount", varargs...)
	ret0, _ := ret[0].(*v1.ReadStorageAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadStorageAccount indicates an expected call of ReadStorageAccount.
func (mr *MockMyGreeterClientMockRecorder) ReadStorageAccount(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadStorageAccount", reflect.TypeOf((*MockMyGreeterClient)(nil).ReadStorageAccount), varargs...)
}

// SayHello mocks base method.
func (m *MockMyGreeterClient) SayHello(ctx context.Context, in *v1.HelloRequest, opts ...grpc.CallOption) (*v1.HelloReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SayHello", varargs...)
	ret0, _ := ret[0].(*v1.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello.
func (mr *MockMyGreeterClientMockRecorder) SayHello(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockMyGreeterClient)(nil).SayHello), varargs...)
}

// StartLongRunningOperation mocks base method.
func (m *MockMyGreeterClient) StartLongRunningOperation(ctx context.Context, in *v1.StartLongRunningOperationRequest, opts ...grpc.CallOption) (*v1.StartLongRunningOperationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartLongRunningOperation", varargs...)
	ret0, _ := ret[0].(*v1.StartLongRunningOperationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartLongRunningOperation indicates an expected call of StartLongRunningOperation.
func (mr *MockMyGreeterClientMockRecorder) StartLongRunningOperation(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartLongRunningOperation", reflect.TypeOf((*MockMyGreeterClient)(nil).StartLongRunningOperation), varargs...)
}

// UpdateResourceGroup mocks base method.
func (m *MockMyGreeterClient) UpdateResourceGroup(ctx context.Context, in *v1.UpdateResourceGroupRequest, opts ...grpc.CallOption) (*v1.UpdateResourceGroupResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateResourceGroup", varargs...)
	ret0, _ := ret[0].(*v1.UpdateResourceGroupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResourceGroup indicates an expected call of UpdateResourceGroup.
func (mr *MockMyGreeterClientMockRecorder) UpdateResourceGroup(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResourceGroup", reflect.TypeOf((*MockMyGreeterClient)(nil).UpdateResourceGroup), varargs...)
}

// UpdateStorageAccount mocks base method.
func (m *MockMyGreeterClient) UpdateStorageAccount(ctx context.Context, in *v1.UpdateStorageAccountRequest, opts ...grpc.CallOption) (*v1.UpdateStorageAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateStorageAccount", varargs...)
	ret0, _ := ret[0].(*v1.UpdateStorageAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStorageAccount indicates an expected call of UpdateStorageAccount.
func (mr *MockMyGreeterClientMockRecorder) UpdateStorageAccount(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStorageAccount", reflect.TypeOf((*MockMyGreeterClient)(nil).UpdateStorageAccount), varargs...)
}
