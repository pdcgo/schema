package user_iface_mock

import (
    "context"
    "reflect"
    gomock "github.com/golang/mock/gomock"
    connect "connectrpc.com/connect"
    v1 "github.com/pdcgo/schema/services/user_iface/v1"
)

type MockAuthService struct {
    ctrl     *gomock.Controller
    recorder *MockAuthServiceMockRecorder
}

type MockAuthServiceMockRecorder struct {
    mock *MockAuthService
}

func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
    mock := &MockAuthService{ctrl: ctrl}
    mock.recorder = &MockAuthServiceMockRecorder{mock}
    return mock
}

func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
    return m.recorder
}

type MockAuthServiceClient struct {
    ctrl     *gomock.Controller
    recorder *MockAuthServiceClientMockRecorder
}

type MockAuthServiceClientMockRecorder struct {
    mock *MockAuthServiceClient
}

func NewMockAuthServiceClient(ctrl *gomock.Controller) *MockAuthServiceClient {
    mock := &MockAuthServiceClient{ctrl: ctrl}
    mock.recorder = &MockAuthServiceClientMockRecorder{mock}
    return mock
}

func (m *MockAuthServiceClient) EXPECT() *MockAuthServiceClientMockRecorder {
    return m.recorder
}

func (m *MockAuthService) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Login", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LoginResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAuthServiceMockRecorder) Login(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), ctx, req)
}

func (m *MockAuthServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Login", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LoginResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAuthServiceClientMockRecorder) Login(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), ctx, req)
}

func (m *MockAuthService) Logout(ctx context.Context, req *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Logout", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LogoutResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAuthServiceMockRecorder) Logout(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthService)(nil).Logout), ctx, req)
}

func (m *MockAuthServiceClient) Logout(ctx context.Context, req *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "Logout", ctx, req)
    ret0, _ := ret[0].(*connect.Response[v1.LogoutResponse])
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

func (mr *MockAuthServiceClientMockRecorder) Logout(ctx, req interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuthService)(nil).Logout), ctx, req)
}

